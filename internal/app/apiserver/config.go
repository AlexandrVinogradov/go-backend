package apiserver 

// Config ...
type Config struct {
	// адрес, на котором мы запускаем наш сервер 
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml: "log_level"`
}

// функция, которая отдает интциализированный конфиг с дефолтныйм параметром, который нас по умолчанию устраивает
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8081",
		LogLevel: "debug",
	}
}