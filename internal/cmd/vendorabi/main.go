// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Command vendorabi copies contract ABIs, deployment addresses, decoded
// constructor arguments and (optionally) creation bytecode from a pinned
// ensdomains/contracts-v2 revision into this repository.
//
// It writes:
//
//	contracts/v2/<pkg>/contract.abi          ABI array, entry order preserved
//	contracts/v2/deployments/<network>.json  address + constructor args per artifact
//	contracts/v2/manifest.json               source ref and sha256 of every ABI
//	contracts/v1/<pkg>/contract.abi          v1 artifacts fetched from ens-contracts
//	internal/simchain/testdata/<Name>.bin    creation bytecode (with -bytecode)
//
// With -check it regenerates everything into a temporary directory and fails
// if the result differs from what is committed.
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ContractsV2Ref is the pinned ensdomains/contracts-v2 commit.
const ContractsV2Ref = "0f10caf50054f1af96197350aed6a4f072e4f205"

// ENSContractsRef is the ensdomains/ens-contracts submodule commit pinned by
// contracts-v2 at ContractsV2Ref; the v1 artifacts are taken from there.
const ENSContractsRef = "f0151aaf8a53b5b8f725a584fc194997c8226fbd"

const (
	contractsV2Raw  = "https://raw.githubusercontent.com/ensdomains/contracts-v2/%s/contracts/deployments/%s/%s.json"
	ensContractsRaw = "https://raw.githubusercontent.com/ensdomains/ens-contracts/%s/deployments/%s/%s.json"
)

var networks = []string{"mainnet", "sepolia"}

// binding maps a deployment artifact to a Go binding package.
type binding struct {
	pkg  string // directory under contracts/v2 (or contracts/v1 when v1 is true)
	file string // artifact file name without .json
	v1   bool
	// mainnetOnly skips the Sepolia compatibility check: the binding follows
	// the mainnet artifact and Sepolia may lag it by a redeploy.
	mainnetOnly bool
}

var bindings = []binding{
	{pkg: "universalresolver", file: "UniversalResolverV2"},
	{pkg: "universalhelper", file: "UniversalHelper"},
	{pkg: "permissionedregistry", file: "UserRegistryImpl"},
	{pkg: "wrapperregistry", file: "WrapperRegistryImpl", mainnetOnly: true},
	{pkg: "ethregistrar", file: "ETHRegistrar"},
	{pkg: "ethrenewerv1", file: "ETHRenewerV1"},
	{pkg: "rentpriceoracle", file: "StandardRentPriceOracle"},
	{pkg: "permissionedresolver", file: "PermissionedResolverImpl"},
	{pkg: "publicresolver", file: "PublicResolverV2"},
	{pkg: "verifiablefactory", file: "VerifiableFactory"},
	{pkg: "reverseregistraradapter", file: "ReverseRegistrarAdapter"},
	{pkg: "defaultreverseregistraradapter", file: "DefaultReverseRegistrarAdapter"},
	{pkg: "labelstore", file: "LabelStore"},
	{pkg: "compositeresolver", file: "ENSV1Resolver"},
	{pkg: "namewrapper", file: "NameWrapper", v1: true},
	{pkg: "defaultreverseregistrar", file: "DefaultReverseRegistrar", v1: true},
}

// simchainBytecode lists artifacts whose creation bytecode the simulated
// backend harness deploys. Mocks exist only on sepolia.
var simchainBytecode = map[string]string{
	"ContractNamer_Implementation": "sepolia",
	"ContractNamer_Proxy":          "sepolia",
	"LabelStore":                   "mainnet",
	"ETHRegistry":                  "mainnet",
	"MockUSDC":                     "sepolia",
	"StandardRentPriceOracle":      "mainnet",
	"ETHRegistrar":                 "mainnet",
	"UniversalResolverV2":          "mainnet",
	"UniversalHelper":              "mainnet",
	"VerifiableFactory":            "mainnet",
	"UserRegistryImpl":             "mainnet",
	"PermissionedResolverImpl":     "mainnet",
}

