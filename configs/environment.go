package configs

import (
	"log"
	"os"
)

var EthNodeURL, EthWallet, EthWalletPwd, SubgraphURL string

func InitEnvironment() {
	var exists bool

	EthNodeURL, exists = os.LookupEnv("ETH_NODE_URL")
	if !exists {
		log.Fatal("ETH_NODE_URL environment variable not set")
	}

	EthWallet, exists = os.LookupEnv("ETH_WALLET")
	if !exists {
		log.Fatal("ETH_WALLET environment variable not set")
	}

	EthWalletPwd, exists = os.LookupEnv("ETH_WALLET_PWD")
	if !exists {
		log.Fatal("ETH_WALLET_PWD environment variable not set")
	}

	SubgraphURL, exists = os.LookupEnv("SUBGRAPH_URL")
	if !exists {
		log.Fatal("SUBGRAPH_URL environment variable not set")
	}

}