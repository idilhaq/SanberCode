package main

import (
	"flag"
	"fmt"
)

func main() {
	var panjang = flag.Int("panjang", 5, "tolong")
	// var lebar = flag.Int("lebar", 4, "tolong")
	// var tinggi = flag.Int("tinggi", 6, "tolong")

	volumeBalok := (*panjang)
	fmt.Println(volumeBalok)
}
