package main

import (
	//	"fmt"

	aero "github.com/aerospike/aerospike-client-go"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// define a client to connect to
	client, err := aero.NewClient("10.3.48.54", 3000)
	panicOnError(err)

	key, err := aero.NewKey("ns_aes_test", "user_info", "key")
	panicOnError(err)

	bin1 := aero.NewBin("data",1) // Value find
	bin2 := aero.NewBin("data",11) // Value Update to

	// write the bins
	err = client.PutBins(nil, key, bin1,bin2)
	panicOnError(err)

	// delete the key, and check if key exists
	//existed, err := client.Delete(nil, key)
	//panicOnError(err)
	//fmt.Printf("Record existed before delete? %v\n", existed)
}
