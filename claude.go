package main

import (
	"context"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type Claude struct {
	APIKey string
	Model  anthropic.Model
	client *anthropic.Client
}

func (c *Claude) Prompt(in string) (out string, err error) {
	if c.client == nil {
		client := anthropic.NewClient(option.WithAPIKey(c.APIKey))
		c.client = &client
	}
	if c.Model == "" {
		c.Model = anthropic.ModelClaude3_7SonnetLatest
	}

	m, err := c.client.Messages.New(context.Background(), anthropic.MessageNewParams{
		MaxTokens: 4096,
		Messages: []anthropic.MessageParam{{
			Role: anthropic.MessageParamRoleUser,
			Content: []anthropic.ContentBlockParamUnion{{
				OfRequestTextBlock: &anthropic.TextBlockParam{Text: in},
			}},
		}},
		Model: c.Model,
	})
	if err != nil {
		return
	}

	ss := make([]string, len(m.Content))
	for i, c := range m.Content {
		ss[i] = c.Text
	}
	out = strings.Join(ss, "")
	return
}

var _ LLM = &Claude{}
