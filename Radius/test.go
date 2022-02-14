package main

import (
	"fmt"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {
	packet := radius.New(radius.CodeAccessRequest, []byte(`secret`))
	rfc2865.UserName_SetString(packet, "tim")
	rfc2865.UserPassword_SetString(packet, "12345")
	fmt.Println(rfc2865.NASIPAddress_Get(packet))
	fmt.Println(rfc2865.NASPort_Strings)
	fmt.Println(rfc2865.NASPort_Get(packet))
}
