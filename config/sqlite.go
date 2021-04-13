package config

type SqliteConfig struct {
	DBFile string `mapstructure:"sqlite_db_file"`
	Debug  bool   `mapstructure:"sqlite_debug"`
}
