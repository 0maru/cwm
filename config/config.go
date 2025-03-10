package config

type Config struct {
	// code-workspace's root directory
	Root string `toml:"root"`

	// editor command code or cursor
	Editor string `toml:"editor"`
}

// New は新しいConfig構造体のインスタンスを返します
func New() *Config {
	return &Config{}
}
