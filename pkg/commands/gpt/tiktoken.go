package gpt

import (
	"github.com/sashabaranov/go-openai"
	"github.com/tiktoken-go/tokenizer"
)

func countMessageTokens(message openai.ChatCompletionMessage, model string) *int {
	ok, tokensPerMessage, tokensPerName := _tokensConfiguration(model)
	if !ok {
		return nil
	}
	enc, err := tokenizer.ForModel(tokenizer.Model(model))
	if err != nil {
		enc, _ = tokenizer.Get(tokenizer.Cl100kBase)
	}
	tokens := _countMessageTokens(enc, tokensPerMessage, tokensPerName, message)
	return &tokens
}

func _countMessageTokens(enc tokenizer.Codec, tokensPerMessage int, tokensPerName int, message openai.ChatCompletionMessage) int {
	tokens := tokensPerMessage
	contentIds, _, _ := enc.Encode(message.Content)
	roleIds, _, _ := enc.Encode(message.Role)
	tokens += len(contentIds)
	tokens += len(roleIds)
	if message.Name != "" {
		tokens += tokensPerName
		nameIds, _, _ := enc.Encode(message.Name)
		tokens += len(nameIds)
	}
	return tokens
}
