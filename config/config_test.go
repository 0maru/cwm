package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestConfigParse(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "cwm.toml")

	testConfig := `
root = "~/workspaces"
editor = "code"
`
	if err := os.WriteFile(configPath, []byte(testConfig), 0644); err != nil {
		t.Fatalf("failed to create test config file: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		expected Config
		wantErr  bool
	}{
		{
			name:  "valid config",
			input: configPath,
			expected: Config{
				Root:   "~/workspaces",
				Editor: "code",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := New()
			_, err := toml.DecodeFile(tt.input, conf)

			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if conf.Root != tt.expected.Root {
				t.Errorf("Root = %v, want %v", conf.Root, tt.expected.Root)
			}
			if conf.Editor != tt.expected.Editor {
				t.Errorf("Editor = %v, want %v", conf.Editor, tt.expected.Editor)
			}
		})
	}
}

func TestConfigParseInvalid(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid_config.toml")

	invalidConfig := `
root = /invalid/path    # need quote
editor = "code"
`
	if err := os.WriteFile(configPath, []byte(invalidConfig), 0644); err != nil {
		t.Fatalf("failed to create invalid test config file: %v", err)
	}

	conf := New()
	_, err := toml.DecodeFile(configPath, conf)
	if err == nil {
		t.Error("DecodeFile() expected error for invalid TOML, got nil")
	}
}
