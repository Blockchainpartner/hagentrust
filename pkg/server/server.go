package server

import (
    "github.com/Blockchainpartner/hagentrust/configs"
)

func Init(port string) error {
    configs.InitEnvironment()

    r := NewRouter()
    err := r.Run(":" + port)
    return err
}
