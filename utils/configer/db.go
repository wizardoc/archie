package configer

type DBConfig struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}

type RedisConfig struct {
	Bind string `json:"bind"`
	Port string `json:"port"`
}

type EmailConfig struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

/** The DB mean's primary DB */
func LoadDBConfig() DBConfig {
	config := DBConfig{}
	configLoader("db.json", &config)

	return config
}

/** redis */
func LoadRedisConfig() RedisConfig {
	config := RedisConfig{}
	configLoader("redis.json", &config)

	return config
}

/** email */
func LoadEmailConfig() EmailConfig {
	config := EmailConfig{}
	configLoader("email.json", &config)

	return config
}
