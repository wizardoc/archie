package configer

type QiNiuConfig struct {
	AK     string
	SK     string
	Bucket string `json:"bucket"`
}

func LoadQiNiuConfig() QiNiuConfig {
	config := QiNiuConfig{}
	configLoader("redis.json", &config)

	return config
}
