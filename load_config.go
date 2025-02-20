package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func LoadConfig(ctx *cli.Context) {
	configPath, err := getConfigPath()
	if err != nil {
		fmt.Println("failed to get config path: %w", err)
		os.Exit(1)
	}

	fmt.Println("config path: %s", configPath)
}

// getConfigPath は設定ファイルのパスを返します
func getConfigPath() (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		// XDG_CONFIG_HOME が設定されていない場合は ~/.config を使用
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		configHome = filepath.Join(home, ".config")
	}

	// アプリケーション名に応じてパスを設定
	// TODO: アプリケーション名を適切に設定してください
	appName := "cwm"
	configDir := filepath.Join(configHome, appName)

	// 設定ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return filepath.Join(configDir, "config.toml"), nil
}
