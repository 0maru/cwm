package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var commandOpen = &cli.Command{
	Name:      "open",
	Aliases:   []string{"o"},
	Usage:     "Open a code-workspace",
	ArgsUsage: "WORKSPACE_NAME",
	Action:    doOpen,
}

func doOpen(c *cli.Context) error {
	// 引数からワークスペース名を取得
	if c.NArg() < 1 {
		return fmt.Errorf("ワークスペース名を指定してください")
	}
	workspaceName := c.Args().First()

	// 設定ファイルを読み込む
	if err := LoadConfig(c); err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
	}

	// ルートディレクトリを取得
	rootDir := conf.Root

	// チルダ(~)を展開
	if strings.HasPrefix(rootDir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("ホームディレクトリの取得に失敗しました: %w", err)
		}
		rootDir = filepath.Join(home, rootDir[2:])
	}

	// ワークスペースファイルのパスを構築
	workspaceFile := filepath.Join(rootDir, workspaceName+".code-workspace")

	// ファイルの存在確認
	if _, err := os.Stat(workspaceFile); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("ワークスペース '%s' が見つかりません", workspaceName)
		}
		return fmt.Errorf("ワークスペースファイルの確認に失敗しました: %w", err)
	}

	// エディタコマンドを取得
	editorCmd := conf.Editor
	if editorCmd == "" {
		// デフォルトのエディタを使用
		editorCmd = "code"
	}

	// エディタでワークスペースを開く
	cmd := exec.Command(editorCmd, workspaceFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("ワークスペース '%s' を開いています...\n", workspaceName)

	// コマンドを実行
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("エディタの起動に失敗しました: %w", err)
	}

	// バックグラウンドで実行するため、Wait()は呼び出さない
	return nil
}
