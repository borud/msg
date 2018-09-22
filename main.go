package main

import (
	"log"
	"os"
	"os/user"

	"github.com/kevinburke/twilio-go"
)

var sid = os.Getenv("TWILIO_SID")
var token = os.Getenv("TWILIO_TOKEN")

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <phone number> <message>", os.Args[0])
	}

	if sid == "" || token == "" {
		log.Fatal("please make sure TWILIO_SID and TWILIO_TOKEN environment variables are set", os.Args[0])
	}

	toNumber := os.Args[1]
	message := os.Args[2]

	from, err := user.Current()
	if err != nil {
		log.Fatal("Unable to determine current user: ", err)
	}

	client := twilio.NewClient(sid, token, nil)
	msg, err := client.Messages.SendMessage(from.Username, toNumber, message, nil)
	if err != nil {
		log.Fatal("Error sending message: ", err)
	}

	log.Print("Message sent: ", msg)
}
