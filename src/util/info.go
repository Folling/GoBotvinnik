package util

import (
    "github.com/bwmarrin/discordgo"
    "strconv"
)

func PrintHelp() string {
    return "Help function currently not implemented\n"
}

func ChannelInfo(channel *discordgo.Channel) string {
    var ret string
    ret = " ChannelID    : " + channel.ID            + "\n" +
          " Name         : " + channel.Name          + "\n" +
          " Type         : " + channel.Type          + "\n" +
          " Last Message : " + channel.LastMessageID + "\n" +
          " Privacy      : " + strconv.FormatBool(channel.IsPrivate) + "\n\n"
    return ret
}

func GuildInfo(guild *discordgo.Guild) string {
    var ret string
    ret = " GuildID      : " + guild.ID                                             + "\n" +
          " Name         : " + guild.Name                                           + "\n" +
          " Region       : " + guild.Region                                         + "\n" +
          " Owner        : " + guild.OwnerID                                        + "\n" +
          " MemberCount  : " + strconv.FormatInt(int64(guild.MemberCount), 10) + "\n"
    return ret
}

