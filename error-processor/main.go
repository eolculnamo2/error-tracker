package main

import (
	"log"
	"github.com/eolculnamo2/error-tracker/error-processor/app"
)

/*
Todo next:
1) Connect to MongoDB
2) Add routes to process errors
3) Add authenticated Routes to return errors 
4) Add CORS restrictions.. If a route that provides info, it should be strictly for dashbaord otherwise any url
*/

func main() {
	log.Println("Starting Error Processor")
	app.StartRoutes()
}