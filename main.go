package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	//TODO 
	if len(os.Args) < 2{
		log.Fatal("usage: ingrecognintion<img_url>")
	}
	fmt.Printf("url:%s",os.Args[1])
}