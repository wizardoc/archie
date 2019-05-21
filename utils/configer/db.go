package configer

type DBConfig struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}

func LoadDBConfig() DBConfig {
	config := DBConfig{}
	configLoader("db.json", &config)

	return config
}
