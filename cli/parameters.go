package cli

//StartArguments holds the values of "start" command flags and options
type StartArguments struct {
	Help       bool
	ConfigFile string
}

//NewStartArguments creates and returns a StartArguments object with the values provided in the function arguments
func NewStartArguments(help bool, configFile string) StartArguments {
	var arguments StartArguments
	arguments.Help = help
	arguments.ConfigFile = configFile
	return arguments
}

//VersionArguments holds the values of "version" command flags and options
type VersionArguments struct {
	Help    bool
	Version string
}

//NewVersionArguments creates and returns a VersionArguments object with the values provided in the function arguments
func NewVersionArguments(help bool, version string) VersionArguments {
	var arguments VersionArguments
	arguments.Help = help
	arguments.Version = version
	return arguments
}

//ServerConfig holds the values contained in the server's config.yaml file
type ServerConfig struct {
	//Network
	IP       string `yaml:"ip"`
	HTTPPort int    `yaml:"http_port"`
	UDPPort  int    `yaml:"udp_port"`

	//World
	WorldFile string `yaml:"world_file"`
	MaxUsers  int    `yaml:"max_users"`

	//Chat
	EnableChat  bool `yaml:"enable_chat"`
	MaxMessages int  `yaml:"max_messages"`
}

//NewServerConfig creates and returns a ServerConfig object with the values provided in the function arguments
func NewServerConfig(ip string, httpPort int, udpPort int, worldFile string, maxUsers int, enableChat bool, maxMessages int) ServerConfig {
	var config ServerConfig
	config.IP = ip
	config.HTTPPort = httpPort
	config.UDPPort = udpPort
	config.WorldFile = worldFile
	config.MaxUsers = maxUsers
	config.EnableChat = enableChat
	config.MaxMessages = maxMessages

	return config
}
