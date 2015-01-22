package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"gopkg.in/gomail.v1"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Server struct {
		TLS      bool   `json:"tls"`
		Host     string `json:"host"`
		Port     int  `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"server"`
	From struct {
		Name string `json:"name"`
		Address string `json:"address"`
	} `json:"from"`
	To []struct {
		Name string `json:"name"`
		Address string `json:"address"`
	} `json:"to"`
}

func main() {
	app := cli.NewApp()
	app.Name = "valiant"
	app.Usage = "A brave mail pigeon"
	app.Author = "Romain Baugue"
	app.Email = "romain.baugue@elwinar.com"
	app.Version = "1-dev"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "configuration,c",
			Value: "valiant.json",
			Usage: "configuration file",
		},
		cli.StringFlag{
			Name:  "body,b",
			Usage: "body file",
		},
		cli.StringFlag{
			Name: "subject,s",
			Usage: "message subject",
		},
	}
	app.Action = action

	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	var cfg Configuration
	var err error
	
	_, _, error := Bootstrap(ctx, &cfg)
	
	message := gomail.NewMessage()
	message.SetHeader("From", fmt.Sprintf("%s <%s>", cfg.From.Name, cfg.From.Address))
	
	var to []string
	for _, t := range cfg.To {
		to = append(to, fmt.Sprintf("%s <%s>", t.Name, t.Address))
	}
	message.SetHeader("To", to...)
	message.SetHeader("Subject", ctx.String("subject"))
	
	body, err := ioutil.ReadFile(ctx.String("body"))
	if err != nil {
		error.Fatalln("Unable to read body file:", err)
	}
	message.SetBody("text/html", string(body))
	
	mailer := gomail.NewMailer(cfg.Server.Host, cfg.Server.User, cfg.Server.Password, cfg.Server.Port)
	if err := mailer.Send(message); err != nil {
		error.Fatalln("Unable to send message:", err)
	}
}
