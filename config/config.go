package config

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func newConfig() (config *Config) {
	return &Config{
		DBDriver: "mysql",
		DBSource: "yezzi:yezzi@tcp(localhost:3306)/users",
	}
}
func LoadConfig() (config Config, err error) {
	// viper

	return *newConfig(), nil
}
