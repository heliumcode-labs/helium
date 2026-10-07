package config

import (
	"testing"

	"github.com/heliumcode-labs/helium/internal/llm/models"
)

// TestLoadWithoutProviderConfigured guards the regression that made a fresh
// install crash with "agent coder not found": when no provider credentials are
// present, the config must still expose a populated agent registry and a
// provider entry the coder agent can be built against.
func TestLoadWithoutProviderConfigured(t *testing.T) {
	for _, key := range []string{
		"ANTHROPIC_API_KEY", "OPENAI_API_KEY", "GEMINI_API_KEY", "GROQ_API_KEY",
		"OPENROUTER_API_KEY", "XAI_API_KEY", "LOCAL_ENDPOINT",
		"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_REGION", "AWS_PROFILE",
		"GOOGLE_APPLICATION_CREDENTIALS", "VERTEXAI_PROJECT", "GOOGLE_CLOUD_PROJECT",
	} {
		t.Setenv(key, "")
	}

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", tmp+"/.config")

	if _, err := Load(t.TempDir(), false); err != nil {
		t.Fatalf("Load failed with no provider configured: %v", err)
	}

	for _, name := range []AgentName{AgentCoder, AgentSummarizer, AgentTask, AgentTitle} {
		agent, ok := cfg.Agents[name]
		if !ok {
			t.Fatalf("agent %q missing from config; startup would fail", name)
		}
		if _, ok := models.SupportedModels[agent.Model]; !ok {
			t.Fatalf("agent %q has unsupported model %q", name, agent.Model)
		}
	}

	coderModel := cfg.Agents[AgentCoder].Model
	provider := models.SupportedModels[coderModel].Provider
	if _, ok := cfg.Providers[provider]; !ok {
		t.Fatalf("provider %q for coder model %q missing; NewAgent would fail", provider, coderModel)
	}
}
