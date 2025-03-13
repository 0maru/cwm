package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var commandList = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List all code-workspaces",
	Action:  doList,
}

func doList(c *cli.Context) error {
	if err := LoadConfig(c); err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
	}

	// 設定ファイルからルートディレクトリを取得
	rootDir := conf.Root

	// チルダ(~)を展開
	if strings.HasPrefix(rootDir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("ホームディレクトリの取得に失敗しました: %w", err)
		}
		rootDir = filepath.Join(home, rootDir[2:])
	}

	// ディレクトリが存在するか確認
	info, err := os.Stat(rootDir)
	if err != nil {
		return fmt.Errorf("ルートディレクトリの確認に失敗しました: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s はディレクトリではありません", rootDir)
	}

	// ディレクトリ内のファイル一覧を取得
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return fmt.Errorf("ディレクトリの読み込みに失敗しました: %w", err)
	}

	// ワークスペースファイルのみをフィルタリング
	workspaces := filterWorkspaces(entries, rootDir)

	// 結果を表示
	if len(workspaces) == 0 {
		fmt.Println("ワークスペースが見つかりませんでした")
		return nil
	}

	fmt.Println("利用可能なワークスペース:")
	for _, ws := range workspaces {
		fmt.Fprintf(os.Stdout, "%s\n", ws)
	}

	return nil
}

// filterWorkspaces は .code-workspace ファイルのみをフィルタリングします
func filterWorkspaces(entries []fs.DirEntry, rootDir string) []string {
	var workspaces []string

	for _, entry := range entries {
		// ディレクトリはスキップ
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		// .code-workspace 拡張子を持つファイルのみを対象とする
		if filepath.Ext(name) == ".code-workspace" {
			// 完全なパスを構築
			fullPath := filepath.Join(rootDir, name)
			workspaces = append(workspaces, fullPath)
		}
	}

	return workspaces
}
