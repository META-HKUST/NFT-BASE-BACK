package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBSource       string `mapstructure:"DB_SOURCE"`
	LogFilePATH    string `mapstructure:"LOG_FILE_PATH"`
	GinPort        string `mapstructure:"GIN_PORT"`
	GrpcPort       string `mapstructure:"GRPC_PORT"`
	CryptoPath     string `mapstructure:"CRYPTO_PATH"`
	CcpPath        string `mapstructure:"CCP_PATH"`
	ChannelName    string `mapstructure:"CHANNEL_NAME"`
	ChaincodeName  string `mapstructure:"CHAINCODE_NAME"`
	CertPathSuffix string `mapstructure:"CERT_PATH_SUFFIX"`
	KeyPathSuffix  string `mapstructure:"KEY_PATH_SUFFIX"`
	WalletPath     string `mapstructure:"WALLET_PATH"`
	MspId          string `mapstructure:"MSP_ID"`
}

var CONFIG Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&CONFIG)
	return err
}
