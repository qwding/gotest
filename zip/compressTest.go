package main

import (
	"fmt"
	"github.com/qwding/cae/zip"
	"os"
	"time"
)

func main() {
	fmt.Println("start!")
	times := time.Now()
	dockName := fmt.Sprintln("Docker", times.Format("2006_01_02"), ".zip")

	f, errZip := os.Create("d:/testcode/" + dockName)
	if errZip != nil {
		fmt.Println("[ errZip ]", errZip)
	}
	defer f.Close()
	z := zip.New(f)
	z.AddEmptyDir("123")
	errAddFile := z.AddFile("main1.go", "d:/testcode/main1.go")
	fmt.Println("list ", z.List("d:/testcode/Docker.zip"))
	if errAddFile != nil {
		fmt.Println("[ errAddFile ]", errAddFile)
	}
	z.AddFile("newmain.go", "d:/testcode/newmain.go")
	z.Flush()
}
