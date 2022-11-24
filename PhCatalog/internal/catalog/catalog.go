package catalog

type Medicine struct {
	Name  string `bson,json:"name"`
	Count int32  `bson,json:"count"`
	Price int32  `bson,json:"price"`
	Id    string `bson,json:"id"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
