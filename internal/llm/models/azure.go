package models

const ProviderAzure ModelProvider = "azure"

const (
	AzureGPT61Sol  ModelID = "azure.gpt-6.1-sol"
	AzureGPT6Luna  ModelID = "azure.gpt-6-luna"
	AzureGPT56Sol  ModelID = "azure.gpt-5.6-sol"
	AzureGPT56Luna ModelID = "azure.gpt-5.6-luna"
	AzureGPT55     ModelID = "azure.gpt-5.5"
	AzureGPT54     ModelID = "azure.gpt-5.4"
	AzureGPT54Mini ModelID = "azure.gpt-5.4-mini"
	AzureGPT54Nano ModelID = "azure.gpt-5.4-nano"
	AzureGPT5      ModelID = "azure.gpt-5"
	AzureGPT5Mini  ModelID = "azure.gpt-5-mini"
	AzureGPT41     ModelID = "azure.gpt-4.1"
	AzureGPT41Mini ModelID = "azure.gpt-4.1-mini"
	AzureGPT41Nano ModelID = "azure.gpt-4.1-nano"
	AzureGPT4o     ModelID = "azure.gpt-4o"
	AzureGPT4oMini ModelID = "azure.gpt-4o-mini"
	AzureO3        ModelID = "azure.o3"
	AzureO4Mini    ModelID = "azure.o4-mini"
)

// azureModel mirrors an OpenAI model: same API model name, pricing, context
// window and capabilities, only the provider and display name differ.
func azureModel(id ModelID, base ModelID, display string) Model {
	src := OpenAIModels[base]
	src.ID = id
	src.Name = display
	src.Provider = ProviderAzure
	return src
}

var AzureModels = map[ModelID]Model{
	AzureGPT61Sol:  azureModel(AzureGPT61Sol, GPT61Sol, "Azure OpenAI – GPT-6.1 Sol"),
	AzureGPT6Luna:  azureModel(AzureGPT6Luna, GPT6Luna, "Azure OpenAI – GPT-6 Luna"),
	AzureGPT56Sol:  azureModel(AzureGPT56Sol, GPT56Sol, "Azure OpenAI – GPT-5.6 Sol"),
	AzureGPT56Luna: azureModel(AzureGPT56Luna, GPT56Luna, "Azure OpenAI – GPT-5.6 Luna"),
	AzureGPT55:     azureModel(AzureGPT55, GPT55, "Azure OpenAI – GPT-5.5"),
	AzureGPT54:     azureModel(AzureGPT54, GPT54, "Azure OpenAI – GPT-5.4"),
	AzureGPT54Mini: azureModel(AzureGPT54Mini, GPT54Mini, "Azure OpenAI – GPT-5.4 mini"),
	AzureGPT54Nano: azureModel(AzureGPT54Nano, GPT54Nano, "Azure OpenAI – GPT-5.4 nano"),
	AzureGPT5:      azureModel(AzureGPT5, GPT5, "Azure OpenAI – GPT-5"),
	AzureGPT5Mini:  azureModel(AzureGPT5Mini, GPT5Mini, "Azure OpenAI – GPT-5 mini"),
	AzureGPT41:     azureModel(AzureGPT41, GPT41, "Azure OpenAI – GPT 4.1"),
	AzureGPT41Mini: azureModel(AzureGPT41Mini, GPT41Mini, "Azure OpenAI – GPT 4.1 mini"),
	AzureGPT41Nano: azureModel(AzureGPT41Nano, GPT41Nano, "Azure OpenAI – GPT 4.1 nano"),
	AzureGPT4o:     azureModel(AzureGPT4o, GPT4o, "Azure OpenAI – GPT-4o"),
	AzureGPT4oMini: azureModel(AzureGPT4oMini, GPT4oMini, "Azure OpenAI – GPT-4o mini"),
	AzureO3:        azureModel(AzureO3, O3, "Azure OpenAI – o3"),
	AzureO4Mini:    azureModel(AzureO4Mini, O4Mini, "Azure OpenAI – o4-mini"),
}
