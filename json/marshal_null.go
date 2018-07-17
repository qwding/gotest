package main
import (
"encoding/json"

)
type A struct{
B string
}

func main(){	
	var a A

	res:=Struct2String(a)
	println("res:",res)
}


func Struct2String(i interface{}) string {
	content, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(content)
}

