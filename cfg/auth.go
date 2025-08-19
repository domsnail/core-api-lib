package cfg

type AuthConfig struct {
	Token string `yaml:"token" env:"AUTH_TOKEN" env-description:"API security token"`
}
