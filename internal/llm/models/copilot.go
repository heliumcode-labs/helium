package models

const (
	ProviderCopilot ModelProvider = "copilot"

	// OpenAI models
	CopilotGPT61Sol   ModelID = "copilot.gpt-6.1-sol"
	CopilotGPT6Astra  ModelID = "copilot.gpt-6-astra"
	CopilotGPT6Luna   ModelID = "copilot.gpt-6-luna"
	CopilotGPT56Sol   ModelID = "copilot.gpt-5.6-sol"
	CopilotGPT56Terra ModelID = "copilot.gpt-5.6-terra"
	CopilotGPT56Luna  ModelID = "copilot.gpt-5.6-luna"
	CopilotGPT55      ModelID = "copilot.gpt-5.5"
	CopilotGPT54      ModelID = "copilot.gpt-5.4"
	CopilotGPT54Mini  ModelID = "copilot.gpt-5.4-mini"
	CopilotGPT54Nano  ModelID = "copilot.gpt-5.4-nano"
	CopilotGPT5Mini   ModelID = "copilot.gpt-5-mini"
	CopilotGPT41      ModelID = "copilot.gpt-4.1"
	CopilotGPT4o      ModelID = "copilot.gpt-4o"
	CopilotO4Mini     ModelID = "copilot.o4-mini"

	// Anthropic models
	CopilotClaudeFable51  ModelID = "copilot.claude-fable-5.1"
	CopilotClaudeFable5   ModelID = "copilot.claude-fable-5"
	CopilotClaudeOpus55   ModelID = "copilot.claude-opus-5.5"
	CopilotClaudeOpus5    ModelID = "copilot.claude-opus-5"
	CopilotClaudeOpus48   ModelID = "copilot.claude-opus-4.8"
	CopilotClaudeOpus47   ModelID = "copilot.claude-opus-4.7"
	CopilotClaudeSonnet55 ModelID = "copilot.claude-sonnet-5.5"
	CopilotClaudeSonnet5  ModelID = "copilot.claude-sonnet-5"
	CopilotClaudeSonnet46 ModelID = "copilot.claude-sonnet-4.6"
	CopilotClaudeHaiku45  ModelID = "copilot.claude-haiku-4.5"

	// Google and xAI models
	CopilotGemini38Flash ModelID = "copilot.gemini-3.8-flash"
	CopilotGemini35Flash ModelID = "copilot.gemini-3.5-flash"
	CopilotGrok47        ModelID = "copilot.grok-4.7"
	CopilotGrok46        ModelID = "copilot.grok-4.6"
	CopilotGrok45        ModelID = "copilot.grok-4.5"
)

// CopilotAnthropicModels are the Copilot models served over Anthropic's SSE
// wire format; the Copilot client picks its parser from this list.
var CopilotAnthropicModels = []ModelID{
	CopilotClaudeFable51,
	CopilotClaudeFable5,
	CopilotClaudeOpus55,
	CopilotClaudeOpus5,
	CopilotClaudeOpus48,
	CopilotClaudeOpus47,
	CopilotClaudeSonnet55,
	CopilotClaudeSonnet5,
	CopilotClaudeSonnet46,
	CopilotClaudeHaiku45,
}

// copilotModel builds a Copilot entry. Every model is free — it is billed
// against the GitHub Copilot subscription — and supports attachments.
// CanReason is only set for the OpenAI family because Copilot's
// OpenAI-compatible endpoint does not accept reasoning_effort for Claude,
// Gemini or Grok models.
func copilotModel(id ModelID, apiModel, display string, ctx, max int64, canReason bool) Model {
	return Model{
		ID:                  id,
		Name:                display,
		Provider:            ProviderCopilot,
		APIModel:            apiModel,
		ContextWindow:       ctx,
		DefaultMaxTokens:    max,
		CanReason:           canReason,
		SupportsAttachments: true,
	}
}

