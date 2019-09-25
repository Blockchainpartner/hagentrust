package util

import (
	"log"
	"strings"

	"github.com/Blockchainpartner/hagentrust/configs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EthClientInit() (auth *bind.TransactOpts, contract string) {
	/*connection to the ethereum client*/
	nodeURL := configs.EthNodeURL
	_, err := ethclient.Dial("https://"+nodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum test client: %v", err)
	}
	/*connection to the user's wallet*/
	wallet := configs.EthWallet
	walletPwd := configs.EthWalletPwd
	auth, err = bind.NewTransactor(strings.NewReader(wallet), walletPwd)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	return auth, contract
}
