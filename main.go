package main

import (
	"Go-APIEmulator/router"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("### Start Go-APIEmulator...")
	port := 8688
	log.Println("HTTP Server Started in :" + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router.NewRouter()))
}
