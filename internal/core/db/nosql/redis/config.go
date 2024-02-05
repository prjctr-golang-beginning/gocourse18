package redis

var defaultConfig = RedisConfig{
	address:  `127.0.0.1:6379`,
	password: `mycomplicatedpassword`,
}

type RedisConfig struct {
	address  string
	password string
	db       int
}

func (c RedisConfig) Address() string {
	return c.address
}

func (c RedisConfig) Password() string {
	return c.password
}

func (c RedisConfig) DB() int {
	return c.db
}
