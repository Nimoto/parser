package configs

import (
	"encoding/json"
	"flag"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ServiceName - service name
const ServiceName = "parser"

var options = []option{
	{"config", "string", "", "config file"},
	{"source.file", "string", "file.csv", "source file"},
	{"source.dir", "string", "./source", "source dir"},
	{"source.metafile", "string", "meta.csv", "meta file"},
	{"period", "int", 5, "seconds"},
}

// Config - main config struct
type Config struct {
	Source struct {
		File     string
		Dir      string
		Metafile string
	}
	Period int
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.Read(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Read read parameters for config.
// Read from environment variables, flags or file.
func (c *Config) Read() error {

	for _, o := range options {
		switch o.typing {
		case "string":
			flag.String(o.name, o.value.(string), o.description)
		case "int":
			flag.Int(o.name, o.value.(int), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}
	viper.SetEnvPrefix(ServiceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic("Read config error: " + err.Error())
	}

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigFile(fileName)
		viper.SetConfigType("toml")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

// Print print config structure
func (c *Config) Print() error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	log.Println(string(b))
	return nil
}
