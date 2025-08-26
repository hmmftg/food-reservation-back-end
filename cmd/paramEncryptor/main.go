package main

import (
	"flag"
	"log"
	"os"

	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"github.com/hmmftg/requestCore/libParams"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	paramFile := flag.String("p", "param.yaml", "Application Params")
	encryptParams := flag.Bool("c", false, "Encrypt Params")

	flag.Parse()

	wsParams, err := libParams.Load[params.FoodReservationParams](*paramFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	keys := params.GetKeys()

	if *encryptParams {
		err = libParams.EncryptParams(keys[0], keys[1], *paramFile, wsParams)
		if err != nil {
			log.Fatalln("error in EncryptParams", err)
		}
		os.Exit(0)
	}
}
