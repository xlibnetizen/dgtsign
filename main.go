package main

import (
	"fmt"

	"github.com/xlibnetizen/dgtsign/dgtsign"
)

func main() {
	kq := dgtsign.Echo()
	fmt.Println("Init main..", kq)
}
