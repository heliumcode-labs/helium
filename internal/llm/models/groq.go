package models

const (
	ProviderGROQ ModelProvider = "groq"

	// GROQ models
	GPTOSS120B            ModelID = "openai/gpt-oss-120b"
	GPTOSS20B             ModelID = "openai/gpt-oss-20b"
	Qwen38_27B            ModelID = "qwen/qwen3.8-27b"
	Qwen36_27B            ModelID = "qwen/qwen3.6-27b"
	Llama3_3_70BVersatile ModelID = "llama-3.3-70b-versatile"
	Llama3_1_8BInstant    ModelID = "llama-3.1-8b-instant"
)

// https://console.groq.com/docs/models
//
// CanReason stays false for every Groq model: the Groq OpenAI-compatible API
// rejects the reasoning_effort parameter that the OpenAI client would send.
var GroqModels = map[ModelID]Model{
	GPTOSS120B: {
		ID:                  GPTOSS120B,
		Name:                "GPT-OSS 120B",
		Provider:            ProviderGROQ,
		APIModel:            "openai/gpt-oss-120b",
		CostPer1MIn:         0.15,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.60,
		ContextWindow:       131_072,
		DefaultMaxTokens:    65_536,
		CanReason:           false,
		SupportsAttachments: false,
	},
	GPTOSS20B: {
		ID:                  GPTOSS20B,
		Name:                "GPT-OSS 20B",
		Provider:            ProviderGROQ,
		APIModel:            "openai/gpt-oss-20b",
		CostPer1MIn:         0.075,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.30,
		ContextWindow:       131_072,
		DefaultMaxTokens:    65_536,
		CanReason:           false,
		SupportsAttachments: false,
	},
	Qwen38_27B: {
		ID:                  Qwen38_27B,
		Name:                "Qwen 3.8 27B",
		Provider:            ProviderGROQ,
		APIModel:            "qwen/qwen3.8-27b",
		CostPer1MIn:         0.80,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        4.00,
		ContextWindow:       131_042,
		DefaultMaxTokens:    16_384,
		CanReason:           false,
		SupportsAttachments: true,
	},
	Qwen36_27B: {
		ID:                  Qwen36_27B,
		Name:                "Qwen 3.6 27B",
		Provider:            ProviderGROQ,
		APIModel:            "qwen/qwen3.6-27b",
		CostPer1MIn:         0.60,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        3.00,
		ContextWindow:       131_072,
		DefaultMaxTokens:    16_384,
		CanReason:           false,
		SupportsAttachments: true,
	},
	Llama3_3_70BVersatile: {
		ID:                  Llama3_3_70BVersatile,
		Name:                "Llama 3.3 70B Versatile",
		Provider:            ProviderGROQ,
		APIModel:            "llama-3.3-70b-versatile",
		CostPer1MIn:         0.59,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.79,
		ContextWindow:       131_072,
		DefaultMaxTokens:    32_768,
		CanReason:           false,
		SupportsAttachments: false,
	},
	Llama3_1_8BInstant: {
		ID:                  Llama3_1_8BInstant,
		Name:                "Llama 3.1 8B Instant",
		Provider:            ProviderGROQ,
		APIModel:            "llama-3.1-8b-instant",
		CostPer1MIn:         0.05,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.08,
		ContextWindow:       131_072,
		DefaultMaxTokens:    131_072,
		CanReason:           false,
		SupportsAttachments: false,
	},
}
