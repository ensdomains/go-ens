package erc20

// Hand-written minimal ERC-20 ABI (plus the mint() of the Sepolia mock tokens
// and the OpenZeppelin v5 insufficiency errors); not vendored from a deployment.

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.5 --v2 --abi contract.abi --pkg erc20 --type ERC20 --out contract.go