type artifact struct {
	Address  common.Address  `json:"address"`
	ABI      json.RawMessage `json:"abi"`
	ArgsData string          `json:"argsData"` //nolint:tagliatelle // upstream artifact naming
	Bytecode string          `json:"bytecode"`
}

type deploymentEntry struct {
	Address         common.Address `json:"address"`
	ConstructorArgs map[string]any `json:"constructorArgs,omitempty"` //nolint:tagliatelle // matches the artifacts
}

type deployments struct {
	ChainID   uint64                     `json:"chainId"` //nolint:tagliatelle // matches the artifacts
	Ref       string                     `json:"ref"`
	Contracts map[string]deploymentEntry `json:"contracts"`
}

type manifest struct {
	ContractsV2Ref  string            `json:"contractsV2Ref"`  //nolint:tagliatelle // matches the artifacts
	ENSContractsRef string            `json:"ensContractsRef"` //nolint:tagliatelle // matches the artifacts
	ABIs            map[string]string `json:"abis"`            // package path -> sha256 of contract.abi
}

var chainIDs = map[string]uint64{"mainnet": 1, "sepolia": 11155111}

func main() {
	var (
		local    = flag.String("local", "", "local contracts-v2 checkout (contracts/deployments dir); fetches from GitHub when empty")
		localV1  = flag.String("local-v1", "", "local directory holding ens-contracts mainnet artifacts (NameWrapper.json ...)")
		root     = flag.String("root", ".", "repository root")
		bytecode = flag.Bool("bytecode", false, "also write simchain creation bytecode")
		check    = flag.Bool("check", false, "verify committed files match the pinned sources")
	)
	flag.Parse()

	src := newSource(*local, *localV1)
	out := *root
	if *check {
		tmp, err := os.MkdirTemp("", "vendorabi-check")
		if err != nil {
			fatal(err)
		}
		defer os.RemoveAll(tmp)
		out = tmp
	}
	if err := run(src, out, *bytecode || *check); err != nil {
		fatal(err)
	}
	if *check {
		if err := compare(*root, out); err != nil {
			fatal(err)
		}
		fmt.Println("vendored ABIs are up to date")
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "vendorabi:", err)
	os.Exit(1)
}

type source struct {
	local   string
	localV1 string
	client  *http.Client
	cache   map[string]*artifact
}

func newSource(local, localV1 string) *source {
	return &source{local: local, localV1: localV1, client: http.DefaultClient, cache: map[string]*artifact{}}
}

func (s *source) artifact(network, name string, v1 bool) (*artifact, error) {
	key := fmt.Sprintf("%v/%s/%s", v1, network, name)
	if a, ok := s.cache[key]; ok {
		return a, nil
	}
	var raw []byte
	var err error
	switch {
	case v1 && s.localV1 != "":
		raw, err = os.ReadFile(filepath.Join(s.localV1, name+".json"))
	case v1:
		raw, err = s.fetch(fmt.Sprintf(ensContractsRaw, ENSContractsRef, network, name))
	case s.local != "":
		raw, err = os.ReadFile(filepath.Join(s.local, network, name+".json"))
	default:
		raw, err = s.fetch(fmt.Sprintf(contractsV2Raw, ContractsV2Ref, network, name))
	}
	if err != nil {
		return nil, fmt.Errorf("%s/%s: %w", network, name, err)
	}
	var a artifact
	if err := json.Unmarshal(raw, &a); err != nil {
		return nil, fmt.Errorf("%s/%s: %w", network, name, err)
	}
	s.cache[key] = &a
	return &a, nil
}

