package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	// ! Set the transactionHashSignature created in step 4
	transactionHashSignature := "..."

	// ! Set the rawUnsignedTransaction created in step 3
	rawUnsignedTransaction := "..."

	// Deserialize the rawUnsignedTransaction
	unsignedTx := &types.Transaction{}
	unsignedTxBytes, err := hex.DecodeString(rawUnsignedTransaction)
	err = rlp.DecodeBytes(unsignedTxBytes, unsignedTx)
	if err != nil {
		panic(err)
	}

	// Deserialize the transactionHashSignature
	transactionHashSignatureBytes, err := hex.DecodeString(transactionHashSignature)

	// Initialize go-ethereum client
	client, err := ethclient.Dial("https://erpc.apothem.network")
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	signer := types.NewEIP155Signer(chainID)

	// Combine the unsigned transaction and the signature to create a signed transaction
	signedTx, err := unsignedTx.WithSignature(signer, transactionHashSignatureBytes)
	if err != nil {
		panic(err)
	}

	raw, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSigned raw transaction (RLP encoded): 0x%x", raw)

	// Broadcast the signed transaction to the blockchain
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nBroadcasted transaction hash: https://apothem.blocksscan.io/txs/%s\n", signedTx.Hash().String())
	}
}
