package main

import (
	"fmt"

	"xlibnetizen/dgtsign"
)

func main() {
	kq := dgtsign.Echo()
	fmt.Println("Init main..", kq)
}
