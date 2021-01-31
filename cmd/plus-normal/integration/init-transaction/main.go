package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/dokshor/transbank-sdk-golang"
	"github.com/dokshor/transbank-sdk-golang/pkg/webpay"
)

func main() {
	amount := int(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := strconv.Itoa(rand.Intn(99999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	service := webpay.NewIntegrationPlusNormal()
	transaction, err := service.InitTransaction(transbank.InitTransaction{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL:  finalURL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
