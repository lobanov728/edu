package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Запускаем цикл
	for {
		fmt.Println(1)
		// Открываем порт
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Printf("conn %+v\n", conn)
		// Будем прослушивать все сообщения разделенные \n
		go func() {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalln(err.Error())
			}

			fmt.Println(2)
			// Распечатываем полученое сообщение
			fmt.Print("Message Received:", string(message))
			time.Sleep(time.Second * 10)
			// Процесс выборки для полученной строки
			newmessage := strings.ToUpper(message)
			// Отправить новую строку обратно клиенту
			conn.Write([]byte(newmessage + "\n"))
		}()
	}
}
