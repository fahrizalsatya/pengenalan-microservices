package config

//Config is struct type, which [Port as string, AuthService as AuthService, Database as Database]
type Config struct {
	Port        string
	AuthService AuthService `mapstructure:"auth_service"`
	Database    Database
}
