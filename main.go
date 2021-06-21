package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {

	banner :=
		`
	___ _____ ___ _                   
	| __|_   _| _ \ |  _  _ ______ _ _ 
	| _|  | | |  _/ |_| || |_ / -_) '_|
	|_|   |_| |_| |____\_, /__\___|_|  
	                    |__/            
   
   
	FTP User Enumeration tool.
	`

	fmt.Println(banner)

	var connHost, connPort, connType, wordList string
	var concureny int
	var findAll = false
	var wg sync.WaitGroup

	flag.StringVar(&connHost, "h", "127.0.0.1", "Enter hostname")
	flag.StringVar(&connPort, "p", "21", "Enter port")
	flag.StringVar(&wordList, "w", "", "Wordlist of usernames")
	flag.IntVar(&concureny, "c", 10, "Concurrency default")
	flag.BoolVar(&findAll, "continue", false, "Will Continue search for username upon first finding ")
	connType = "tcp"

	flag.Parse()

	file, err := os.Open(wordList)
	if err != nil {
		fmt.Println("Please make sure file exist")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	usernames := make(chan string)

	for i := 0; i < concureny; i++ {
		wg.Add(1)
		go func() {
			for user := range usernames {
				fmt.Printf("\rTesting username: %s", user)
				conn, err := net.Dial(connType, connHost+":"+connPort)
				if err != nil {
					fmt.Println("Error connecting:", err.Error())
					os.Exit(1)
				}

				reader := bufio.NewReader(conn)

				for {
					message, _ := reader.ReadString('\n')
					status_code := message[0:3]
					if status_code == "220" {
						conn.Write([]byte("user " + user + "\r\n"))
					} else if status_code == "331" {
						fmt.Printf("\n\nFOUND : %s\n\n", user)
						if !findAll {
							os.Exit(0)
						}
						break
					} else if status_code == "530" {
						break
					} else {
						fmt.Println(message)
					}
				}
				conn.Close()
			}
			wg.Done()
		}()
	}

	fmt.Printf("Attacking %s on port %s \n", connHost, connPort)
	for scanner.Scan() {
		user := scanner.Text()
		usernames <- user
	}
	close(usernames)
	wg.Wait()
}
