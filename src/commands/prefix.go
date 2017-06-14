package commands

import (
    "github.com/bwmarrin/discordgo"
)

func GetPrefix(command string, content string, msg *discordgo.Message, session *discordgo.Session){
    session.ChannelMessageSend(msg.ChannelID, "Current prefix is \"" + Prefix + "\"")
}