func (s *source) fetch(url string) ([]byte, error) {
	resp, err := s.client.Get(url) //nolint:noctx // one-shot CLI
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: HTTP %d", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// extraArtifacts are deployed contracts that addresses.md does not list but
// the simulated chain needs. They are skipped on networks that lack them.
var extraArtifacts = []string{"ContractNamer_Implementation", "ContractNamer_Proxy", "MockUSDC", "MockDAI"}

// listNetwork returns the artifact names in a network namespace: the table
// in addresses.md (read locally or fetched) plus the extras above. Both
// modes read the same table so their output is identical.
func (s *source) listNetwork(network string) ([]string, error) {
	var raw []byte
	var err error
	if s.local != "" {
		raw, err = os.ReadFile(filepath.Join(s.local, network, "addresses.md"))
	} else {
		raw, err = s.fetch(fmt.Sprintf("https://raw.githubusercontent.com/ensdomains/contracts-v2/%s/contracts/deployments/%s/addresses.md", ContractsV2Ref, network))
	}
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var names []string
	add := func(n string) {
		if !seen[n] {
			seen[n] = true
			names = append(names, n)
		}
	}
	for _, line := range strings.Split(string(raw), "\n") {
		if !strings.HasPrefix(line, "| ") || strings.HasPrefix(line, "| Contract") || strings.HasPrefix(line, "| ---") {
			continue
		}
		cells := strings.Split(line, "|")
		if len(cells) < 3 {
			continue
		}
		add(strings.TrimSpace(cells[1]))
	}
	for _, extra := range extraArtifacts {
		if s.exists(network, extra) {
			add(extra)
		}
	}
	sort.Strings(names)
	return names, nil
}

// exists reports whether an artifact can be read for the network.
func (s *source) exists(network, name string) bool {
	_, err := s.artifact(network, name, false)
	return err == nil
}

func run(src *source, out string, withBytecode bool) error {
	m := manifest{ContractsV2Ref: ContractsV2Ref, ENSContractsRef: ENSContractsRef, ABIs: map[string]string{}}

	// ABIs.
	for _, b := range bindings {
		a, err := src.artifact("mainnet", b.file, b.v1)
		if err != nil {
			return err
		}
		ver := "v2"
		if b.v1 {
			ver = "v1"
		}
		dir := filepath.Join(out, "contracts", ver, b.pkg)
		formatted, err := formatABI(a.ABI)
		if err != nil {
			return fmt.Errorf("%s: %w", b.file, err)
		}
		if err := writeFile(filepath.Join(dir, "contract.abi"), formatted); err != nil {
			return err
		}
		sum := sha256.Sum256(formatted)
		m.ABIs["contracts/"+ver+"/"+b.pkg] = hex.EncodeToString(sum[:])
		if !b.v1 && !b.mainnetOnly {
			// Sepolia may lag mainnet by a redeploy; its functions and events
			// must still match (the bindings are shared), while custom errors
			// may differ since newer builds only add them.
			sa, err := src.artifact("sepolia", b.file, false)
			if err != nil {
				return err
			}
			if err := compatibleABIs(a.ABI, sa.ABI); err != nil {
				return fmt.Errorf("%s: mainnet and sepolia ABIs are incompatible: %w", b.file, err)
			}
		}
	}

	// Deployment tables.
	for _, network := range networks {
		names, err := src.listNetwork(network)
		if err != nil {
			return err
		}
		d := deployments{ChainID: chainIDs[network], Ref: ContractsV2Ref, Contracts: map[string]deploymentEntry{}}
		for _, name := range names {
			a, err := src.artifact(network, name, false)
			if err != nil {
				return err
			}
			entry := deploymentEntry{Address: a.Address}
			if args, err := decodeConstructorArgs(a); err != nil {
				return fmt.Errorf("%s/%s constructor args: %w", network, name, err)
			} else if len(args) > 0 {
				entry.ConstructorArgs = args
			}
			d.Contracts[name] = entry
			if withBytecode {
				if net, ok := simchainBytecode[name]; ok && net == network {
					code := strings.TrimPrefix(a.Bytecode, "0x")
					if err := writeFile(filepath.Join(out, "internal", "simchain", "testdata", name+".bin"), []byte(code+"\n")); err != nil {
						return err
					}
				}
			}
		}
		raw, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return err
		}
		if err := writeFile(filepath.Join(out, "contracts", "v2", "deployments", network+".json"), append(raw, '\n')); err != nil {
			return err
		}
	}

	raw, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(filepath.Join(out, "contracts", "v2", "manifest.json"), append(raw, '\n'))
}

