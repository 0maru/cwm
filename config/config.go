package config

type Config struct {
	Root   string `toml:"root"`
	Editor string `toml:"editor"`
}

// New は新しいConfig構造体のインスタンスを返します
func New() *Config {
	return &Config{}
}
