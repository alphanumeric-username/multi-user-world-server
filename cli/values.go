package cli

import (
	"fmt"
	"os"
)

//Version is the running server version
const Version string = "0.0.3"

func defaultConfigFilePath() (string, error) {
	cwd, err := os.Getwd()

	return fmt.Sprintf("%s\\config.yaml", cwd), err
}
