package repositories

import (
	"auto_translate_manga_backend/internal/configs"
	"auto_translate_manga_backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MangaRepo struct {
	coll *mongo.Collection
}

func NewMangaRepo(client *mongo.Client, cfg *configs.Config) *MangaRepo {
	return &MangaRepo{
		coll: client.Database(cfg.DBName).Collection("mangas"),
	}
}

func mangaPipeline(id string, name string, genres []string, status string) (mongo.Pipeline, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "genres"},
			{Key: "localField", Value: "genres"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "genres"},
		}}},
	}

	if id != "" {
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid id: %v", err)
		}
		pipeline = append(mongo.Pipeline{{{Key: "$match", Value: bson.D{{Key: "_id", Value: objectID}}}}}, pipeline...)
	}

	matchConditions := bson.A{}

	if name != "" {
		matchConditions = append(matchConditions, bson.D{
			{Key: "title", Value: bson.D{
				{Key: "$regex", Value: name},
				{Key: "$options", Value: "i"},
			}},
		})
	}

	if status != "" {
		matchConditions = append(matchConditions, bson.D{
			{Key: "status", Value: bson.D{
				{Key: "$regex", Value: status},
				{Key: "$options", Value: "i"},
			}},
		})
	}

	if len(genres) > 0 {
		orGenres := bson.A{}
		for _, g := range genres {
			g = strings.TrimSpace(g)
			orGenres = append(orGenres, bson.D{
				{Key: "genres.name", Value: bson.D{
					{Key: "$regex", Value: g},
					{Key: "$options", Value: "i"},
				}},
			})
		}

		if len(orGenres) > 0 {
			matchConditions = append(matchConditions, bson.D{{Key: "$or", Value: orGenres}})
		}
	}

	if len(matchConditions) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: matchConditions}}}})
	}

	return pipeline, nil
}

func (mangaRepo *MangaRepo) CreateManga(ctx context.Context, manga *models.Manga) (*models.Manga, error) {
	manga.CreatedAt = time.Now()
	manga.UpdatedAt = time.Now()

	result, err := mangaRepo.coll.InsertOne(ctx, manga)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("title '%s' already exists", manga.Title)
		}
		return nil, err
	}

	var oid bson.ObjectID
	genreResult, _ := json.Marshal(result.InsertedID)
	if err := json.Unmarshal(genreResult, &oid); err != nil {
		return nil, err
	}

	manga.ID = oid
	return manga, nil
}

func (mangaRepo *MangaRepo) GetAllManga(ctx context.Context) ([]models.MangaWithGenres, error) {
	pipeline, err := mangaPipeline("", "", []string{}, "")
	if err != nil {
		return nil, err
	}
	cursor, err := mangaRepo.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate failed: %v", err)
	}
	defer cursor.Close(ctx)

	var mangas []models.MangaWithGenres
	if err := cursor.All(ctx, &mangas); err != nil {
		return nil, err
	}

	if len(mangas) == 0 {
		return []models.MangaWithGenres{}, nil
	}

	return mangas, nil
}

func (mangaRepo *MangaRepo) GetMangaById(ctx context.Context, id string) (*models.MangaWithGenres, error) {
	pipeline, err := mangaPipeline(id, "", []string{}, "")
	if err != nil {
		return nil, err
	}
	cursor, err := mangaRepo.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate failed: %v", err)
	}
	defer cursor.Close(ctx)

	var result models.MangaWithGenres
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode manga: %v", err)
		}
		return &result, nil
	}

	return nil, fmt.Errorf("manga not found")
}

func (mangaRepo *MangaRepo) FilterManga(ctx context.Context, name string, genre []string, status string) (*[]models.MangaWithGenres, error) {
	pipeline, err := mangaPipeline("", name, genre, status)
	if err != nil {
		return nil, err
	}
	cursor, err := mangaRepo.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate failed: %v", err)
	}
	defer cursor.Close(ctx)

	var result []models.MangaWithGenres
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("failed to decode manga results: %v", err)
	}

	return &result, nil
}

func (mangaRepo *MangaRepo) DeleteManga(ctx context.Context, id string) (string, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("invalid id: %s", id)
	}

	res, err := mangaRepo.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return "", fmt.Errorf("failed to delete manga: %v", err)
	}

	if res.DeletedCount == 0 {
		return "", fmt.Errorf("manga not found")
	}

	return "manga deleted successfully", nil
}
