package main

import (
	"fmt"
	"github.com/alapierre/go-ksef-client/ksef/api"
	"github.com/alapierre/go-ksef-client/ksef/model"
	"os"
)

func main() {

	nip := "9781399259"
	token := "30AC53BF6313480A4C12278907E718C82086E19FD56DF3F43C889A28572FDD4A"
	pathToKey := "keys/test/publicKey.pem"

	client := api.New(api.Test)
	sessionService := api.NewSessionService(client)
	invoiceService := api.NewInvoiceService(client)

	sessionToken, err := sessionService.LoginByToken(nip, model.ONIP, token, pathToKey)

	if err != nil {
		panic(err)
	}

	fmt.Printf("session token: %s\n", sessionToken.SessionToken.Token)

	content, err := os.ReadFile("FA2.xml")
	if err != nil {
		fmt.Println("Can't read invoice file")
		panic(err)
	}

	invoice, err := invoiceService.SendInvoice(content, sessionToken.SessionToken.Token)
	if err != nil {
		fmt.Println("Can't send invoice")
		panic(err)
	}

	fmt.Printf("invoice elementReferenceNumber: %s, processing code: %d\n", invoice.ElementReferenceNumber, invoice.ProcessingCode)
}
