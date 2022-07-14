package main

import (
	"fmt"

	go_odoo "github.com/skilld-labs/go-odoo"
)

var (
	ClientPm go_odoo.Client
)

var (
	OdooLoginPM    = "loginPM"
	OdooPasswordPM = "passwordPM"
	OdooDatabasePM = "databasePM"
	OdooUrlPM      = "urlPM"
)

func main() {

	odooClientPm, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    OdooLoginPM,
		Password: OdooPasswordPM,
		Database: OdooDatabasePM,
		URL:      OdooUrlPM,
	})
	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
		return
	}

	ClientPm = *odooClientPm

	GetCatalogFrom1c()
}
