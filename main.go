package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

var turkeyCommands string = "wzrd gobble: make some weird turkey noises \n" +
	"wzrd peck: peck with that little beak of yours. aim for sean \n"

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provied bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "nut") {
		s.ChannelMessageSend(m.ChannelID, "Do you have a nut allergy? \n"+
			"https://www.betterhealth.vic.gov.au/health/conditionsandtreatments/nut-allergies")
	}

	// If the message is "fard" reply with "shid"
	if m.Content == "fard" {
		s.ChannelMessageSend(m.ChannelID, "shid")
	}

	// If the message is "shid" reply with "fard"
	if m.Content == "shid" {
		s.ChannelMessageSend(m.ChannelID, "fard")
	}

	if m.Content == "wzrd start" {
		s.ChannelMessageSend(m.ChannelID, "Welcome you little wizards and wizardesses! \n "+
			"First things first, you are straight up just a turkey right now. \n"+
			"If you are able to return to your true form then you can cast spells and shit. \n"+
			"Here are some commands to get you started: \n"+turkeyCommands)
	}
}
