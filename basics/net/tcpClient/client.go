package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Подключаемся к сокету
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err.Error())
		}
		// Отправляем в socket

		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Print("Message from server: " + message)
	}
}
