package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/0maru/cwm/config"
	"github.com/urfave/cli/v2"
)

func TestDoOpenNoArgs(t *testing.T) {
	// 引数なしの場合のテスト
	app := cli.NewApp()
	ctx := cli.NewContext(app, nil, nil)

	err := doOpen(ctx)
	if err == nil {
		t.Error("引数なしの場合にエラーが発生しませんでした")
	}
}

func TestDoOpenNonExistentWorkspace(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テスト用の設定を作成
	conf = &config.Config{
		Root:   tmpDir,
		Editor: "echo", // テスト用に実際に実行できるコマンドを指定
	}

	// テスト用のコンテキストを作成（存在しないワークスペース名を指定）
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Parse([]string{"non-existent-workspace"})
	ctx := cli.NewContext(app, set, nil)

	// テスト対象の関数を実行
	err := doOpen(ctx)

	// 存在しないワークスペースを指定した場合はエラーが発生するはず
	if err == nil {
		t.Error("存在しないワークスペースを指定した場合にエラーが発生しませんでした")
	}
}

func TestDoOpenExistingWorkspace(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テスト用のワークスペースファイルを作成
	workspaceName := "test-project"
	workspaceFile := filepath.Join(tmpDir, workspaceName+".code-workspace")
	if err := os.WriteFile(workspaceFile, []byte("test"), 0644); err != nil {
		t.Fatalf("テストファイルの作成に失敗しました: %v", err)
	}

	// テスト用の設定を作成
	conf = &config.Config{
		Root:   tmpDir,
		Editor: "echo", // テスト用に実際に実行できるコマンドを指定
	}

	// テスト用のコンテキストを作成（存在するワークスペース名を指定）
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Parse([]string{workspaceName})
	ctx := cli.NewContext(app, set, nil)

	// テスト対象の関数を実行
	err := doOpen(ctx)

	// エラーがないことを確認
	if err != nil {
		t.Errorf("doOpen関数がエラーを返しました: %v", err)
	}
}
