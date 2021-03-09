package main

import(
	"github.com/koala-proptech/goconf/consul"
	"fmt"
)
func main(){

	a:=consul.GetString("database.read")
	fmt.Println(a)
}