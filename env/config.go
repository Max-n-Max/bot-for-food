package env

// I would use viper: https://github.com/spf13/viper
type ExchangeConfig struct {
	PublicKey        string                     `yaml:"public_key"`        // Represents the public key used to connect to Exchange API.
	SecretKey        string                     `yaml:"secret_key"`        // Represents the secret key used to connect to Exchange API.

}


// BotConfig contains all config data of the bot, which can be also loaded from config file.
type BotConfig struct {
	ExchangeConfigs  []ExchangeConfig `yaml:"exchange_configs"` // Represents the current exchange configuration.
}
