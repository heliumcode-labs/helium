package models

const (
	ProviderOpenAI ModelProvider = "openai"

	// GPT-6 family
	GPT61Sol  ModelID = "gpt-6.1-sol"
	GPT6Astra ModelID = "gpt-6-astra"
	GPT6Luna  ModelID = "gpt-6-luna"

	// GPT-5.6 family
	GPT56Sol   ModelID = "gpt-5.6-sol"
	GPT56Terra ModelID = "gpt-5.6-terra"
	GPT56Luna  ModelID = "gpt-5.6-luna"

	// GPT-5 family
	GPT55     ModelID = "gpt-5.5"
	GPT54     ModelID = "gpt-5.4"
	GPT54Mini ModelID = "gpt-5.4-mini"
	GPT54Nano ModelID = "gpt-5.4-nano"
	GPT5      ModelID = "gpt-5"
	GPT5Mini  ModelID = "gpt-5-mini"
	GPT5Nano  ModelID = "gpt-5-nano"

	// GPT-4 family
	GPT41     ModelID = "gpt-4.1"
	GPT41Mini ModelID = "gpt-4.1-mini"
	GPT41Nano ModelID = "gpt-4.1-nano"
	GPT4o     ModelID = "gpt-4o"
	GPT4oMini ModelID = "gpt-4o-mini"

	// Reasoning models
	O3     ModelID = "o3"
	O4Mini ModelID = "o4-mini"
)

// https://developers.openai.com/api/docs/models
//
// CostPer1MInCached is the cache *write* price (only GPT-5.6/6 quote one) and
// CostPer1MOutCached is the cached-input read price: 10% of input for the
// GPT-5/6 families, 25% for the GPT-4.1 family, 50% for GPT-4o.
var OpenAIModels = map[ModelID]Model{
	GPT61Sol: {
		ID:                  GPT61Sol,
		Name:                "GPT-6.1 Sol",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-6.1-sol",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   2.5,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        10.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT6Astra: {
		ID:                  GPT6Astra,
		Name:                "GPT-6 Astra",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-6-astra",
		CostPer1MIn:         10.00,
		CostPer1MInCached:   12.5,
		CostPer1MOutCached:  1.00,
		CostPer1MOut:        50.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT6Luna: {
		ID:                  GPT6Luna,
		Name:                "GPT-6 Luna",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-6-luna",
		CostPer1MIn:         0.10,
		CostPer1MInCached:   0.125,
		CostPer1MOutCached:  0.01,
		CostPer1MOut:        0.50,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT56Sol: {
		ID:                  GPT56Sol,
		Name:                "GPT-5.6 Sol",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.6-sol",
		CostPer1MIn:         4.00,
		CostPer1MInCached:   5.0,
		CostPer1MOutCached:  0.40,
		CostPer1MOut:        20.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT56Terra: {
		ID:                  GPT56Terra,
		Name:                "GPT-5.6 Terra",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.6-terra",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   2.5,
		CostPer1MOutCached:  0.20,
		CostPer1MOut:        12.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT56Luna: {
		ID:                  GPT56Luna,
		Name:                "GPT-5.6 Luna",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.6-luna",
		CostPer1MIn:         0.20,
		CostPer1MInCached:   0.25,
		CostPer1MOutCached:  0.02,
		CostPer1MOut:        1.20,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT55: {
		ID:                  GPT55,
		Name:                "GPT-5.5",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.5",
		CostPer1MIn:         5.00,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        30.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT54: {
		ID:                  GPT54,
		Name:                "GPT-5.4",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.4",
		CostPer1MIn:         2.50,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.25,
		CostPer1MOut:        15.00,
		ContextWindow:       1_050_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT54Mini: {
		ID:                  GPT54Mini,
		Name:                "GPT-5.4 mini",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.4-mini",
		CostPer1MIn:         0.75,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.075,
		CostPer1MOut:        4.50,
		ContextWindow:       400_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT54Nano: {
		ID:                  GPT54Nano,
		Name:                "GPT-5.4 nano",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5.4-nano",
		CostPer1MIn:         0.20,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.02,
		CostPer1MOut:        1.25,
		ContextWindow:       400_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT5: {
		ID:                  GPT5,
		Name:                "GPT-5",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5",
		CostPer1MIn:         1.25,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.125,
		CostPer1MOut:        10.00,
		ContextWindow:       400_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT5Mini: {
		ID:                  GPT5Mini,
		Name:                "GPT-5 mini",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5-mini",
		CostPer1MIn:         0.25,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.025,
		CostPer1MOut:        2.00,
		ContextWindow:       400_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT5Nano: {
		ID:                  GPT5Nano,
		Name:                "GPT-5 nano",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-5-nano",
		CostPer1MIn:         0.05,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.005,
		CostPer1MOut:        0.40,
		ContextWindow:       400_000,
		DefaultMaxTokens:    128_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	GPT41: {
		ID:                  GPT41,
		Name:                "GPT-4.1",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-4.1",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        8.00,
		ContextWindow:       1_047_576,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: true,
	},
	GPT41Mini: {
		ID:                  GPT41Mini,
		Name:                "GPT-4.1 mini",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-4.1-mini",
		CostPer1MIn:         0.40,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.10,
		CostPer1MOut:        1.60,
		ContextWindow:       1_047_576,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: true,
	},
	GPT41Nano: {
		ID:                  GPT41Nano,
		Name:                "GPT-4.1 nano",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-4.1-nano",
		CostPer1MIn:         0.10,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.025,
		CostPer1MOut:        0.40,
		ContextWindow:       1_047_576,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: true,
	},
	GPT4o: {
		ID:                  GPT4o,
		Name:                "GPT-4o",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-4o",
		CostPer1MIn:         2.50,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  1.25,
		CostPer1MOut:        10.00,
		ContextWindow:       128_000,
		DefaultMaxTokens:    16_384,
		SupportsAttachments: true,
	},
	GPT4oMini: {
		ID:                  GPT4oMini,
		Name:                "GPT-4o mini",
		Provider:            ProviderOpenAI,
		APIModel:            "gpt-4o-mini",
		CostPer1MIn:         0.15,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.075,
		CostPer1MOut:        0.60,
		ContextWindow:       128_000,
		DefaultMaxTokens:    16_384,
		SupportsAttachments: true,
	},
	O3: {
		ID:                  O3,
		Name:                "o3",
		Provider:            ProviderOpenAI,
		APIModel:            "o3",
		CostPer1MIn:         2.00,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.50,
		CostPer1MOut:        8.00,
		ContextWindow:       200_000,
		DefaultMaxTokens:    100_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	O4Mini: {
		ID:                  O4Mini,
		Name:                "o4-mini",
		Provider:            ProviderOpenAI,
		APIModel:            "o4-mini",
		CostPer1MIn:         1.10,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.275,
		CostPer1MOut:        4.40,
		ContextWindow:       200_000,
		DefaultMaxTokens:    100_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
}
