package server

import (
    "github.com/Blockchainpartner/hagentrust/configs"
    "github.com/Blockchainpartner/hagentrust/pkg/db"
)

func Init(port string) error {
    configs.InitEnvironment()
    db.Init()

    r := NewRouter()
    err := r.Run(":" + port)
    return err
}
