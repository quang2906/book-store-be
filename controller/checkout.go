package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
)

var PriceId = "price_1IuSOTKNyEKTIrMRyg1Mh93X"

func checkout(email string) (*stripe.CheckoutSession, error) {
	var discounts []*stripe.CheckoutSessionDiscountParams

	discounts = []*stripe.CheckoutSessionDiscountParams{
		&stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String("FMARC"),
		},
	}

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("FinalEmail", email)
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, err
	}

	meta := map[string]string{
		"FinalEmail": email,
	}

	log.Println("Creating meta for user: ", meta)

	params := &stripe.CheckoutSessionParams{
		Customer:   &newCustomer.ID,
		SuccessURL: stripe.String("https://www.youtube.com/channel/UCzgn3FvGR1UK_0M0B6GiLug"),
		CancelURL:  stripe.String("https://www.youtube.com/channel/UCzgn3FvGR1UK_0M0B6GiLug"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Discounts: discounts,
		Mode:      stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(PriceId),
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(7),
			Metadata:        meta,
		},
	}
	return session.New(params)
}

type EmailInput struct {
	Email string `json:"email"`
}

type SessionOutput struct {
	Id string `json:"id"`
}

func CheckoutCreator(w http.ResponseWriter, req *http.Request) {
	input := &EmailInput{}
	err := json.NewDecoder(req.Body).Decode(input)
	if err != nil {
		log.Fatal(err)
	}

	stripeSession, err := checkout(input.Email)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(&SessionOutput{Id: stripeSession.ID})

	if err != nil {
		log.Fatal(err)
	}
}
