package constant

import (
	"calc/src/commands"
)

type CommandsType map[string]func()

var Commands = CommandsType{
	"exit": commands.Exit,
}
