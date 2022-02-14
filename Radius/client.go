package main

import (
	"context"
	"log"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {
	packet := radius.New(radius.CodeAccessRequest, []byte(`123456`))
	rfc2865.UserName_SetString(packet, "thaianh47")
	//rfc2865.UserPassword_SetString(packet, "")
	response, err := radius.Exchange(context.Background(), packet, "192.168.3.104:1812")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Code:", response.Code)
}