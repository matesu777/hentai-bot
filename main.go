package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	TOKEN := os.Getenv("TOKEN")
	disc, err := discordgo.New("Bot " + TOKEN)
	if err != nil{
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	fmt.Println("My bot running!!!!!")

}