// compatibleABIs reports an error when two ABIs differ in any function or
// event signature. Custom errors may differ.
func compatibleABIs(a, b json.RawMessage) error {
	sigs := func(raw json.RawMessage) (map[string]bool, error) {
		var entries []struct {
			Type   string `json:"type"`
			Name   string `json:"name"`
			Inputs []struct {
				Type string `json:"type"`
			} `json:"inputs"`
		}
		if err := json.Unmarshal(raw, &entries); err != nil {
			return nil, err
		}
		out := map[string]bool{}
		for _, e := range entries {
			if e.Type != "function" && e.Type != "event" {
				continue
			}
			var types []string
			for _, in := range e.Inputs {
				types = append(types, in.Type)
			}
			out[e.Type+" "+e.Name+"("+strings.Join(types, ",")+")"] = true
		}
		return out, nil
	}
	as, err := sigs(a)
	if err != nil {
		return err
	}
	bs, err := sigs(b)
	if err != nil {
		return err
	}
	for sig := range as {
		if !bs[sig] {
			return fmt.Errorf("%s missing on sepolia", sig)
		}
	}
	for sig := range bs {
		if !as[sig] {
			return fmt.Errorf("%s missing on mainnet", sig)
		}
	}
	return nil
}

// formatABI re-indents an ABI array without reordering entries. abigen derives
// overload suffixes (addr/addr0) from entry order, so order is significant.
func formatABI(raw json.RawMessage) ([]byte, error) {
	var entries []json.RawMessage
	if err := json.Unmarshal(raw, &entries); err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("empty abi")
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		return nil, err
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

func decodeConstructorArgs(a *artifact) (map[string]any, error) {
	data := strings.TrimPrefix(a.ArgsData, "0x")
	if data == "" {
		return nil, nil
	}
	parsed, err := abi.JSON(bytes.NewReader(a.ABI))
	if err != nil {
		return nil, err
	}
	if len(parsed.Constructor.Inputs) == 0 {
		return nil, nil
	}
	raw, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	values, err := parsed.Constructor.Inputs.Unpack(raw)
	if err != nil {
		return nil, err
	}
	out := map[string]any{}
	for i, in := range parsed.Constructor.Inputs {
		name := in.Name
		if name == "" {
			name = fmt.Sprintf("arg%d", i)
		}
		out[name] = jsonValue(values[i])
	}
	return out, nil
}

// jsonValue converts an ABI-decoded Go value into a JSON-friendly form:
// addresses as checksummed hex, big integers as decimal strings, byte arrays
// as hex, tuples as objects, slices as arrays.
func jsonValue(v any) any {
	switch x := v.(type) {
	case common.Address:
		return x.Hex()
	case *big.Int:
		return x.String()
	case []byte:
		return "0x" + hex.EncodeToString(x)
	case string, bool, uint8, uint16, uint32, uint64, int8, int16, int32, int64:
		return x
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() { //nolint:exhaustive // scalars handled above, rest stringified
	case reflect.Array:
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			b := make([]byte, rv.Len())
			reflect.Copy(reflect.ValueOf(b), rv)
			return "0x" + hex.EncodeToString(b)
		}
		fallthrough
	case reflect.Slice:
		out := make([]any, rv.Len())
		for i := range out {
			out[i] = jsonValue(rv.Index(i).Interface())
		}
		return out
	case reflect.Struct:
		out := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Type().Field(i)
			if !f.IsExported() {
				continue
			}
			out[lowerFirst(f.Name)] = jsonValue(rv.Field(i).Interface())
		}
		return out
	default:
		return fmt.Sprint(v)
	}
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil { //nolint:gosec // generated source files
		return err
	}
	return os.WriteFile(path, data, 0o644) //nolint:gosec // generated source files
}

// compare checks that every file generated under tmp is byte-identical to the
// committed copy under root.
func compare(root, tmp string) error {
	var mismatched []string
	err := filepath.WalkDir(tmp, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(tmp, path)
		if err != nil {
			return err
		}
		want, err := os.ReadFile(path) //nolint:gosec // paths come from our own temp dir
		if err != nil {
			return err
		}
		got, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil || !bytes.Equal(want, got) {
			mismatched = append(mismatched, rel)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(mismatched) > 0 {
		return fmt.Errorf("stale vendored files (run `go generate ./contracts/v2`): %s", strings.Join(mismatched, ", "))
	}
	return nil
}
