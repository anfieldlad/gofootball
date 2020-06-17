package config

//Config struct
type Config struct {
	DB *DBConfig
}

//DBConfig struct
type DBConfig struct {
	Dialect  string
	Server   string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

//GetConfig function
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Server:   "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "September2019",
			Name:     "football",
			Charset:  "utf8",
		},
	}
}
