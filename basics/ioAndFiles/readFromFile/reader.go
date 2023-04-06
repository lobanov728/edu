package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	read()

}

func read() {
	f1, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f1.Close()
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		if scanner.Text() == "7414497771202375935" {
			fmt.Println("found")
		}
	}

	f2, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f2.Close()

	b := make([]byte, 1)

	for {
		_, err := f2.Read(b)
		if err == io.EOF {
			break
		}
	}

	ff, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer ff.Close()
	bb, err := ioutil.ReadAll(ff)

	fmt.Println(err)
	fmt.Println(bb)
}
