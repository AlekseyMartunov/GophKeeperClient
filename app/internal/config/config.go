package config

import "os"

type Config struct {
	dataBaseDSN          string `env:"POSTGRES_DSN"`
	serverADDR           string `env:"SERVER_ADDRESS"`
	token                string
	minioAccessKeyID     string `env:"MINIO_ACCESS_KEY_ID"`
	minioSecretAccessKey string `env:"MINIO_SECRET_KEY_ID"`
	minioEndpoint        string `env:"MINIO_ENDPOINT"`
	minioBucketName      string `env:"BUCKET_NAME"`
	clientUserLogin      string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ParseFlags() {
	if key, ok := os.LookupEnv("POSTGRES_DSN"); ok {
		c.dataBaseDSN = key
	}

	if key, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		c.serverADDR = key
	}

	if key, ok := os.LookupEnv("MINIO_ACCESS_KEY_ID"); ok {
		c.minioAccessKeyID = key
	}

	if key, ok := os.LookupEnv("MINIO_SECRET_KEY_ID"); ok {
		c.minioSecretAccessKey = key
	}

	if key, ok := os.LookupEnv("MINIO_ENDPOINT"); ok {
		c.minioEndpoint = key
	}

	if key, ok := os.LookupEnv("BUCKET_NAME"); ok {
		c.minioBucketName = key
	}

}

func (c *Config) PostgresDSN() string {
	return c.dataBaseDSN
}

func (c *Config) ServerADDR() string {
	return c.serverADDR
}

func (c *Config) Token() string {
	return c.token
}

func (c *Config) UpdateToken(t string) {
	c.token = t
}

func (c *Config) MinioAccessKeyID() string {
	return c.minioAccessKeyID
}

func (c *Config) MinioSecretAccessKey() string {
	return c.minioSecretAccessKey
}

func (c *Config) MinioEndpoint() string {
	return c.minioEndpoint
}

func (c *Config) MinioBucketName() string {
	return c.minioBucketName
}

func (c *Config) ClientUserLogin(name string) {
	c.clientUserLogin = name
}

func (c *Config) GetClientUserLogin() string {
	return c.clientUserLogin
}
