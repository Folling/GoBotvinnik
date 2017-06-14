package errorLog

import (
    "github.com/bwmarrin/discordgo"
    "time"
)

type UserMessageError struct {
    User *discordgo.User
    Time time.Time
}

type ChannelCreateError struct {
    Channel *discordgo.Channel
    Time time.Time
}