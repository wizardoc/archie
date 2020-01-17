package configer

type QiNiuConfig struct {
	AK     string
	SK     string
	Bucket string `json:"bucket"`
}

func LoadQiNiuConfig() QiNiuConfig {
	config := QiNiuConfig{}
	configLoader("qiniu.json", &config)

	return config
}
