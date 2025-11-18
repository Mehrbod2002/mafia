package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   Server
	Database Database
	Redis    Redis
	RabbitMQ RabbitMQ
	Payment  Payment
	WebRTC   WebRTC
	Logging  Logging
}

type Server struct{ Port string; Debug bool }
type Database struct{ URL string }
type Redis struct{ Addr string }
type RabbitMQ struct{ URL string }
type Payment struct{ Zarinpal string }
type WebRTC struct{ ICEServers []ICEServer }
type ICEServer struct{ URLs []string }
type Logging struct{ Level, Format string }

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil { log.Println("Config file not found, using env") }
	return &Config{
		Server:   Server{Port: viper.GetString("server.port"), Debug: viper.GetBool("server.debug")},
		Database: Database{URL: viper.GetString("database.url")},
		Redis:    Redis{Addr: viper.GetString("redis.addr")},
		RabbitMQ: RabbitMQ{URL: viper.GetString("rabbitmq.url")},
		Payment:  Payment{Zarinpal: viper.GetString("payment.zarinpal")},
		WebRTC:   WebRTC{ICEServers: parseICEServers()},
		Logging:  Logging{Level: viper.GetString("logging.level"), Format: viper.GetString("logging.format")},
	}
}

func parseICEServers() []ICEServer {
	var servers []ICEServer
	viper.UnmarshalKey("webrtc.ice_servers", &servers)
	return servers
}
