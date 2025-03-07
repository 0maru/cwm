package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

<<<<<<< HEAD
func doList(ctx *cli.Context) error {
	// config変数はload_config.goで定義されているグローバル変数を使用
	if config.WorkspaceDir == "" {
		return fmt.Errorf("workspace_dirが設定されていません")
	}

	fmt.Printf("ワークスペースディレクトリ: %s\n", config.WorkspaceDir)
	// TODO: workspace_dirからワークスペース一覧を取得して表示する処理を実装

||||||| ce3fcf3
func doList(ctx *cli.Context) error {
	fmt.Println("list")
=======
var commandList = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
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
		fmt.Printf("- %s\n", ws)
	}

>>>>>>> 100e2923d64e621000e777a6e14efd05259f43bf
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
			// 拡張子を除いたファイル名を追加
			baseName := strings.TrimSuffix(name, ".code-workspace")
			workspaces = append(workspaces, baseName)
		}
	}

	return workspaces
}
