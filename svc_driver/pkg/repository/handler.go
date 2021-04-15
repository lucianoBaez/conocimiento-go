// Package repository Implements repository access for accounty
package repository

import (
	"context"
	"fmt"
	"github.com/d-Una-Interviews/svc_driver/internal/utils"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // PostgreSQL migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // for migration
	"go.uber.org/zap"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))

	dsnMongo        = ""
	userMongo       = ""
	passwordMongo   = ""
	authSourceMongo = ""
	shutdowns       []func() error
)

func init() {
	dsnMongo = fmt.Sprintf("mongodb://%s:%s",
		utils.GetEnv("MONGO_HOST", "localhost"),
		utils.GetEnv("MONGO_PORT", "27017"))

	userMongo = utils.GetEnv("USER_MONGO", "dunauser")
	passwordMongo = utils.GetEnv("PASSWORD_MONGO", "password")
	authSourceMongo = utils.GetEnv("AUTH_MONGO", "drivers_mongo")
}

func InitMongoRepository() *mongo.Database {

	clientOpts := options.Client().ApplyURI(dsnMongo).SetAuth(getMongoCredential())
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Congratulations, you're already connected to MongoDB!")

	return client.Database(authSourceMongo)
}

func getMongoCredential() options.Credential {
	var cred options.Credential
	cred.Username = userMongo
	cred.Password = passwordMongo
	cred.AuthSource = authSourceMongo
	return cred
}
