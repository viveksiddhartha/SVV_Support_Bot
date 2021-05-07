package types

import (
	"errors"
)

// Commands is the command array
var Commands = map[string]string{
	"/start": "Start the bot help",
	"/news":  "Get the top headlines",
	// "/port":  "Get stock portfolio rates right now",
	"/word": "Get word of the day",
}

// BotName is the name of the bot
const BotName = "@fetchitemsbot"

// ParseCommand parses the string to get the appropriate command
func ParseCommand(userCommand string) (string, error) {
	for key := range Commands {
		if len(userCommand) >= len(key) {
			if userCommand[:len(key)] == key {
				return userCommand[:len(key)], nil
			}
		}
	}

	return "", errors.New("Invalid Command")
}

// ParseArguments parses the string to get the appropriate arguments
func ParseArguments(userCommand string) string {
	for key := range Commands {
		if len(userCommand) >= len(key) {
			if userCommand[:len(key)] == key {
				userCommand = userCommand[len(key):]
			}
		}
	}

	if len(userCommand) >= len(BotName) {
		if userCommand[:len(BotName)] == BotName {
			userCommand = userCommand[len(BotName):]
		}
	}

	return userCommand
}
