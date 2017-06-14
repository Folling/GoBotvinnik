package launch

import (
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    "strings"
    "GoBotvinnik/src/commands"
)


func OnReady(session *discordgo.Session, event *discordgo.Ready){
    commands.Initialize()
    channel, err := session.UserChannelCreate("135818878176460801")
    if err != nil {
        raven.CaptureError(err, nil)
    }
    session.ChannelMessageSend(channel.ID, "Bot started")
}

func OnMessage(session *discordgo.Session, mesg *discordgo.MessageCreate) {

    if mesg.Author.Bot || mesg.MentionEveryone {
        return
    }
    channel, err := session.Channel(mesg.ChannelID)
    if err != nil {
        raven.CaptureError(err, nil)
    }
    if channel.IsPrivate {
        return
    }
    msg := mesg.Content

    if !strings.HasPrefix(msg, commands.Prefix) {
        return
    }
    subs := strings.Split(msg, " ")

    command := strings.TrimPrefix(subs[0], commands.Prefix)
    content := strings.Join(subs[1:], " ")
    if _, ok := commands.Commands[command]; !ok {
        return
    }
    commands.Commands[command](command, content, mesg.Message, session)

}
