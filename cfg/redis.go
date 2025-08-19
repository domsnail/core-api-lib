package cfg

import "time"

type RedisConfig struct {
	Host string `yaml:"host" env:"REDIS_HOST" env-default:"0.0.0.0" env-description:"Redis cache server host"`
	Port uint64 `yaml:"port" env:"REDIS_PORT" env-default:"6379" env-description:"Redis cache server port"`

	Database int `yaml:"database" env:"REDIS_DB" env-description:"Redis database name"`

	Username string `yaml:"username" env:"REDIS_USER" env-default:"default" env-description:"Redis connection username"`
	Password string `yaml:"password" env:"REDIS_PASS" env-default:"" env-description:"Redis connection user password"`

	Consumer string   `yaml:"queue_consumer" env:"REDIS_QUEUE_CONSUMER" env-default:"" env-description:"Redis queue consumer identity"`
	Group    string   `yaml:"queue_group" env:"REDIS_QUEUE_GROUP" env-default:"whois_tasks" env-description:"Redis queue name"`
	Streams  []string `yaml:"queue_streams" env:"REDIS_QUEUE_STREAMS" env-description:"List of Redis queue streams"`

	MaxRetries uint64        `yaml:"max_retries" env:"REDIS_MAX_RETRIES" env-default:"3" env-description:"Redis query max retries"`
	DialTimout time.Duration `yaml:"dial_timout" env:"REDIS_DIAL_TIMOUT" env-default:"3000ms" env-description:"Redis max dial timout"`
}
