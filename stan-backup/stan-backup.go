package main

import (
	"fmt"
	"os"
)


func main() {

	fmt.Println("Backup")

	args := os.Args[1:]

	subject := args[0]
	bak := args[1]
	idx := args[2]

	fmt.Printf("Sub: %s\nBak: %s\nIdx: %s\n",subject,bak,idx)

	err := os.Rename( "/Users/thomas/tutorial/stan/" + bak, "/Users/thomas/tutorial/stan/backup/" + bak)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Rename( "/Users/thomas/tutorial/stan/" + idx, "/Users/thomas/tutorial/stan/backup/" + idx)
	if err != nil {
		fmt.Println(err)
	}

}
