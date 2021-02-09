package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main(){

	eth_client, err := ethclient.Dial("ws://127.0.0.1:7545")
	if err != nil{
		fmt.Printf("%v", err)
	}

	contract, err := NewTestContract(common.HexToAddress("0xe1eDC959706BBD30D1CB881E1EF75E1edEFe753A"), eth_client)
	if err != nil{
		fmt.Printf("%v", err)
	}

	//private key
	ecdsa, err := crypto.HexToECDSA("0beefda68b751c53c27c5bb0b5840f1b77dbb4e40cc182aea45af0212cf5b744")
	if err != nil{
		fmt.Printf("%v", err)
	}

	//uses the private key to sign the transaction. Only needed for the store function
	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(ecdsa, big.NewInt(5777))
	if err != nil{
		fmt.Printf("%v", err)
	}

	broadcastedTransaction, err := contract.Store(keyedTransactor,"giulia&Michael")
	if err != nil{
		fmt.Printf("%v", err)
	}
	fmt.Printf("%s",broadcastedTransaction.Hash().String())

}
