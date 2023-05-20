package main

import (
	"context"
	"flag"
	"github.com/Syfaro/telegram-bot-api"
	"gopkg.in/yaml.v3"
	telegrampb "learnability_bot/pb"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var conf = flag.String("config", "config.yaml", "Configuration")

type ConfCourse struct {
	CourseID int64   `yaml:"id"`
	ChatID   []int64 `yaml:"chats"`
}

type BotService struct {
	Courses  []ConfCourse `yaml:"course"`
	BotToken string
	bot      *tgbotapi.BotAPI
	ctx      context.Context
}

func (c *BotService) getConf() *BotService {

	yamlFile, err := os.ReadFile(*conf)

	if err != nil {
		log.Panic(err)
	}

	if err = yaml.Unmarshal(yamlFile, &c); err != nil {
		log.Panic(err)
	} else {
		log.Printf("File read")
	}

	log.Printf("file is ", c.Courses)

	return c
}

func notify(courseId int64, course string, name string, typeElement telegrampb.ElementType, typeAction telegrampb.ActionType) error {

	flag.Parse()
	// подключаемся к боту с помощью токена

	svcContext, svcCancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	defer svcCancel()
	svc := new(BotService)
	svc.getConf()
	svc.BotToken = os.Getenv("botkey")
	svc.ctx = svcContext
	if err := svc.start(); err != nil {
		log.Panic("Unable to start bot", err)
		return err
	}

	var action string
	switch typeAction {
	case telegrampb.ActionType_create:
		action = "Добавлен"
	case telegrampb.ActionType_update:
		action = "Обновлён"
	case telegrampb.ActionType_delete:
		action = "Удалён"
	}

	var element string
	switch typeElement {
	case telegrampb.ElementType_topic:
		element = "тема"
	case telegrampb.ElementType_test:
		element = "тест"
	case telegrampb.ElementType_lecture:
		element = "лекция"
	}

	for _, c := range svc.Courses {
		log.Printf("ID course", c.CourseID)
		if courseId == c.CourseID {
			for _, chat := range c.ChatID {
				msg := tgbotapi.NewMessage(chat, "В курсе '"+course+"' произошли небольшие изменения.\n"+action+" элемент '"+name+"' ("+element+").")
				if _, err := svc.bot.Send(msg); err != nil {
					log.Panic("Unable to send", err)
					return err
				}
			}
		}
	}

	return nil
}

func (c *BotService) start() error {
	if bot, err := tgbotapi.NewBotAPI(c.BotToken); err != nil {
		return err
	} else {
		bot.Debug = true
		c.bot = bot
		log.Printf("Authorized on account", bot.Self.UserName)
		var ucfg = tgbotapi.NewUpdate(0)
		ucfg.Timeout = 60
	}
	return nil
}
