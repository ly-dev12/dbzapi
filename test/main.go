package main

import (
	dbzapi "dbzapi/lib"
	"fmt"
)

func main() {
	dbzmethods := dbzapi.SetConfig("Kakarot", 23453)

	lp, msg := dbzapi.GetLevel(dbzmethods)

	fmt.Println(lp, msg)
}
