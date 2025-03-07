package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/0maru/cwm/config"
	"github.com/urfave/cli/v2"
)

func TestFilterWorkspaces(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テスト用のファイルを作成
	files := []string{
		"test1.code-workspace",
		"test2.code-workspace",
		"test3.txt",
		"test4.md",
	}

	for _, file := range files {
		filePath := filepath.Join(tmpDir, file)
		if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
			t.Fatalf("テストファイルの作成に失敗しました: %v", err)
		}
	}

	// サブディレクトリも作成
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("サブディレクトリの作成に失敗しました: %v", err)
	}

	// ディレクトリ内のエントリを取得
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("ディレクトリの読み込みに失敗しました: %v", err)
	}

	// テスト対象の関数を実行
	workspaces := filterWorkspaces(entries, tmpDir)

	// 期待される結果
	expected := []string{"test1", "test2"}

	// 結果の検証
	if len(workspaces) != len(expected) {
		t.Errorf("ワークスペースの数が一致しません。期待: %d, 実際: %d", len(expected), len(workspaces))
	}

	// 各ワークスペース名の検証
	for i, ws := range expected {
		found := false
		for _, actual := range workspaces {
			if actual == ws {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ワークスペース '%s' が見つかりません", expected[i])
		}
	}
}

func TestDoList(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テスト用のワークスペースファイルを作成
	workspaceFiles := []string{
		"project1.code-workspace",
		"project2.code-workspace",
	}

	for _, file := range workspaceFiles {
		filePath := filepath.Join(tmpDir, file)
		if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
			t.Fatalf("テストファイルの作成に失敗しました: %v", err)
		}
	}

	// テスト用の設定を作成
	conf = &config.Config{
		Root: tmpDir,
	}

	// テスト用のコンテキストを作成
	app := cli.NewApp()
	ctx := cli.NewContext(app, nil, nil)

	// テスト対象の関数を実行
	err := doList(ctx)

	// エラーがないことを確認
	if err != nil {
		t.Errorf("doList関数がエラーを返しました: %v", err)
	}
}
