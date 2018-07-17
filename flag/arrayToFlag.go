package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	first := flag.String("f", "", "first flag")
	second := flag.String("s", "", "second flag")
	flag.Parse()

	fmt.Println("first ", *first, "second", *second)

	flagArray := flag.Args()

	fmt.Println("flagArray first ", flagArray)

	if len(flagArray) > 0 {
		flagSet := flag.NewFlagSet(flagArray[0], flag.ExitOnError)
		so := flagSet.String("so", "", "second one")
		ss := flagSet.String("ss", "", "second two")
		st := flagSet.Bool("st", false, "second three")

		errPasreSecond := flagSet.Parse(flagArray[1:])
		if errPasreSecond != nil {
			fmt.Println("errParseSecond", errPasreSecond)
		}
		flagSecondArray := flagSet.Args()

		fmt.Println("flag set out put")
		flagSet.SetOutput(os.Stdout)

		fmt.Printf("flag set %#v\n", flagSet)

		fmt.Println("so", *so, "ss", *ss, "st", *st)
		fmt.Println("flagSecondArray", flagSecondArray)
	}

}
