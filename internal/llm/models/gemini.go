package models

const (
	ProviderGemini ModelProvider = "gemini"

	// Models
	Gemini38Flash     ModelID = "gemini-3.8-flash"
	Gemini35Flash     ModelID = "gemini-3.5-flash"
	Gemini35FlashLite ModelID = "gemini-3.5-flash-lite"
	Gemini31FlashLite ModelID = "gemini-3.1-flash-lite"
	Gemini31Pro       ModelID = "gemini-3.1-pro"
	Gemini3Flash      ModelID = "gemini-3-flash"
	Gemini25Pro       ModelID = "gemini-2.5-pro"
	Gemini25Flash     ModelID = "gemini-2.5-flash"
	Gemini25FlashLite ModelID = "gemini-2.5-flash-lite"
)

// https://ai.google.dev/gemini-api/docs/models
var GeminiModels = map[ModelID]Model{
	Gemini38Flash: {
		ID:                  Gemini38Flash,
		Name:                "Gemini 3.8 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3.8-flash",
		CostPer1MIn:         0.75,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.075,
		CostPer1MOut:        3.75,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini35Flash: {
		ID:                  Gemini35Flash,
		Name:                "Gemini 3.5 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3.5-flash",
		CostPer1MIn:         1.50,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.15,
		CostPer1MOut:        9.00,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini35FlashLite: {
		ID:                  Gemini35FlashLite,
		Name:                "Gemini 3.5 Flash-Lite",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3.5-flash-lite",
		CostPer1MIn:         0.30,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.03,
		CostPer1MOut:        2.50,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini31FlashLite: {
		ID:                  Gemini31FlashLite,
		Name:                "Gemini 3.1 Flash-Lite",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3.1-flash-lite",
		CostPer1MIn:         0.25,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.025,
		CostPer1MOut:        1.50,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini31Pro: {
		ID:                  Gemini31Pro,
		Name:                "Gemini 3.1 Pro",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3.1-pro-preview",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.2,
		CostPer1MOut:        12.00,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini3Flash: {
		ID:                  Gemini3Flash,
		Name:                "Gemini 3 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-3-flash-preview",
		CostPer1MIn:         0.50,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.05,
		CostPer1MOut:        3.00,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini25Pro: {
		ID:                  Gemini25Pro,
		Name:                "Gemini 2.5 Pro",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-pro",
		CostPer1MIn:         1.25,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.3125,
		CostPer1MOut:        10.00,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini25Flash: {
		ID:                  Gemini25Flash,
		Name:                "Gemini 2.5 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-flash",
		CostPer1MIn:         0.30,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.15,
		CostPer1MOut:        2.50,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
	Gemini25FlashLite: {
		ID:                  Gemini25FlashLite,
		Name:                "Gemini 2.5 Flash-Lite",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-flash-lite",
		CostPer1MIn:         0.10,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.05,
		CostPer1MOut:        0.40,
		ContextWindow:       1_048_576,
		DefaultMaxTokens:    65_536,
		CanReason:           true,
		SupportsAttachments: true,
	},
}
