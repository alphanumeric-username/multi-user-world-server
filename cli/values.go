package cli

import (
	"fmt"
	"os"
)

const version string = "0.0.2"

func defaultConfigFilePath() (string, error) {
	cwd, err := os.Getwd()

	return fmt.Sprintf("%s\\config.yaml", cwd), err
}
