package models

const (
	ProviderXAI ModelProvider = "xai"

	XAIGrok47 ModelID = "grok-4.7"
	XAIGrok46 ModelID = "grok-4.6"
	XAIGrok45 ModelID = "grok-4.5"
	XAIGrok43 ModelID = "grok-4.3"
)

// https://docs.x.ai/docs/models
//
// CanReason stays false so the OpenAI-compatible client does not send
// reasoning_effort/max_completion_tokens parameters that the xAI API rejects.
var XAIModels = map[ModelID]Model{
	XAIGrok47: {
		ID:                  XAIGrok47,
		Name:                "Grok 4.7",
		Provider:            ProviderXAI,
		APIModel:            "grok-4.7",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOut:        6.00,
		CostPer1MOutCached:  0,
		ContextWindow:       500_000,
		DefaultMaxTokens:    50_000,
		SupportsAttachments: true,
	},
	XAIGrok46: {
		ID:                  XAIGrok46,
		Name:                "Grok 4.6",
		Provider:            ProviderXAI,
		APIModel:            "grok-4.6",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOut:        6.00,
		CostPer1MOutCached:  0,
		ContextWindow:       500_000,
		DefaultMaxTokens:    50_000,
		SupportsAttachments: true,
	},
	XAIGrok45: {
		ID:                  XAIGrok45,
		Name:                "Grok 4.5",
		Provider:            ProviderXAI,
		APIModel:            "grok-4.5",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOut:        6.00,
		CostPer1MOutCached:  0,
		ContextWindow:       500_000,
		DefaultMaxTokens:    50_000,
		SupportsAttachments: true,
	},
	XAIGrok43: {
		ID:                  XAIGrok43,
		Name:                "Grok 4.3",
		Provider:            ProviderXAI,
		APIModel:            "grok-4.3",
		CostPer1MIn:         1.25,
		CostPer1MInCached:   0,
		CostPer1MOut:        2.50,
		CostPer1MOutCached:  0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    30_000,
		SupportsAttachments: true,
	},
}
