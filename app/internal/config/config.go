package config

import (
	"os"
	"strings"
	"errors"
	"net"
	"fmt"
)

type Config struct {
	dataBaseDSN          string `env:"POSTGRES_DSN"`
	serverADDR           string `env:"SERVER_ADDRESS"`
	token                string
	minioAccessKeyID     string `env:"MINIO_ACCESS_KEY_ID"`
	minioSecretAccessKey string `env:"MINIO_SECRET_KEY_ID"`
	minioEndpoint        string `env:"MINIO_ENDPOINT"`
	minioBucketName      string `env:"BUCKET_NAME"`
	clientUserLogin      string
	migrationPath        string `env:"MIGRATION_PATH"`
}

func NewConfig() *Config {
	return &Config{
		serverADDR:           "http://127.0.0.1:8080",
		dataBaseDSN:          "postgres://admin:1234@localhost:5431/testClient?sslmode=disable",
		minioAccessKeyID:     "minio",
		minioSecretAccessKey: "minio123",
		minioEndpoint:        "127.0.0.1:9000",
		minioBucketName:      "default-bucket",
		migrationPath:        "./migrations",
	}
}

func (c *Config) ParseFlags() error {
	if key, ok := os.LookupEnv("POSTGRES_DSN"); ok {
		c.dataBaseDSN = key
	}

	if key, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		ip, err := getIPByDNSName(key)
		if err != nil {
			return err
		}
		c.serverADDR = ip
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

	if key, ok := os.LookupEnv("MIGRATION_PATH"); ok {
		c.migrationPath = key
	}

	return nil
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

func (c *Config) MigrationPath() string {
	return c.migrationPath
}

func getIPByDNSName(nameDNS string) (string, error) {
	arr := strings.Split(nameDNS, ":")
	if len(arr) < 2 {
		return "",  errors.New("parsing dns name error")
	}

	port := arr[1]

	ips, err := net.LookupIP(arr[0])
	if err != nil {
		return "", err
	}

	if len(ips) < 1 {
		return "", errors.New("not found ip add by name ")
	}

	ip := ips[0].String()


	return fmt.Sprintf("http://%s:%s", ip, port), nil


}

