package crypto_exchange_automation

import (
	"log"
	"time"
)

// keep track of start time
// on 24 hours since start time do a transfer to bank

// TODO: (s)
// create a systems manager like cts that maps user to WorkingTransfersProvider
// give a init 1 user to ^^
// give a init all user to ^^
// move the start time to db
// move the last run time to db
// create a clean stop function

type WorkingTransfersProvider struct {
	StartTime time.Time
	LastRunTime time.Time
	TransferCurrency string
	TransferAmount string
}

type CustomerTransferServer struct {
	UserId string
	VenueWorkingTransferProviders map[string]WorkingTransfersProvider // map[venue]WorkingTransfersProvider
}

func (c *CustomerTransferServer) RunLoop() {
	for true {
		log.Printf("hello")

	}
}
