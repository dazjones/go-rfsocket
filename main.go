package main

import (
	"bufio"
	"flag"
	"log"
	"os/exec"
)

func main() {
	sender_bin := flag.String("sender", "", "Path to sender")
	receiver_bin := flag.String("receiver", "", "Path to receiver")

	flag.Parse()

	if *sender_bin == "" {
		log.Fatalf("Flag sender is empty")
	}

	if *receiver_bin == "" {
		//log.Fatalf("Flag receiver is empty")
	}

	receiver := make(chan string)

	go func() {
		cmd := exec.Command(*sender_bin)
		out, _ := cmd.StdoutPipe()

		cmd.Start()

		scanner := bufio.NewScanner(out)

		for scanner.Scan() {
			text := scanner.Text()
			receiver <- text
		}
	}()

	log.Println("Ready")

	for {
		received_code := <-receiver
		log.Print(received_code)
	}

}
