package config

type Config struct {
	Port string
	Auth struct {
		Pepper   string
		SigningKey string
		TokenTTL   int
	}
	AWS struct {
		Region string
		DynamoDB struct {
			URI           string
			Name          string
			UserTableName string
		}
	}
}