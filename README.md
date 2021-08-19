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
