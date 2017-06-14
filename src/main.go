package main

import (
    "os"
    "os/signal"
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    "GoBotvinnik/src/launch"
)


func init() {
    raven.SetDSN("https://43ec317525d64825a5b799abb3a46e27:4bca63ea5ed04409bddce476971cbc25@sentry.io/178389")
}

func main() {
    discord, err := discordgo.New("Bot " + "Mjk5MTg4ODgwMzk2MTI0MTYw.DBwKBQ.qp8cMI_lK3dYw4GyIh0c-04L5fU")
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