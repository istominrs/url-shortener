package config

type (
	HTTPConfig struct {
		Port string
	}

	PostgresConfig struct {
		DSN string
	}

	EnvConfig struct {
		Env string
	}

	Config struct {
		HttpConfig     HTTPConfig
		PostgresConfig PostgresConfig
		EnvConfig      EnvConfig
	}
)
