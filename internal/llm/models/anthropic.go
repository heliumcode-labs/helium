package models

const (
	ProviderAnthropic ModelProvider = "anthropic"

	// Models
	ClaudeFable51  ModelID = "claude-fable-5.1"
	ClaudeFable5   ModelID = "claude-fable-5"
	ClaudeOpus55   ModelID = "claude-opus-5.5"
	ClaudeOpus5    ModelID = "claude-opus-5"
	ClaudeOpus48   ModelID = "claude-opus-4.8"
	ClaudeOpus45   ModelID = "claude-opus-4.5"
	ClaudeSonnet55 ModelID = "claude-sonnet-5.5"
	ClaudeSonnet5  ModelID = "claude-sonnet-5"
	ClaudeSonnet46 ModelID = "claude-sonnet-4.6"
	ClaudeSonnet45 ModelID = "claude-sonnet-4.5"
	ClaudeHaiku45  ModelID = "claude-haiku-4.5"
)

// https://platform.claude.com/docs/en/about-claude/models
//
// CostPer1MInCached is the 5-minute cache *write* price and CostPer1MOutCached
// is the cache *read* price, matching Anthropic's pricing table.
var AnthropicModels = map[ModelID]Model{
	ClaudeFable51: {
		ID:                  ClaudeFable51,
		Name:                "Claude Fable 5.1",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-fable-5-1",
		CostPer1MIn:         10.0,
		CostPer1MInCached:   12.50,
		CostPer1MOutCached:  0.25,
		CostPer1MOut:        50.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeFable5: {
		ID:                  ClaudeFable5,
		Name:                "Claude Fable 5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-fable-5",
		CostPer1MIn:         10.0,
		CostPer1MInCached:   12.50,
		CostPer1MOutCached:  1.00,
		CostPer1MOut:        50.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeOpus55: {
		ID:                  ClaudeOpus55,
		Name:                "Claude Opus 5.5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-opus-5-5",
		CostPer1MIn:         4.0,
		CostPer1MInCached:   5.0,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        20.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeOpus5: {
		ID:                  ClaudeOpus5,
		Name:                "Claude Opus 5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-opus-5",
		CostPer1MIn:         5.0,
		CostPer1MInCached:   6.25,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        25.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeOpus48: {
		ID:                  ClaudeOpus48,
		Name:                "Claude Opus 4.8",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-opus-4-8",
		CostPer1MIn:         5.0,
		CostPer1MInCached:   6.25,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        25.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeOpus45: {
		ID:                  ClaudeOpus45,
		Name:                "Claude Opus 4.5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-opus-4-5",
		CostPer1MIn:         5.0,
		CostPer1MInCached:   6.25,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        25.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    64_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeSonnet55: {
		ID:                  ClaudeSonnet55,
		Name:                "Claude Sonnet 5.5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-sonnet-5-5",
		CostPer1MIn:         2.0,
		CostPer1MInCached:   2.50,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        10.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeSonnet5: {
		ID:                  ClaudeSonnet5,
		Name:                "Claude Sonnet 5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-sonnet-5",
		CostPer1MIn:         2.0,
		CostPer1MInCached:   2.50,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        10.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeSonnet46: {
		ID:                  ClaudeSonnet46,
		Name:                "Claude Sonnet 4.6",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-sonnet-4-6",
		CostPer1MIn:         3.0,
		CostPer1MInCached:   3.75,
		CostPer1MOutCached:  0.30,
		CostPer1MOut:        15.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeSonnet45: {
		ID:                  ClaudeSonnet45,
		Name:                "Claude Sonnet 4.5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-sonnet-4-5",
		CostPer1MIn:         3.0,
		CostPer1MInCached:   3.75,
		CostPer1MOutCached:  0.30,
		CostPer1MOut:        15.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    64_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	ClaudeHaiku45: {
		ID:                  ClaudeHaiku45,
		Name:                "Claude Haiku 4.5",
		Provider:            ProviderAnthropic,
		APIModel:            "claude-haiku-4-5",
		CostPer1MIn:         1.0,
		CostPer1MInCached:   1.25,
		CostPer1MOutCached:  0.10,
		CostPer1MOut:        5.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    64_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
}
