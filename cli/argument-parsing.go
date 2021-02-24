package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//ParseVersion parse command line flags and options for version command.
//It returns a VersionArguments object containing the values of options passed
//to the program and a pointer to a FlagSet object used to parse the options
func ParseVersion(args []string) (VersionArguments, *flag.FlagSet, error) {
	var arguments VersionArguments
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	helpPtr := versionCmd.Bool("help", false, "prints this page.")
	parseErr := versionCmd.Parse(args)
	if parseErr == nil {
		arguments = NewVersionArguments(*helpPtr, version)
	}
	return arguments, versionCmd, parseErr
}

//ParseStart parse command line flags and options for start command.
//It returns a StartArguments object containing the values of options passed
//to the program and a pointer to a FlagSet object used to parse the options
func ParseStart(args []string) (StartArguments, *flag.FlagSet, error) {
	var arguments StartArguments
	configFilePath, err := defaultConfigFilePath()
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	if err != nil {
		fmt.Println("Error while parsing start command:")
		fmt.Printf("\t%s", err.Error())
		os.Exit(1)
	}
	helpPtr := startCmd.Bool("help", false, "prints this page.")
	configFilePtr := startCmd.String("config-file", configFilePath, "specify the configuration file path.")
	parseErr := startCmd.Parse(args)
	if parseErr == nil {
		arguments = NewStartArguments(*helpPtr, *configFilePtr)
	}
	return arguments, startCmd, parseErr
}

//LoadServerConfig loads server configuration file specified by "configFilePath" argument
//and returns a ServerConfig object with the configuration file contents
func LoadServerConfig(configFilePath string) (ServerConfig, error) {
	var config ServerConfig

	configRaw, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configRaw, &config)

	return config, err
}
