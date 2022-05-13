package main

import (
	"log"

	"github.com/WisdomEnigma/vault-keygen/vault"
	"github.com/hashicorp/vault/api"
)

const VAULT_PATH = "secret/data/enigma_secret"

func main() {

	client, err := api.NewClient(&api.Config{Address: "http://127.0.0.1:8200"})
	if err != nil {
		log.Fatalln("Client instance failed:", err.Error())
		return
	}

	client.SetToken("")

	_vault := vault.NewClient(client)

	service := map[string]interface{}{
		"data": map[string]interface{}{
			"hello": "world"},
	}

	data, err := _vault.SaveKeygen(vault.Keygen{Vault_path: VAULT_PATH, Vault_record: service})
	if err != nil {
		log.Fatalln("Write keygen failed:", err)
		return
	}

	log.Println("Write keygen succeeded:", data)
	secret, err := _vault.GetKeygen(vault.Keygen{Vault_path: VAULT_PATH})
	if err != nil {
		log.Fatalln("Get keygen failed:", err)
		return
	}
	log.Println("Get keygen succeeded:", secret.Data["data"])
}
