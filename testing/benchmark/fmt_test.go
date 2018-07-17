package testings

import(
	"testing"
	"fmt"
)

func BenchmarkPrintln(b *testing.B){
	for i:=0;i<b.N;i++{
		fmt.Println("xyz")
	}
}

func BenchmarkPrintf(b *testing.B){
	for i:=0;i<b.N;i++{
		fmt.Printf("%s\n","xyz")
	}
}


func BenchmarkPrint(b *testing.B){
	for i:=0;i<b.N;i++{
		print("xyz")
	}
}
