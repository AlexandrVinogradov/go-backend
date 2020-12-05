// точка входа нашего пиложения (API сервера)
package main

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

// чтобы передавать конфиг в качестве флага нашего бинарника
var (
	configPath string
)

func init() {
	// хотим парсить переменную config 
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	// парсим наши флаги и записываем их в нужные переменные
	flag.Parse()

	//передаем config в функцию apiserver.NewConfig()
	config := apiserver.NewConfig()
	// пользуемся нашей библиотекой, которую мы добавили в качестве завсимостей toml 
	// для того чтобы прочитать этот файл, распарсить его и записать все кго значения в 
	// нашу переменную config
	_, err := toml.DecodeFile(configPath, config)
	// проверяем, что не произошло ошибки
	if err != nil {
		log.Fatal(err)
	}

	// создаем API сервер, который мы будем запускать
	s := apiserver.New(config)

	// запускаем его 
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}