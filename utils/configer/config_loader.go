package configer

import (
	"archie/utils"
	"archie/utils/env_utils"
	"encoding/json"
	"os"
	"path/filepath"
)

func configLoader(fileName string, target interface{}) {
	rootPath, err := os.Getwd()

	utils.Check(err)

	defaultFolder := "dev"
	if env_utils.Env.IsProd() {
		defaultFolder = "prod"
	}

	file, err := os.Open(filepath.Join(rootPath, "configs", defaultFolder, fileName))

	utils.Check(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&target)

	utils.Check(err)
}
