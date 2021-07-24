package controller

import (
	"log"
	"net/http"

	"github.com/quang2906/book_store_be/util"
	"github.com/stripe/stripe-go/customer"
)

func HandleEvent(w http.ResponseWriter, req *http.Request) {
	event, err := util.GetEvent(w, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(event.Type)

	if event.Type == "customer.subscription.created" {
		c, err := customer.Get(event.Data.Object["customer"].(string), nil)
		if err != nil {
			log.Fatal(err)
		}
		email := c.Metadata["FinalEmail"]
		log.Println("Subscription created by", email)
	}

}
