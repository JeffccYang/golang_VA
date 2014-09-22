package main

import (
	"fmt"
	//"syscall"
	//"unsafe"
	///"va"
	//"web"
	"os/exec"
)

func main() {

	fmt.Printf("format")

	///	va.PrintVa()

	/*	c, err := exec.Command("gcc", "-v").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(c))
	*/
	/*
		c := exec.Command("cmd", "/C", "gcc", "-v")

		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	*/
	cmd := exec.Command("cmd", "/c", "gcc", "-v")
	out, _ := cmd.Output()
	fmt.Print(string(out))

}
