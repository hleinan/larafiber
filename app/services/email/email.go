package email

import (
	sp "github.com/SparkPost/gosparkpost"
	"log"
	"os"
)

func SendEmail(to string, title string, message string) {
	apiKey := os.Getenv("sparkpost")
	cfg := &sp.Config{
		BaseUrl:    "https://api.eu.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
		log.Println(err)
	}

	// Create a Transmission using an inline Recipient List
	// and inline email Content.
	tx := &sp.Transmission{
		Recipients: []string{to},
		Content: sp.Content{
			HTML:    message,
			From:    "post@aning.no",
			Subject: title,
		},
	}
	id, _, err := client.Send(tx)
	if err != nil {
		log.Println(err)
	}

	// The second value returned from Send
	// has more info about the HTTP response, in case
	// you'd like to see more than the Transmission id.
	log.Printf("Transmission sent with id [%s]\n", id)
}

// 25e2c2140b342b724c7b145f980c139a35fb7c26
