package zap

import "fmt"

//go:generate chmod +x ./zz.sh "Global invocation"
//go:generate ./zz.sh

func Hello() {
	main()
	//go:generate chmod +x ./zz.sh "Hello fn invocation"
	//go:generate ./zz.sh
}

func main() {
	fmt.Println("zzzap")
}
