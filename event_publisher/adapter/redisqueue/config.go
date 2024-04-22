package redisqueue

type Config struct {
	Host     string `env:"REDIS_HOST,notEmpty"`
	Port     string `env:"REDIS_PORT,notEmpty"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int64  `env:"REDIS_DB,notEmpty"`
}
