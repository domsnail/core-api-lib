package cfg

type ServerConfig struct {
	Host string `yaml:"host" env:"SRV_HOST" env-default:"0.0.0.0" env-description:"Server listener host"`
	Port uint64 `yaml:"port" env:"SRV_PORT" env-default:"6379" env-description:"Server listener port"`

	Reflection bool `yaml:"reflection" env:"SRV_REFLECTION" env-default:"false" env-description:"Enables GRPC reflection"`
	Health     bool `yaml:"health" env:"SRV_HEALTH" env-default:"true" env-description:"Enables GRPC health checks"`
}
