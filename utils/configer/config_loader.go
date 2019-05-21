package configer

import (
	"archie/utils"
	"encoding/json"
	"os"
	"path/filepath"
)

func configLoader(fileName string, target interface{}) {
	rootPath, err := os.Getwd()

	utils.Check(err)

	file, err := os.Open(filepath.Join(rootPath, "configs", fileName))

	utils.Check(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&target)

	utils.Check(err)
}