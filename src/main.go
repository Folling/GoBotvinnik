package main

import (
    "os"
    "os/signal"
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    "GoBotvinnik/src/launch"
)


func init() {
    raven.SetDSN("ravenDSN")
}

func main() {
    discord, err := discordgo.New("Bot " + "token")
    if err != nil {
        raven.CaptureError(err, nil)
    }
    discord.AddHandler(launch.OnReady)
    discord.AddHandler(launch.OnMessage)

    err = discord.Open()
    if err != nil {
        raven.CaptureError(err, nil)
    }

    channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt, os.Kill)

    // Wait until the os wants us to shutdown
    <-channel

    discord.Close()
}
