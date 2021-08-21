package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	main1()
	return
	//Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	//Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Printf("address[%d][%v]\n", len(address), address)

	//Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Printf("privateKey[%d][%v]\n", len(privateKey), privateKey)

}

func main1() {
	for i := 0; i < 10; i++ {
		//Create an account
		key, err := crypto.GenerateKey()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		//Get the address
		address := crypto.PubkeyToAddress(key.PublicKey).Hex()
		//Get the private key
		privateKey := hex.EncodeToString(key.D.Bytes())
		fmt.Printf(",[\"%v\",\"%v\"]", address, privateKey)
	}

}
