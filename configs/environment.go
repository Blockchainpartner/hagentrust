package configs

import (
	"log"
	"os"
)

var EthNodeURL, EthWallet, EthWalletPwd, MongodbDBName, MongodbURI, SubgraphURL string

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

	MongodbDBName, exists = os.LookupEnv("MONGODB_DB_NAME")
	if !exists {
		log.Fatal("MONGODB_DB_NAME environment variable not set")
	}

	MongodbURI, exists = os.LookupEnv("MONGODB_URI")
	if !exists {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	SubgraphURL, exists = os.LookupEnv("SUBGRAPH_URL")
	if !exists {
		log.Fatal("SUBGRAPH_URL environment variable not set")
	}

}