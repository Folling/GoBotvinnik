package commands

import "github.com/bwmarrin/discordgo"

func Calculate(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelMessageSend(msg.ChannelID, "44")
}
