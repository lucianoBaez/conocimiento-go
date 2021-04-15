package model

import (
	"context"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	driverCollection = "drivers"
)

// User is a model that maps to users table.
type Driver struct {
	Name      string    `bson:"name" validate:"required,min=2,max=50"`
	LastName  string    `bson:"lastname" validate:"required,min=2,max=50"`
	Email     string    `bson:"email" validate:"required,email"`
	Location  int       `bson:"location""`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// actual implementation
type driverRepository struct {
	repository *mongo.Database
}

// UserRepository interfaces for accessing user data
type DriverRepository interface {
	FindAll(ctx context.Context) (error, []*Driver)
	FindAllPaged(ctx context.Context, limit int64, page int64) (error error, paged map[string]interface{})
	FindByRadius(ctx context.Context, radius int) (error, []*Driver)
	Create(ctx context.Context, driver *Driver) error
}

func (dr driverRepository) FindAll(ctx context.Context) (error, []*Driver) {
	var drivers []*Driver
	findOpts := options.Find()
	collection := dr.repository.Collection(driverCollection)
	cur, err := collection.Find(ctx, bson.D{{}}, findOpts)

	for cur.Next(ctx) {
		var s Driver
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}
		drivers = append(drivers, &s)
	}

	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return err, drivers
}

func (dr driverRepository) FindAllPaged(ctx context.Context, limit int64, page int64) (error error, paged map[string]interface{}) {
	var drivers []*Driver
	var drivers1 []Driver

	filter := bson.M{}
	collection := dr.repository.Collection(driverCollection)
	projection := bson.D{
		{"name", 1},
		{"location", 1},
		{"lastname", 1},
		{"email", 1},
		{"created_at", 1},
		{"updated_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort("name", 1).Sort("lastname", 1).Select(projection).Filter(filter).Decode(&drivers1).Find()
	if err != nil {
		panic(err)
	}

	for _, raw := range drivers1 {
		drivers = append(drivers, &raw)
	}

	var resp = map[string]interface{}{}
	resp["drivers"] = drivers1 //Store the token in the response
	resp["paging"] = paginatedData.Pagination

	return err, resp
}

func (dr driverRepository) FindByRadius(ctx context.Context, radius int) (error, []*Driver) {
	var drivers []*Driver
	filter := bson.M{"location": bson.M{"$lte": radius}}
	opts := options.Find()
	opts.SetSort(bson.D{{"location", 1}})
	collection := dr.repository.Collection(driverCollection)
	cur, err := collection.Find(ctx, filter, opts)
	for cur.Next(ctx) {
		var s Driver
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}
		drivers = append(drivers, &s)
	}

	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return err, drivers
}

func (dr driverRepository) Create(ctx context.Context, driver *Driver) error {
	collection := dr.repository.Collection(driverCollection)
	_, err := collection.InsertOne(ctx, driver)
	if err != nil {
		return err
	}
	return nil
}

// NewContactRepository returns a new repository
func NewDriverRepository(repo *mongo.Database) DriverRepository {
	return driverRepository{
		repository: repo,
	}
}
