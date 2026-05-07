package main

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/richxcame/gosms/src/sms"
)

func main() {
	log.Fatal(sms.Run(context.Background()))
}
