package main

import (
	"flag"
	"log"
	"os/exec"
	"bufio"
	"strings"
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

	go func () {
		cmd := exec.Command(*sender_bin)
		out, _ := cmd.StdoutPipe()

		cmd.Start()

		for {
			r := bufio.NewReader(out)
			line, err := r.ReadBytes('\n')

			if err != nil {
				log.Fatalf("Error reading bytes")
			}

			s := string(line)
			strings.Replace(s, "\n", "", -1)

			receiver <- s
		}

	}()

	log.Println("Ready")

	for {
		received_code := <- receiver
		log.Print(received_code)
	}

}
