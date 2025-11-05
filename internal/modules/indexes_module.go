package modules

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SetupIndexes(db *mongo.Database) {
	// users collection
	_, err := db.Collection("users").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Fatal("Failed to create index for users:", err)
	}

	// genres collection
	_, err = db.Collection("genres").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true).
			SetCollation(&options.Collation{
				Locale:   "en",
				Strength: 2,
			}),
	})
	if err != nil {
		log.Fatal("Failed to create index for genres:", err)
	}

	// mangas collection
	_, err = db.Collection("mangas").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true).
			SetCollation(&options.Collation{
				Locale:   "en",
				Strength: 2,
			}),
	})
	if err != nil {
		log.Fatal("Failed to create index for genres:", err)
	}

	// favorite collection
	_, err = db.Collection("favorites").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "user", Value: 1},
			{Key: "manga", Value: 1},
		},
		Options: options.Index().
			SetUnique(true),
		// SetName("unique_user_manga"), // optional, for clarity
	})
	if err != nil {
		log.Fatal("Failed to create index for favorites:", err)
	}

	log.Println("✅ All indexes created successfully")
}
