package main

import (
	BasketHandler "gigsib/handler/basket"
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

func main() {
	router := vestigo.NewRouter()
	router.Post("/basket", BasketHandler.Create)
	log.Fatal(http.ListenAndServe(":3000", router))
}
