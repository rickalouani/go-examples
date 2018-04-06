package main

import (
	"fmt"
//	"errors"
//	"io/ioutil"
)

func main() {
/*	var file []byte
	var err error
	file, err = ioutil.ReadFile("src/foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", file) 
*/
/*

        err := errors.New("Something went wrong")
        if err != nil {
         fmt.Println(err)
        }
*/
        name, role := "Richard Jupp", "Drummer"
        err := fmt.Errorf("The %v %v quit", role, name)
       if err != nil {
                fmt.Println(err)
        }

}


