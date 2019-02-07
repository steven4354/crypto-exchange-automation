package main

import (
	"github.com/routefire/routefire-core"
	"github.com/steven4354/go-gdax"
	"log"
	"time"
)

// make sure these items have "transfer" permission
const (
	secret = ""
	key = ""
	passphrase = ""
)

var (
	appStartTime = time.Now()
	lastTransferRunTime = time.Time{}
	firstRunBegan = false
)

func DoOneDollaWithdrawalOnGDAX() {
	uh := routefire.NewGDAXUserHandler("test", secret, key, passphrase, []routefire.CryptoAssetSymbol{})
	b, err := uh.AllBalances()
	log.Printf("all balances %+v err %v \n", b, err)

	// get id of payment method

	client := gdax.NewClient(
		secret,
		key,
		passphrase,
	)

	log.Printf("starting getting payment info, please wait...")

	accs, err := client.GetPaymentMethods()

	var acc gdax.PaymentMethod
	var id string
	if len(accs) > 0 {
		acc = accs[0]
		id = acc.ID
	}

	log.Printf("id %+v err %v", id, err)

	// take 1% of usd balance and move to bank

	log.Printf("starting withdrawing, please wait...")

	info := gdax.PaymentMethodWithdrawalRequest{1.00, "USD", id}
	res, err := client.WithdrawToPaymentMethods(info)

	log.Printf("res %+v err %v", res, err)

	// make a loop

	// add a time thing to do once a day
}

func main(){
	for true {
		log.Println("checking...")
		t2 := time.Now()
		diff := t2.Sub(lastTransferRunTime)

		if !firstRunBegan {
			DoOneDollaWithdrawalOnGDAX()
			lastTransferRunTime = time.Now()
		} else if diff.Seconds() > 86400.0 {
			DoOneDollaWithdrawalOnGDAX()
			lastTransferRunTime = time.Now()
		}

		firstRunBegan = true
		time.Sleep(2 * time.Second)
	}
}

