#!/bin/bash

set -e

# ビルド情報
VERSION="0.0.1"
REVISION=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
PACKAGE="github.com/0maru/cwm"
LDFLAGS="-X main.revision=$REVISION"

# ビルド対象のOS/アーキテクチャ
TARGETS=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
)

# ビルドディレクトリの作成
mkdir -p build

echo "Building cwm version $VERSION (revision: $REVISION)"

# 各ターゲット向けにビルド
for target in "${TARGETS[@]}"; do
  os=${target%/*}
  arch=${target#*/}

  echo "Building for $os/$arch..."

  output="build/cwm"
  if [ "$os" = "windows" ]; then
    output="$output.exe"
  fi
  output="${output}_${os}_${arch}"

  GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o "$output" .

  echo "Built: $output"
done

# 現在のプラットフォーム用のバイナリをルートディレクトリにコピー
current_os=$(go env GOOS)
current_arch=$(go env GOARCH)
current_binary="build/cwm_${current_os}_${current_arch}"
if [ "$current_os" = "windows" ]; then
  current_binary="${current_binary}.exe"
  cp "$current_binary" "cwm.exe"
  echo "Copied to cwm.exe"
else
  cp "$current_binary" "cwm"
  chmod +x "cwm"
  echo "Copied to cwm"
fi

echo "Build completed successfully!"
