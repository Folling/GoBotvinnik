package commands

import "github.com/bwmarrin/discordgo"

var (
    Commands = make(map[string]func(command string, content string, msg *discordgo.Message, session *discordgo.Session))
)

func Initialize() {
    Commands["calculate"] = Calculate
}