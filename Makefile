.PHONY: build clean test install

# ビルド情報
VERSION := 0.0.1
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS := -ldflags "-X main.revision=$(REVISION)"

# デフォルトターゲット
all: build

# ビルド
build:
	@echo "Building cwm version $(VERSION) (revision: $(REVISION))"
	go build $(LDFLAGS) -o cwm .

# クロスコンパイル
cross-build:
	@echo "Cross-building cwm version $(VERSION) (revision: $(REVISION))"
	@mkdir -p build
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o build/cwm_darwin_amd64 .
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o build/cwm_darwin_arm64 .
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o build/cwm_linux_amd64 .
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o build/cwm_linux_arm64 .
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o build/cwm_windows_amd64.exe .
	@echo "Cross-build completed successfully!"

# テスト実行
test:
	go test -v ./...

# インストール
install: build
	@echo "Installing cwm to $(GOPATH)/bin"
	@cp cwm $(GOPATH)/bin/

# クリーンアップ
clean:
	@echo "Cleaning up..."
	@rm -rf cwm build/
	@echo "Cleanup completed!"

# ヘルプ
help:
	@echo "Available targets:"
	@echo "  build       - Build cwm binary"
	@echo "  cross-build - Build cwm for multiple platforms"
	@echo "  test        - Run tests"
	@echo "  install     - Install cwm to GOPATH/bin"
	@echo "  clean       - Remove build artifacts"
	@echo "  help        - Show this help message"
