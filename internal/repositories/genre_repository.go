package repositories

import (
	"auto_translate_manga_backend/internal/configs"
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type GenreRepo struct {
	coll *mongo.Collection
}

func NewGenreRepo(client *mongo.Client, cfg *configs.Config) *GenreRepo {
	return &GenreRepo{
		coll: client.Database(cfg.DBName).Collection("genres"),
	}
}

func (genreRepo *GenreRepo) CreateGenre(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	result, err := genreRepo.coll.InsertOne(ctx, genre)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("name '%s' already exists", genre.Name)
		}
		return nil, err
	}

	var oid bson.ObjectID
	genreResult, _ := json.Marshal(result.InsertedID)
	if err := json.Unmarshal(genreResult, &oid); err != nil {
		return nil, err
	}

	genre.ID = oid
	return genre, nil
}

func (genreRepo *GenreRepo) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	cursor, err := genreRepo.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var genres []models.Genre
	if err := cursor.All(ctx, &genres); err != nil {
		return nil, err
	}
	return genres, nil
}

func (genreRepo *GenreRepo) GetGenreById(ctx context.Context, id string) (*models.Genre, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %s", id)
	}

	var genre models.Genre
	err = genreRepo.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&genre)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("genre not found")
		}
		return nil, err
	}

	return &genre, nil
}

func (genreRepo *GenreRepo) UpdateGenreById(ctx context.Context, id string, genre *dtos.UpdateGenreDTO) (*models.Genre, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %s", id)
	}

	updateField := bson.M{}
	if genre.Name != nil {
		updateField["name"] = *genre.Name
	}

	if genre.Description != nil {
		updateField["description"] = *genre.Description
	}

	if len(updateField) == 0 {
		return nil, fmt.Errorf("genre not found")
	}

	update := bson.M{"$set": updateField}

	var updatedGenre models.Genre
	err = genreRepo.coll.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedGenre)

	if err != nil {
		return nil, err
	}

	return &updatedGenre, nil
}

func (genreRepo *GenreRepo) DeleteGenreById(ctx context.Context, id string) (string, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("invalid id: %s", id)
	}

	res, err := genreRepo.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return "", fmt.Errorf("failed to delete genre: %v", err)
	}

	if res.DeletedCount == 0 {
		return "", fmt.Errorf("genre not found")
	}

	return "genre deleted successfully", nil
}
