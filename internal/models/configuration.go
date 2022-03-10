package models

type Configuration struct {
	Engine string `mapstructure:"DB_ENGINE"`
	Server string `mapstructure:"DB_SERVER"`
	Port string `mapstructure:"DB_PORT"`
	User string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	Database string `mapstructure:"DB_NAME"`
	Datapath string `mapstructure:"DATAPATH"`
}
