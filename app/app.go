package app

import (
	"context"

	"github.com/simpleittools/go-rag/chat"
	"github.com/simpleittools/go-rag/config"
	"github.com/simpleittools/go-rag/llm"
)

func Run(ctx context.Context, cfg config.Config) error {
	client := llm.New(cfg)
	return chat.RunREPL(ctx, client, chat.Options{
		SystemPromptFile: cfg.SystemPromptFile,
	})
}
