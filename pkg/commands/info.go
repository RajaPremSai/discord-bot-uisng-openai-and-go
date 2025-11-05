package commands

import (
	"github.com/RajaPremSai/go-openai-dicord-bot/pkg/bot"
	"github.com/RajaPremSai/go-openai-dicord-bot/pkg/constants"
	discord "github.com/bwmarrin/discordgo"
)

const (
	infoCommandName = "info"
)

func infoHandler(ctx *bot.Context) {
	ctx.Respond(&discord.InteractionResponse{
		Type: discord.InteractionResponseChannelMessageWithSource,
		Data: &discord.InteractionResponseData{
			// Note: only visible to the user who invoked the command
			Flags: discord.MessageFlagsEphemeral,
			// Content: "Surprise!",
			Components: []discord.MessageComponent{
				discord.ActionsRow{
					Components: []discord.MessageComponent{
						&discord.Button{
							Label: "Source code",
							Style: discord.LinkButton,
							URL:   "https://github.com/RajaPremSai/go-openai-dicord-bot",
						},
					},
				},
			},
			Embeds: []*discord.MessageEmbed{
				{
					Title:       "Bot Version",
					Description: "Version: " + constants.Version,
					Color:       0x00bfff,
				},
			},
		},
	})
}

func InfoCommand() *bot.Command {
	return &bot.Command{
		Name:                     infoCommandName,
		Description:              "Show information about current version of Open AI bot",
		DMPermission:             true,
		DefaultMemberPermissions: discord.PermissionViewChannel,
		Handler:                  bot.HandlerFunc(infoHandler),
	}
}
