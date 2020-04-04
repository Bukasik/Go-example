package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(time.Format("Mon, 02 Jan 2006 15:04:05 MST"))
	}
}
