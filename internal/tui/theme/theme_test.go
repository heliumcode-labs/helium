package theme

import (
	"testing"
)

func TestThemeRegistration(t *testing.T) {
	// Get list of available themes
	availableThemes := AvailableThemes()

	// Check if "catppuccin" theme is registered
	catppuccinFound := false
	for _, themeName := range availableThemes {
		if themeName == "catppuccin" {
			catppuccinFound = true
			break
		}
	}

	if !catppuccinFound {
		t.Errorf("Catppuccin theme is not registered")
	}

	// Check if "gruvbox" theme is registered
	gruvboxFound := false
	for _, themeName := range availableThemes {
		if themeName == "gruvbox" {
			gruvboxFound = true
			break
		}
	}

	if !gruvboxFound {
		t.Errorf("Gruvbox theme is not registered")
	}

	// Check if "monokai" theme is registered
	monokaiFound := false
	for _, themeName := range availableThemes {
		if themeName == "monokai" {
			monokaiFound = true
			break
		}
	}

	if !monokaiFound {
		t.Errorf("Monokai theme is not registered")
	}

	// Try to get the themes and make sure they're not nil
	catppuccin := GetTheme("catppuccin")
	if catppuccin == nil {
		t.Errorf("Catppuccin theme is nil")
	}

	gruvbox := GetTheme("gruvbox")
	if gruvbox == nil {
		t.Errorf("Gruvbox theme is nil")
	}

	monokai := GetTheme("monokai")
	if monokai == nil {
		t.Errorf("Monokai theme is nil")
	}

	// Test switching theme
	originalTheme := CurrentThemeName()

	err := SetTheme("gruvbox")
	if err != nil {
		t.Errorf("Failed to set theme to gruvbox: %v", err)
	}

	if CurrentThemeName() != "gruvbox" {
		t.Errorf("Theme not properly switched to gruvbox")
	}

	err = SetTheme("monokai")
	if err != nil {
		t.Errorf("Failed to set theme to monokai: %v", err)
	}

	if CurrentThemeName() != "monokai" {
		t.Errorf("Theme not properly switched to monokai")
	}

	// Switch back to original theme
	_ = SetTheme(originalTheme)
}
func TestHeliumThemeRegistered(t *testing.T) {
	found := false
	for _, themeName := range AvailableThemes() {
		if themeName == "helium" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("helium theme is not registered")
	}

	// The brand theme must define both variants of every base color.
	theme := NewHeliumTheme()
	for name, color := range map[string]struct{ dark, light string }{
		"primary":    {theme.Primary().Dark, theme.Primary().Light},
		"secondary":  {theme.Secondary().Dark, theme.Secondary().Light},
		"accent":     {theme.Accent().Dark, theme.Accent().Light},
		"error":      {theme.Error().Dark, theme.Error().Light},
		"warning":    {theme.Warning().Dark, theme.Warning().Light},
		"success":    {theme.Success().Dark, theme.Success().Light},
		"info":       {theme.Info().Dark, theme.Info().Light},
		"text":       {theme.Text().Dark, theme.Text().Light},
		"muted":      {theme.TextMuted().Dark, theme.TextMuted().Light},
		"emphasized": {theme.TextEmphasized().Dark, theme.TextEmphasized().Light},
		"background": {theme.Background().Dark, theme.Background().Light},
		"border":     {theme.BorderNormal().Dark, theme.BorderNormal().Light},
	} {
		if color.dark == "" || color.light == "" {
			t.Errorf("%s: missing dark (%q) or light (%q) variant", name, color.dark, color.light)
		}
		if color.dark == color.light {
			t.Errorf("%s: dark and light variants are identical (%q)", name, color.dark)
		}
	}
}
