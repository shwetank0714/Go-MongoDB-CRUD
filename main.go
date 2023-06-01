package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shwetank0714/mongodbapi/router"
)

func main()  {
	fmt.Println(" ------ MongoDB Crud Operations -------")
	fmt.Println("Local Server is Getting started.....")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8002",r))

	fmt.Print("Server running on Port 8002 ...")
	
}