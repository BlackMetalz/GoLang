package main

import (
	"fmt"
	"os"

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
	fmt.Printf("Passed connect to client")

	key, err := aero.NewKey("ns_aes_test", "user_info",1)
	panicOnError(err)

	rec, err := client.Get(nil, key) // reads all the bins
	panicOnError(err)

	if rec != nil {
		printOK("%v", rec.Bins)
	}


}


func printOK(format string, a ...interface{}) {
	fmt.Printf("ok: "+format+"\n", a...)
	os.Exit(0)
}
