package config

type Config struct {
	dataBaseDSN          string
	serverADDR           string
	token                string
	minioAccessKeyID     string
	minioSecretAccessKey string
	minioEndpoint        string
	minioBucketName      string
	clientUserLogin      string
}

func NewConfig() *Config {
	return &Config{
		dataBaseDSN:          "postgres://admin:1234@localhost:5431/testClient?sslmode=disable",
		serverADDR:           "http://127.0.0.1:8080",
		minioAccessKeyID:     "minio",
		minioSecretAccessKey: "minio123",
		minioEndpoint:        "127.0.0.1:9000",
		minioBucketName:      "default-bucket",
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
