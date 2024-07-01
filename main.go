package zap

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func init() {
	exeCmd := exec.Command("go", "generate")
	var outBuffer bytes.Buffer
	multiWriter := io.MultiWriter(os.Stdout, &outBuffer)
	exeCmd.Stdout = multiWriter
	exeCmd.Stderr = multiWriter

	if err := exeCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

//go:generate chmod +x "$PWD/zz.sh"
//go:generate "$PWD/zz.sh" "Global invocation"

func Hello() {
	main()
	//go:generate chmod +x "$PWD/zz.sh"
	//go:generate "$PWD/zz.sh" "Hello invocation"
}

func main() {
	fmt.Println("zzzap")
	exeCmd := exec.Command("go", "generate", ".")
	var outBuffer bytes.Buffer
	multiWriter := io.MultiWriter(os.Stdout, &outBuffer)
	exeCmd.Stdout = multiWriter
	exeCmd.Stderr = multiWriter

	if err := exeCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
