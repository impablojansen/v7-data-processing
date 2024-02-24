package main

import (
	"fmt"
	"time"
	"tratar_base/reg_manager"
)

func main() {
	start := time.Now()
	reg_manager.RegManager("./reg_01.xlsx")
	end := time.Now()

	fmt.Println(end.Sub(start))
}
