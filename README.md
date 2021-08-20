# Go Discord Bot
This is a little project to practice golang and make a discord bot

## Project setup
From inside the project directory run this to install all required modules
```
go get -u ./...
```
Then build the code before you can run it. (Also need to build after making any changes)
```
go build
```

### Run the server
Once you build the code you can run it with this command. Pass the discord token with the -t flag
```
./go-discord-bot -t <Token>
```

### TODO! If you want to help out here are some things that need to be done
1. Convert our janky ass discord text commands to more legit approach. Here is an [example](https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go)
2. Setup MongoDB to use as a datastore for the bot
3. Figure out a deployment strategy. Containerize? AWS? Heroku? Follow your heart
4. Functionality! Currently this bot doesn't really do anything. We need to write code...
