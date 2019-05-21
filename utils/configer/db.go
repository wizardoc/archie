package configer

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
}

func LoadDBConfig() DBConfig {
	config := DBConfig{}
	configLoader("db.json", &config)

	return config
}