package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func write() {
	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln("f", err.Error())
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1; i++ {
		str := fmt.Sprintf("%d\n", rand.Int())
		_, err := f.WriteString(str)
		fmt.Println(str)

		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
