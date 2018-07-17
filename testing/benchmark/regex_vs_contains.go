package testings

import (
	"testing"
	"strings"
	"regexp"
)

var str = "qwertyuiopqwertyuiowertyui"


func BenchmarkStrings(b *testing.B){
	for i:= 0;i < b.N;i++{
		strings.Index(str,"qwer")
	}
}

func BenchmarkRegex(b *testing.B){
	for i:=0;i<b.N;i++{
		regexp.MatchString("qwer",str)
	}
}