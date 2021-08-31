package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	dbUser := os.Getenv("WZRD_DB_USER")
	dbPass := os.Getenv("WZRD_DB_PASS")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.oqon8.mongodb.net/WzrdBot?retryWrites=true&w=majority", dbUser, dbPass)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

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
	} else {
		fmt.Println("Message info: ", m.Author.ID)
		fmt.Println("Message info: ", m.Author.Username)
	}

	if strings.Contains(m.Content, "nut") {
		s.ChannelMessageSend(m.ChannelID, "Do you have a nut allergy? \n"+
			"https://www.betterhealth.vic.gov.au/health/conditionsandtreatments/nut-allergies")
	}

	if strings.HasPrefix(m.Content, "wzrd") {
		// s.ChannelMessageSend(m.ChannelID, wzrdMsgHandler(strings.TrimSpace(m.Content)))
		s.ChannelMessageSend(m.ChannelID, wzrdMsgHandler(m, s))
	}
}

func wzrdMsgHandler(m *discordgo.MessageCreate, s *discordgo.Session) string {

	var msg = strings.TrimSpace(m.Content)

	switch msg {
	case "wzrd start":
		fmt.Println("the guild id: ", m.GuildID)
		createPlayers(m.GuildID, s)
		return "Welcome you little wizards and wizardesses! \n " +
			"First things first, you are straight up just a turkey right now. \n" +
			"If you are able to return to your true form then you can cast spells and shit. \n" +
			"Here are some commands to get you started: \n\n" + turkeyCommands
	case "wzrd gobble":
		return "You emit a creepy little gobble. You little freak. Slutty little turkey"
	case "wzrd peck":
		return "One day you will be able to peck someone. But for now you just smash your beak on the ground"
	default:
		return "Invalid wzrd command. Do better plz"
	}

}

func createPlayers(guildId string, s *discordgo.Session) {
	// fmt.Println("I need to create characters for these fools ", s.GuildMembers(guildId))
	// var guildMembers = s.GuildMembers(guildId, "", 100)
	guildMembers, err := s.GuildMembers(guildId, "", 100)
	if err != nil {
		fmt.Println("error getting guild memebers,", err)
		return
	}
	fmt.Println("Guild Members: ", guildMembers)
	for i, mem := range guildMembers {
		fmt.Println(i, mem)
		// TODO create all the users in the data store. add any info necessary
		// wat da heck is mongo?

		// I don't think we need this junk but it is a helpful example
		// userInfo, err := s.User(mem.User.ID)
		// if err != nil {
		// 	fmt.Println("error getting user info", err)
		// }
		// fmt.Println(userInfo)
	}
}
