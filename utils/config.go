package utils

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func newConfig() (config *Config) {
	return &Config{
		DBDriver: "11111",
		DBSource: "22222",
	}
}
func LoadConfig() (config Config, err error) {
	// viper

	return *newConfig(), nil
}
