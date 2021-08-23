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
Need to set the DB config variables as env vars
```
export WZRD_DB_USER=
export WZRD_DB_PASS=
export WZRD_DB_TABLE=
```

### Run the server
Once you build the code you can run it with this command. Pass the discord token with the -t flag
```
./go-discord-bot -t <Token>
```

### Clean up go mod
If you get errors around your go mod file like "missing go sum entry" you can run the following command:
```
go mod tidy
```
This command goes through the go.mod file to resolve dependencies, deletes packages that aren't needed, downloads necessary packages, and updates the go.sum file


### TODO! If you want to help out here are some things that need to be done
1. Convert our janky ass discord text commands to more legit approach. Here is an [example](https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go)
2. Setup MongoDB to use as a datastore for the bot
3. Figure out a deployment strategy. Containerize? AWS? Heroku? Follow your heart
4. Functionality! Currently this bot doesn't really do anything. We need to write code...
