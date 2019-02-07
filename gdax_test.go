package crypto_exchange_automation

/*
package routefire

import (
	"github.com/steven4354/go-gdax"
	"log"
	"testing"
)

const (
	secret = ""
	key = ""
	passphrase = ""
)

// this test will show logs such as 2018/11/06 21:00:08 WS> 0 messages processed last cycle...
// this is normal and will stop once the trades are gotten
func TestGetPaymentInfo(t *testing.T) {
	log.Printf("\n---------\nGDAX PAYMENT INFO\n---------\n")

	//mkts := []CryptoAssetSymbol{CryptoAssetSymbol("btc"), CryptoAssetSymbol("eur"), CryptoAssetSymbol("usd")}
	//stevensGdax := NewGDAXUserHandler("GDAXPAYMENTTEST", "KDL1srj8KGzyjQ+rEqDSQRhIcrRIo2Hdm0h6jp2q/eygh9ZycN8JE2Yf6fMtzNNCj0OV7oIRo6t7NMpOf4Zs+w==", "5cb53aac08af25140f8e9a1ed04b6831", "54y4q99xkzn", mkts)

	client := gdax.NewClient(
		"",
		"",
		"",
		)

	log.Printf("starting getting payment info, please wait...")

	accs, err := client.GetPaymentMethods()

	var acc gdax.PaymentMethod
	if len(accs) > 0 {
		acc = accs[0]
	}

	log.Printf("acc %+v err %v", acc, err)
}

func TestWithdrawToPaymentMethod(t *testing.T) {

	log.Printf("\n---------\nGDAX WITHDRAW TO PAYMENT METHOD\n---------\n")

	uh := NewGDAXUserHandler("test", secret, key, passphrase, []CryptoAssetSymbol{})

	b, err := uh.AllBalances()

	log.Printf("gdax balance %+v err %v", b, err)
}
 */