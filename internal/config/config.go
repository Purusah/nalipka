package config

type serviceConfig struct {
	Url  string
}

type storageConfig struct {
	Url string
}

type Config struct {
	Service serviceConfig
	Storage storageConfig
}

func GetConfig() (Config, error) {
	// TODO Read env
	return Config{
		Service: serviceConfig{
			Url:  "0.0.0.0:8080",
		},
		Storage: storageConfig{
			Url: "postgres://postgres:postgres@0.0.0.0:5432/default",
		},
	}, nil
}
