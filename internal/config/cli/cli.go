package cli

import (
	"flag"
	"log"
	"os"
	"url-shortener/internal/config"
)

var isHelp = flag.Bool("help", false, "show this help message")
var httpPort = flag.String("httpPort", "5678", "set up port for http handler")
var pgDSN = flag.String("pgDSN", "", "set up DSN for postgres")
var env = flag.String("env", "local", "set up env for program")

func MustLoad() config.Config {
	flag.Parse()
	if *isHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *pgDSN == "" {
		log.Fatal("please, set Postgres DSN  (use flag -pgDSN)")
	}

	return config.Config{
		HttpConfig: config.HTTPConfig{
			Port: *httpPort,
		},
		PostgresConfig: config.PostgresConfig{
			DSN: *pgDSN,
		},
		EnvConfig: config.EnvConfig{
			Env: *env,
		},
	}
}
