package main

import (
	"log"
	"github.com/yenchieh/APIinGo/router"
)



func main() {
	r := router.New()
	log.Fatal(r.Run(":8081"))
}


