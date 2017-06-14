package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    //"ChessBot/src/util"
    "regexp"
    "strings"
    "ChessBot/src/commands"
)

var (
    Prefix = ">"
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
    if strings.HasPrefix(msg, Prefix) {
        subs := strings.Split(msg, " ")
        if len(subs) == 0 {
            return
        }
        bmsg := []byte(msg)

        switch {
        case regexp.MustCompile(`[a-z]\w+`).Match(bmsg):
            command := strings.TrimPrefix(subs[0], Prefix)
            content := strings.Join(subs[1:], " ")
            if content == "" {
                return
            }
            if _, ok := commands.Commands[command]; !ok {
                session.ChannelMessageSend(channel.ID, "Unkown command :confused:")
                return
            }
            commands.Commands[command](command, content, mesg.Message, session)
        }
    }

}
