package main

var (
	GitTag    = ""
	BuildTime = ""
)

func main() {
	println("tag:", GitTag)
	println("build:", BuildTime)
}
