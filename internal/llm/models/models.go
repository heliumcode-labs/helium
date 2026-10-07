package models

import "maps"

type (
	ModelID       string
	ModelProvider string
)

type Model struct {
	ID                  ModelID       `json:"id"`
	Name                string        `json:"name"`
	Provider            ModelProvider `json:"provider"`
	APIModel            string        `json:"api_model"`
	CostPer1MIn         float64       `json:"cost_per_1m_in"`
	CostPer1MOut        float64       `json:"cost_per_1m_out"`
	CostPer1MInCached   float64       `json:"cost_per_1m_in_cached"`
	CostPer1MOutCached  float64       `json:"cost_per_1m_out_cached"`
	ContextWindow       int64         `json:"context_window"`
	DefaultMaxTokens    int64         `json:"default_max_tokens"`
	CanReason           bool          `json:"can_reason"`
	SupportsAttachments bool          `json:"supports_attachments"`
}

// Model IDs
const (
	// Bedrock
	BedrockClaudeFable51  ModelID = "bedrock.claude-fable-5.1"
	BedrockClaudeOpus55   ModelID = "bedrock.claude-opus-5.5"
	BedrockClaudeSonnet55 ModelID = "bedrock.claude-sonnet-5.5"
	BedrockClaudeSonnet46 ModelID = "bedrock.claude-sonnet-4.6"
	BedrockClaudeHaiku45  ModelID = "bedrock.claude-haiku-4.5"
)

const (
	ProviderBedrock ModelProvider = "bedrock"
	// ForTests
	ProviderMock ModelProvider = "__mock"
)

// Providers in order of popularity
var ProviderPopularity = map[ModelProvider]int{
	ProviderCopilot:    1,
	ProviderAnthropic:  2,
	ProviderOpenAI:     3,
	ProviderGemini:     4,
	ProviderGROQ:       5,
	ProviderOpenRouter: 6,
	ProviderXAI:        7,
	ProviderBedrock:    8,
	ProviderAzure:      9,
	ProviderVertexAI:   10,
}

var SupportedModels = map[ModelID]Model{
	// Bedrock model IDs are prefixed "bedrock." while the API model keeps the
	// "anthropic." prefix required by the Bedrock runtime.
	BedrockClaudeFable51: {
		ID:                  BedrockClaudeFable51,
		Name:                "Bedrock: Claude Fable 5.1",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-fable-5-1",
		CostPer1MIn:         10.0,
		CostPer1MInCached:   12.5,
		CostPer1MOutCached:  0.25,
		CostPer1MOut:        50.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaudeOpus55: {
		ID:                  BedrockClaudeOpus55,
		Name:                "Bedrock: Claude Opus 5.5",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-opus-5-5",
		CostPer1MIn:         4.0,
		CostPer1MInCached:   5.0,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        20.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaudeSonnet55: {
		ID:                  BedrockClaudeSonnet55,
		Name:                "Bedrock: Claude Sonnet 5.5",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-sonnet-5-5",
		CostPer1MIn:         2.0,
		CostPer1MInCached:   2.5,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        10.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaudeSonnet46: {
		ID:                  BedrockClaudeSonnet46,
		Name:                "Bedrock: Claude Sonnet 4.6",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-sonnet-4-6",
		CostPer1MIn:         3.0,
		CostPer1MInCached:   3.75,
		CostPer1MOutCached:  0.30,
		CostPer1MOut:        15.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaudeHaiku45: {
		ID:                  BedrockClaudeHaiku45,
		Name:                "Bedrock: Claude Haiku 4.5",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-haiku-4-5-20251001-v1:0",
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

func init() {
	maps.Copy(SupportedModels, AnthropicModels)
	maps.Copy(SupportedModels, OpenAIModels)
	maps.Copy(SupportedModels, GeminiModels)
	maps.Copy(SupportedModels, GroqModels)
	maps.Copy(SupportedModels, AzureModels)
	maps.Copy(SupportedModels, OpenRouterModels)
	maps.Copy(SupportedModels, XAIModels)
	maps.Copy(SupportedModels, VertexAIGeminiModels)
	maps.Copy(SupportedModels, CopilotModels)
}
