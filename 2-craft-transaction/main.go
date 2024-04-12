package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {

	// ! Set the wallet address created in step 1
	walletAddress := common.HexToAddress("0x...")

	// ! Set the destination address that funds will be sent to (can use faucet)
	destinationAddress := common.HexToAddress("0x45e9B0983dD4544C8f5e8Ba55f8e7Cb4Bf0543cd")

	// Initialize go-ethereum client
	client, err := ethclient.Dial("https://erpc.apothem.network")
	if err != nil {
		panic(err)
	}

	// Check wallet balance at chain path m/44/51
	balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		panic(err)
	}
	// Convert balance from wei to XDC
	balance = new(big.Int).Div(balance, big.NewInt(1000000000000000000))
	fmt.Println("Balance at account:", walletAddress, "=", (balance), "XDC")

	// Build unsigned transaction for sending 1 XDC to destination address
	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	fmt.Println("Nonce at account", walletAddress, ":", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	fmt.Println("Current gas price:", gasPrice)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Current chainID:", chainID)
	signer := types.NewEIP155Signer(chainID)

	// Craft Legacy type transaction
	unsignedTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &destinationAddress,
		Value:    big.NewInt(1000000000000000000), // 1 XDC
		Gas:      uint64(3000000),
		GasPrice: gasPrice,
		Data:     nil,
	})

	// Serialize the unsigned transaction
	bytes, err := rlp.EncodeToBytes(unsignedTx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Raw unsigned transaction:", hex.EncodeToString(bytes))
	fmt.Println("Unsigned transaction hash (to be signed):", hex.EncodeToString(signer.Hash(unsignedTx).Bytes()))
}
