/*
 * @author:kiky
 * @date: 2021/12/02 10:42
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Print(os.Args)
	if(len(os.Args) > 1){
		fmt.Println("test",os.Args[1])
	}
	os.Exit(-1)
}
