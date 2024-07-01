package main

import "fmt"

//go:generate chmod +x ./zz.sh
//go:generate ./zz.sh

func Hello() {
	main()
}

func main() {
	fmt.Println("zzzap")
}
