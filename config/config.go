package config

type Config struct {
	// code-workspace's root directory
	Root string `toml:"root"`
}

// New は新しいConfig構造体のインスタンスを返します
func New() *Config {
	return &Config{}
}
