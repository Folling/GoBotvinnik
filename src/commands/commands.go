package commands

import "github.com/bwmarrin/discordgo"

type commandFunc func(command string, content string, msg *discordgo.Message, session *discordgo.Session)

var (
    Commands = make(map[string]commandFunc)
    Prefix = ">:"
)

func Initialize() {
    Commands["calculate"] = Calculate
    Commands["prefix"] = GetPrefix
}
