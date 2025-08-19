package cfg

type LogsConfig struct {
	AddSource bool   `yaml:"add_source" env:"LOGS_ADD_SOURCE" env-default:"false" env-description:"Defines if function source is added to log line"`
	Level     int    `yaml:"level" env:"LOGS_LEVEL" env-default:"0" env-description:"Minimal log level"`
	Format    string `yaml:"format" env:"LOGS_FORMAT" env-default:"txt" env-description:"Logging default format"`
	File      string `yaml:"file" env:"LOGS_FILE" env-description:"Logging file path"`
}
