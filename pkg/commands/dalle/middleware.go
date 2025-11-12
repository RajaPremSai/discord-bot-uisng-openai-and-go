package dalle

import (
	"context"
	"log"

	"github.com/RajaPremSai/go-openai-dicord-bot/pkg/bot"
	discord "github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
)

func imageInteractionResponseMiddleware(ctx *bot.Context) {
	log.Printf("[GID:%s,i.ID:%s] Image interaction invoked by UserID: %s\n", ctx.Interaction.GuildID, ctx.Interaction.ID)

	err := ctx.Respond(&discord.InteractionResponse{
		Type: discord.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("[GID:%s,i.ID:%s] Failed to respond to interation with the error %v", ctx.Interaction.GuildID, ctx.Interaction.ID, err)
		return
	}

	ctx.Next()
}
func imageModerationMiddleware(ctx *bot.Context, client *openai.Client) {
	log.Printf("[GId : %s,i.ID:%s] Performing interaction moderation middeware\n", ctx.Interaction.GuildID, ctx.Interaction.ID)

	var prompt string
	if option, ok := ctx.Options[imageCommandOptionPrompt.String()]; ok {
		prompt = option.StringValue()
	} else {
		log.Printf("[GID : %s,i.ID:%s] Failed to parse prompt option\n", ctx.Interaction.GuildID, ctx.Interaction.ID)
		ctx.Respond(&discord.InteractionResponse{
			Type: discord.InteractionResponseChannelMessageWithSource,
			Data: &discord.InteractionResponseData{
				Content: "ERROR : Failed to parse prompt option",
			},
		})
		return
	}
	resp, err := client.Moderations(
		context.Background(),
		openai.ModerationRequest{
			Input: prompt,
		},
	)
	if err != nil {
		log.Printf("[GID: %s, i.ID:%s] OPENAI Moderation API request failed with the error:%v\n", ctx.Interaction.GuildID, ctx.Interaction.ID, err)
		ctx.Next()
		return
	}

	if resp.Results[0].Flagged {
		log.Printf("[GID: %s, i.ID: %s] Ineraction was flagged y Moderation API,prompt: \"%s\"n", ctx.Interaction.GuildID, ctx.Interaction.ID, prompt)
		ctx.FollowupMessageCreate(ctx.Interaction, true, &discord.WebhookParams{
			Embeds: []*discord.MessageEmbed{
				{
					Title:       "❌ Error",
					Description: "The provided prompt contains text that violates OpenAI's usage policies and is not allowed by their safety system",
					Color:       0xff0000,
				},
			},
		})
		return
	}
	ctx.Next()
}
