package models

const (
	ProviderOpenRouter ModelProvider = "openrouter"

	// Anthropic
	OpenRouterClaudeFable51  ModelID = "openrouter.anthropic/claude-fable-5.1"
	OpenRouterClaudeOpus55   ModelID = "openrouter.anthropic/claude-opus-5.5"
	OpenRouterClaudeSonnet55 ModelID = "openrouter.anthropic/claude-sonnet-5.5"
	OpenRouterClaudeSonnet5  ModelID = "openrouter.anthropic/claude-sonnet-5"
	OpenRouterClaudeSonnet46 ModelID = "openrouter.anthropic/claude-sonnet-4.6"
	OpenRouterClaudeHaiku45  ModelID = "openrouter.anthropic/claude-haiku-4.5"

	// OpenAI
	OpenRouterGPT61Sol  ModelID = "openrouter.openai/gpt-6.1-sol"
	OpenRouterGPT6Luna  ModelID = "openrouter.openai/gpt-6-luna"
	OpenRouterGPT56Sol  ModelID = "openrouter.openai/gpt-5.6-sol"
	OpenRouterGPT56Luna ModelID = "openrouter.openai/gpt-5.6-luna"
	OpenRouterGPT5      ModelID = "openrouter.openai/gpt-5"
	OpenRouterGPT41     ModelID = "openrouter.openai/gpt-4.1"
	OpenRouterGPT41Mini ModelID = "openrouter.openai/gpt-4.1-mini"

	// Google
	OpenRouterGemini31Pro   ModelID = "openrouter.google/gemini-3.1-pro-preview"
	OpenRouterGemini38Flash ModelID = "openrouter.google/gemini-3.8-flash"
	OpenRouterGemini25Pro   ModelID = "openrouter.google/gemini-2.5-pro"

	// Others
	OpenRouterGrok47      ModelID = "openrouter.x-ai/grok-4.7"
	OpenRouterDeepSeekV32 ModelID = "openrouter.deepseek/deepseek-v3.2"
)

// sourceModel looks a model up in the provider catalogs it belongs to. The
// provider maps are read directly instead of SupportedModels because package
// level variables are initialized before init() merges them together.
func sourceModel(id ModelID) Model {
	for _, catalog := range []map[ModelID]Model{
		AnthropicModels, OpenAIModels, GeminiModels, XAIModels,
	} {
		if m, ok := catalog[id]; ok {
			return m
		}
	}
	return Model{}
}

// openRouterModel mirrors a model from another provider through OpenRouter.
// Context window, output cap and capabilities are inherited; OpenRouter bills
// at the upstream provider's rates.
func openRouterModel(id ModelID, base ModelID, display string) Model {
	src := sourceModel(base)
	src.ID = id
	src.Name = display
	src.Provider = ProviderOpenRouter
	return src
}

// https://openrouter.ai/models
var OpenRouterModels = map[ModelID]Model{
	OpenRouterClaudeFable51: openRouterModel(OpenRouterClaudeFable51, ClaudeFable51,
		"OpenRouter – Claude Fable 5.1"),
	OpenRouterClaudeOpus55: openRouterModel(OpenRouterClaudeOpus55, ClaudeOpus55,
		"OpenRouter – Claude Opus 5.5"),
	OpenRouterClaudeSonnet55: openRouterModel(OpenRouterClaudeSonnet55, ClaudeSonnet55,
		"OpenRouter – Claude Sonnet 5.5"),
	OpenRouterClaudeSonnet5: openRouterModel(OpenRouterClaudeSonnet5, ClaudeSonnet5,
		"OpenRouter – Claude Sonnet 5"),
	OpenRouterClaudeSonnet46: openRouterModel(OpenRouterClaudeSonnet46, ClaudeSonnet46,
		"OpenRouter – Claude Sonnet 4.6"),
	OpenRouterClaudeHaiku45: openRouterModel(OpenRouterClaudeHaiku45, ClaudeHaiku45,
		"OpenRouter – Claude Haiku 4.5"),

	OpenRouterGPT61Sol:  openRouterModel(OpenRouterGPT61Sol, GPT61Sol, "OpenRouter – GPT-6.1 Sol"),
	OpenRouterGPT6Luna:  openRouterModel(OpenRouterGPT6Luna, GPT6Luna, "OpenRouter – GPT-6 Luna"),
	OpenRouterGPT56Sol:  openRouterModel(OpenRouterGPT56Sol, GPT56Sol, "OpenRouter – GPT-5.6 Sol"),
	OpenRouterGPT56Luna: openRouterModel(OpenRouterGPT56Luna, GPT56Luna, "OpenRouter – GPT-5.6 Luna"),
	OpenRouterGPT5:      openRouterModel(OpenRouterGPT5, GPT5, "OpenRouter – GPT-5"),
	OpenRouterGPT41:     openRouterModel(OpenRouterGPT41, GPT41, "OpenRouter – GPT 4.1"),
	OpenRouterGPT41Mini: openRouterModel(OpenRouterGPT41Mini, GPT41Mini, "OpenRouter – GPT 4.1 mini"),

	OpenRouterGemini31Pro: openRouterModel(OpenRouterGemini31Pro, Gemini31Pro,
		"OpenRouter – Gemini 3.1 Pro"),
	OpenRouterGemini38Flash: openRouterModel(OpenRouterGemini38Flash, Gemini38Flash,
		"OpenRouter – Gemini 3.8 Flash"),
	OpenRouterGemini25Pro: openRouterModel(OpenRouterGemini25Pro, Gemini25Pro,
		"OpenRouter – Gemini 2.5 Pro"),

	OpenRouterGrok47: openRouterModel(OpenRouterGrok47, XAIGrok47,
		"OpenRouter – Grok 4.7"),
	OpenRouterDeepSeekV32: {
		ID:                  OpenRouterDeepSeekV32,
		Name:                "OpenRouter – DeepSeek V3.2",
		Provider:            ProviderOpenRouter,
		APIModel:            "deepseek/deepseek-v3.2",
		CostPer1MIn:         0.28,
		CostPer1MOutCached:  0.03,
		CostPer1MOut:        0.42,
		ContextWindow:       163_840,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: false,
	},
}
