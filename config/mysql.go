package config

type MysqlConfig struct {
	Host     string `mapstructure:"mysql_host"`
	User     string `mapstructure:"mysql_user"`
	Password string `mapstructure:"mysql_password"`
	Port     int    `mapstructure:"mysql_port"`
	Database string `mapstructure:"mysql_database"`
	Debug    bool   `mapstructure:"mysql_debug"`
	Charset  string `mapstructure:"mysql_charset"`
}
