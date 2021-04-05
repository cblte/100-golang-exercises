package main

import (
	"fmt"
)

func main() {

	want := "HELLO WORLD!"

	ReadString()
	got := PrintString()
	//got = "afdsas"

	if got != want {
		fmt.Println(fmt.Errorf("Error: Ex005() = %v, want %v", got, want))
		return
	}

	fmt.Println(got)

}