// https://docs.github.com/en/copilot/using-github-copilot/using-github-copilot-chat/asking-github-copilot-questions-in-your-ide
var CopilotModels = map[ModelID]Model{
	CopilotGPT61Sol: copilotModel(CopilotGPT61Sol, "gpt-6.1-sol",
		"GitHub Copilot GPT-6.1 Sol", 1_050_000, 128_000, true),
	CopilotGPT6Astra: copilotModel(CopilotGPT6Astra, "gpt-6-astra",
		"GitHub Copilot GPT-6 Astra", 1_050_000, 128_000, true),
	CopilotGPT6Luna: copilotModel(CopilotGPT6Luna, "gpt-6-luna",
		"GitHub Copilot GPT-6 Luna", 1_050_000, 128_000, true),
	CopilotGPT56Sol: copilotModel(CopilotGPT56Sol, "gpt-5.6-sol",
		"GitHub Copilot GPT-5.6 Sol", 1_050_000, 128_000, true),
	CopilotGPT56Terra: copilotModel(CopilotGPT56Terra, "gpt-5.6-terra",
		"GitHub Copilot GPT-5.6 Terra", 1_050_000, 128_000, true),
	CopilotGPT56Luna: copilotModel(CopilotGPT56Luna, "gpt-5.6-luna",
		"GitHub Copilot GPT-5.6 Luna", 1_050_000, 128_000, true),
	CopilotGPT55: copilotModel(CopilotGPT55, "gpt-5.5",
		"GitHub Copilot GPT-5.5", 1_050_000, 128_000, true),
	CopilotGPT54: copilotModel(CopilotGPT54, "gpt-5.4",
		"GitHub Copilot GPT-5.4", 1_050_000, 128_000, true),
	CopilotGPT54Mini: copilotModel(CopilotGPT54Mini, "gpt-5.4-mini",
		"GitHub Copilot GPT-5.4 mini", 400_000, 128_000, true),
	CopilotGPT54Nano: copilotModel(CopilotGPT54Nano, "gpt-5.4-nano",
		"GitHub Copilot GPT-5.4 nano", 400_000, 128_000, true),
	CopilotGPT5Mini: copilotModel(CopilotGPT5Mini, "gpt-5-mini",
		"GitHub Copilot GPT-5 mini", 264_000, 64_000, true),
	CopilotGPT41: copilotModel(CopilotGPT41, "gpt-4.1",
		"GitHub Copilot GPT 4.1", 1_047_576, 32_768, false),
	CopilotGPT4o: copilotModel(CopilotGPT4o, "gpt-4o",
		"GitHub Copilot GPT-4o", 128_000, 16_384, false),
	CopilotO4Mini: copilotModel(CopilotO4Mini, "o4-mini",
		"GitHub Copilot o4-mini", 200_000, 100_000, false),

	CopilotClaudeFable51: copilotModel(CopilotClaudeFable51, "claude-fable-5-1",
		"GitHub Copilot Claude Fable 5.1", 1_000_000, 128_000, false),
	CopilotClaudeFable5: copilotModel(CopilotClaudeFable5, "claude-fable-5",
		"GitHub Copilot Claude Fable 5", 1_000_000, 128_000, false),
	CopilotClaudeOpus55: copilotModel(CopilotClaudeOpus55, "claude-opus-5-5",
		"GitHub Copilot Claude Opus 5.5", 1_000_000, 128_000, false),
	CopilotClaudeOpus5: copilotModel(CopilotClaudeOpus5, "claude-opus-5",
		"GitHub Copilot Claude Opus 5", 1_000_000, 64_000, false),
	CopilotClaudeOpus48: copilotModel(CopilotClaudeOpus48, "claude-opus-4-8",
		"GitHub Copilot Claude Opus 4.8", 200_000, 64_000, false),
	CopilotClaudeOpus47: copilotModel(CopilotClaudeOpus47, "claude-opus-4-7",
		"GitHub Copilot Claude Opus 4.7", 200_000, 32_000, false),
	CopilotClaudeSonnet55: copilotModel(CopilotClaudeSonnet55, "claude-sonnet-5-5",
		"GitHub Copilot Claude Sonnet 5.5", 1_000_000, 128_000, false),
	CopilotClaudeSonnet5: copilotModel(CopilotClaudeSonnet5, "claude-sonnet-5",
		"GitHub Copilot Claude Sonnet 5", 1_000_000, 128_000, false),
	CopilotClaudeSonnet46: copilotModel(CopilotClaudeSonnet46, "claude-sonnet-4-6",
		"GitHub Copilot Claude Sonnet 4.6", 200_000, 32_000, false),
	CopilotClaudeHaiku45: copilotModel(CopilotClaudeHaiku45, "claude-haiku-4-5",
		"GitHub Copilot Claude Haiku 4.5", 200_000, 64_000, false),

	CopilotGemini38Flash: copilotModel(CopilotGemini38Flash, "gemini-3.8-flash",
		"GitHub Copilot Gemini 3.8 Flash", 1_048_576, 64_000, false),
	CopilotGemini35Flash: copilotModel(CopilotGemini35Flash, "gemini-3.5-flash",
		"GitHub Copilot Gemini 3.5 Flash", 200_000, 64_000, false),
	CopilotGrok47: copilotModel(CopilotGrok47, "grok-4.7",
		"GitHub Copilot Grok 4.7", 500_000, 128_000, false),
	CopilotGrok46: copilotModel(CopilotGrok46, "grok-4.6",
		"GitHub Copilot Grok 4.6", 500_000, 128_000, false),
	CopilotGrok45: copilotModel(CopilotGrok45, "grok-4.5",
		"GitHub Copilot Grok 4.5", 500_000, 128_000, false),
}
