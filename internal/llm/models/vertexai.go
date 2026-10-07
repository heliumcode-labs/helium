package models

const (
	ProviderVertexAI ModelProvider = "vertexai"

	// Models
	VertexAIGemini38Flash     ModelID = "vertexai.gemini-3.8-flash"
	VertexAIGemini35Flash     ModelID = "vertexai.gemini-3.5-flash"
	VertexAIGemini31Pro       ModelID = "vertexai.gemini-3.1-pro"
	VertexAIGemini25Pro       ModelID = "vertexai.gemini-2.5-pro"
	VertexAIGemini25Flash     ModelID = "vertexai.gemini-2.5-flash"
	VertexAIGemini25FlashLite ModelID = "vertexai.gemini-2.5-flash-lite"
)

// vertexGeminiModel mirrors a Gemini model on Google Cloud Vertex AI: the same
// API model, pricing and context window with the Vertex AI provider attached.
func vertexGeminiModel(id ModelID, base ModelID, display string) Model {
	src := GeminiModels[base]
	src.ID = id
	src.Name = display
	src.Provider = ProviderVertexAI
	return src
}

var VertexAIGeminiModels = map[ModelID]Model{
	VertexAIGemini38Flash: vertexGeminiModel(VertexAIGemini38Flash, Gemini38Flash, "VertexAI: Gemini 3.8 Flash"),
	VertexAIGemini35Flash: vertexGeminiModel(VertexAIGemini35Flash, Gemini35Flash, "VertexAI: Gemini 3.5 Flash"),
	VertexAIGemini31Pro:   vertexGeminiModel(VertexAIGemini31Pro, Gemini31Pro, "VertexAI: Gemini 3.1 Pro"),
	VertexAIGemini25Pro:   vertexGeminiModel(VertexAIGemini25Pro, Gemini25Pro, "VertexAI: Gemini 2.5 Pro"),
	VertexAIGemini25Flash: vertexGeminiModel(VertexAIGemini25Flash, Gemini25Flash, "VertexAI: Gemini 2.5 Flash"),
	VertexAIGemini25FlashLite: vertexGeminiModel(
		VertexAIGemini25FlashLite, Gemini25FlashLite, "VertexAI: Gemini 2.5 Flash-Lite"),
}
