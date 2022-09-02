package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sergera/marketplace-vendor-edge-api/internal/api"
	"github.com/sergera/marketplace-vendor-edge-api/internal/conf"
)

func main() {
	conf := conf.GetConf()

	mux := http.NewServeMux()

	vendorAPI := api.NewVendorAPI()

	mux.HandleFunc("/send-order", vendorAPI.SendOrder)

	srv := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: mux,
	}

	fmt.Printf("starting application on port %s", conf.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
