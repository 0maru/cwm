package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()
	
	// テスト用の設定ファイルを作成
	testConfig := `workspace_dir = "/test/workspace"`
	configPath := filepath.Join(tmpDir, "test_config.toml")
	err := os.WriteFile(configPath, []byte(testConfig), 0644)
	assert.NoError(t, err)

	// 環境変数を設定
	originalConfig := os.Getenv("CWM_CONFIG")
	t.Setenv("CWM_CONFIG", configPath)
	defer func() {
		os.Setenv("CWM_CONFIG", originalConfig)
	}()

	// テスト実行
	LoadConfig(nil)

	// 設定が正しく読み込まれたか確認
	assert.Equal(t, "/test/workspace", config.WorkspaceDir)
}

func TestGetConfigPath(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(t *testing.T) string
		expectedPath string
		expectError  bool
	}{
		{
			name: "CWM_CONFIG環境変数が設定されている場合",
			setup: func(t *testing.T) string {
				tmpDir := t.TempDir()
				configPath := filepath.Join(tmpDir, "cwm.toml")
				err := os.WriteFile(configPath, []byte{}, 0644)
				assert.NoError(t, err)
				t.Setenv("CWM_CONFIG", configPath)
				return configPath
			},
			expectError: false,
		},
		{
			name: "XDG_CONFIG_HOMEが設定されている場合",
			setup: func(t *testing.T) string {
				tmpDir := t.TempDir()
				configDir := filepath.Join(tmpDir, "cwm")
				err := os.MkdirAll(configDir, 0755)
				assert.NoError(t, err)
				
				configPath := filepath.Join(configDir, "cwm.toml")
				err = os.WriteFile(configPath, []byte{}, 0644)
				assert.NoError(t, err)
				
				t.Setenv("XDG_CONFIG_HOME", tmpDir)
				t.Setenv("CWM_CONFIG", "")
				return configPath
			},
			expectError: false,
		},
		{
			name: "設定ファイルが見つからない場合",
			setup: func(t *testing.T) string {
				t.Setenv("CWM_CONFIG", "")
				t.Setenv("XDG_CONFIG_HOME", "")
				t.Setenv("HOME", t.TempDir())
				return ""
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPath := tt.setup(t)
			path, err := getConfigPath()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedPath, path)
			}
		})
	}
} 