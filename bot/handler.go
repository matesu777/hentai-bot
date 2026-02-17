package bot

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/matesu777/DiscordBot-Go/media"
)

var botprefix string = "$"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID{
		return
	}
	switch m.Content{
		case botprefix + "hentai":{
			rand.Seed(time.Now().UnixNano())
			hentai, err := media.RandomImage("media/Hentai")
			file, err := os.Open(hentai)

			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Erro ao abrir a imagem!")
				return
			}
			
			defer file.Close()
			
			fmt.Println("imagem enviada: ", hentai)

			_, err = s.ChannelFileSend(m.ChannelID, "hentai.jpg" ,file)
			
			fmt.Println("Arquivo enviado!")

			if err != nil {
            	fmt.Println(err)
        	}
			fmt.Println("Arquivo enviado!")

			return
		}
		case botprefix + "blowjob":{
			res, err := media.GetApi("nsfw", "blowjob")
			if err != nil{
				s.ChannelMessageSend(m.ChannelID, "Ao achar imagem!")
				fmt.Println(err)
				return
			}
	
			_, err = s.ChannelMessageSend(m.ChannelID, res.Url)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println("Arquivo enviado!")

			return
		}
		case botprefix + "trap":{
			res, err := media.GetApi("nsfw", "trap")
			if err != nil{
				s.ChannelMessageSend(m.ChannelID, "Ao achar imagem!")
				fmt.Println(err)
				return
			}
			_, err = s.ChannelMessageSend(m.ChannelID, res.Url)
			fmt.Println("imagem enviada: ", res.Url)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println("Arquivo enviado!")

			return
		}
		case botprefix + "neko":{
			res, err := media.GetApi("nsfw", "neko")
			if err != nil{
				s.ChannelMessageSend(m.ChannelID, "Ao achar imagem!")
				fmt.Println(err)
				return
			}
	
			_, err = s.ChannelMessageSend(m.ChannelID, res.Url)
			fmt.Println("imagem enviada: ", res.Url)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println("Arquivo enviado!")

			return
		}
		case botprefix + "waifu":{
			res, err := media.GetApi("nsfw", "waifu")

			if err != nil{
				s.ChannelMessageSend(m.ChannelID, "Ao achar imagem!")
				fmt.Println(err)
				return
			}

			_, err = s.ChannelMessageSend(m.ChannelID, res.Url)
			fmt.Println("imagem enviada: ", res.Url)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println("Arquivo enviado!")


			return
		}
	}
}