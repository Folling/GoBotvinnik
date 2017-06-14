package errorLog

import (
    "github.com/bwmarrin/discordgo"
    "GoBotvinnik/src/util"
)

func Send(s *discordgo.Session, err *ChannelCreateError, msg string) {
    s.ChannelMessageSend("323055650424881152" ,
                         util.ChannelInfo(err.Channel) +
                         "\n\n" + msg )
}