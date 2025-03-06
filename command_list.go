package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func doList(ctx *cli.Context) error {
	// config変数はload_config.goで定義されているグローバル変数を使用
	if config.WorkspaceDir == "" {
		return fmt.Errorf("workspace_dirが設定されていません")
	}

	fmt.Printf("ワークスペースディレクトリ: %s\n", config.WorkspaceDir)
	// TODO: workspace_dirからワークスペース一覧を取得して表示する処理を実装

	return nil
}
