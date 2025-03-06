package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
)

// Config はTOMLファイルの構造を定義します
type Config struct {
	WorkspaceDir string `toml:"workspace_dir"`
}

var config Config

func LoadConfig(ctx *cli.Context) {
	configPath, err := getConfigPath()
	if err != nil {
		fmt.Printf("設定ファイルのパスの取得に失敗しました: %v\n", err)
		os.Exit(1)
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("設定ファイルを読み込みました: %s\n", configPath)
}

// getConfigPath は設定ファイルのパスを返す
// 取得の優先度は以下の通り
// 1. $CWM_CONFIG
// 2. $XDG_CONFIG_HOME/cwm/cwm.toml
// 3. ~/.config/cwm/cwm.toml
// 4. ~/.cwm.toml
func getConfigPath() (string, error) {
	// 1. $CWM_CONFIG
	cwmConfig := os.Getenv("CWM_CONFIG")
	if cwmConfig != "" {
		if _, err := os.Stat(cwmConfig); err == nil {
			return cwmConfig, nil
		}
	}

	// 2. $XDG_CONFIG_HOME/cwm/cwm.toml
	if configHome := os.Getenv("XDG_CONFIG_HOME"); configHome != "" {
		xdgConfigHomePath := filepath.Join(configHome, "cwm", "cwm.toml")
		if _, err := os.Stat(xdgConfigHomePath); err == nil {
			return xdgConfigHomePath, nil
		}
	}

	// 3. ~/.config/cwm/cwm.toml
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".config", "cwm", "cwm.toml")
	if _, err := os.Stat(configPath); err == nil {
		return configPath, nil
	}

	// 4. ~/.cwm.toml
	homePath := filepath.Join(home, ".cwm.toml")
	if _, err := os.Stat(homePath); err == nil {
		return homePath, nil
	}

	return "", fmt.Errorf("failed to get config path")
}
