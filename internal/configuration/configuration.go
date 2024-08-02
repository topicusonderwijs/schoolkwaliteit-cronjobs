package configuration

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"skw.mijnschoolteam/internal/utils"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Deployment    string            `yaml:"deployment" env:"DEPLOYMENT" env-default:"development"`
	Database      DatabaseConf      `yaml:"database"`
	Logging       LoggingConf       `yaml:"logging"`
	MonitorServer MonitorServerConf `yaml:"monitorserver"`
}

type MonitorServerConf struct {
	Host string `yaml:"host" env:"MONITOR_SERVER_HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"MONITOR_SERVER_PORT" env-default:"4445"`
}

type DatabaseConf struct {
	Host          string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
	Port          int    `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
	User          string `yaml:"user" env:"DATABASE_USER" env-default:"postgres"`
	Password      string `yaml:"password" env:"DATABASE_PASSWORD" env-default:"postgres"`
	Name          string `yaml:"database" env:"DATABASE_NAME" env-default:"postgres"`
	Scheme        string `yaml:"scheme"  env:"DATABASE_SCHEME" env-default:"public"`
	MigrationPath string `yaml:"migrationPath" env:"DATABASE_MIGRATION_PATH" env-default:"migrations"`
	Driver        string `yaml:"driver" env:"DATABASE_DRIVER" env-default:"postgres"`
}

type LoggingConf struct {
	LogLevel    string `yaml:"level" env:"LOG_LEVEL"`
	PrettyPrint bool   `yaml:"prettyPrint" env:"LOG_PRETTY_PRINT" env-default:"false"`
}

func CreateConfiguration() (*Configuration, error) {
	var configPath string
	var cfg Configuration

	// Use golang flags to check if a config file location is passed as argument
	f := flag.NewFlagSet("Sync service", 1)
	f.StringVar(&configPath, "config", "config.yml", "Path to configuration file")

	fu := f.Usage

	header := `This server can be configured using environment variables or with a config file.
By default the config.yaml in this directory is used, if you want to override this start the application with the --config flag`
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(&cfg, &header)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	err := f.Parse(utils.ApplicationArguments())
	if err != nil {
		return nil, err
	}

	// Use the configured config location, or default values when config location does not exist
	var configError error
	if _, err := os.Stat(configPath); err == nil {
		configError = cleanenv.ReadConfig(configPath, &cfg)
		slog.Info(fmt.Sprintf("Read config from file %s", configPath))
	} else {
		slog.Warn("Config file (" + configPath + ") not found, using environment variables and default values.")
		configError = cleanenv.ReadEnv(&cfg)
	}

	if configError != nil {
		return nil, configError
	}

	err = cfg.Validate()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Configuration) Validate() error {
	err := validateDeployment(c.Deployment)
	if err != nil {
		return err
	}

	if c.Logging.LogLevel == "" {
		if c.IsEnvironment("development") {
			c.Logging.LogLevel = "debug"
			c.Logging.PrettyPrint = true
		} else {
			c.Logging.LogLevel = "info"
			c.Logging.PrettyPrint = false
		}
	}

	return nil
}

func validateDeployment(deployment string) error {
	allowedDeployments := map[string]string{
		"development": "development",
		"test":        "test",
		"acceptance":  "acceptance",
		"production":  "production",
	}

	if _, exists := allowedDeployments[deployment]; !exists {
		return fmt.Errorf("deployment [%s] is not a valid deployment. Valid deployments are: %s", deployment, reflect.ValueOf(allowedDeployments).MapKeys())
	}
	return nil
}

func (c *Configuration) IsEnvironment(environment string) bool {
	return c.Deployment == environment
}
