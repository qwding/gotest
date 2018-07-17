package main 


func main(){
	run()
}

func run(){
	a := 1
	handle := func(xyz int){
		a = a+xyz
	}

	handle(2)
	println(a)
	handle(3)
	println(a)
}