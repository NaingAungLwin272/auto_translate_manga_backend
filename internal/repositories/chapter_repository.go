package repositories

import (
	"auto_translate_manga_backend/internal/configs"
	"auto_translate_manga_backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChapterRepository struct {
	coll *mongo.Collection
}

func NewChapterRepository(client *mongo.Client, cfg *configs.Config) *ChapterRepository {
	return &ChapterRepository{
		coll: client.Database(cfg.DBName).Collection("chapters"),
	}
}

func chapterPipeline(mangaId string) (mongo.Pipeline, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mangas"},
			{Key: "localField", Value: "manga"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "manga"},
		}}},

		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$manga"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "genres"},
			{Key: "localField", Value: "manga.genres"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "manga.genres"},
		}}},
	}

	if mangaId != "" {
		objectID, err := bson.ObjectIDFromHex(mangaId)
		if err != nil {
			return nil, fmt.Errorf("invalid manga id format: %v", err)
		}

		pipeline = append(mongo.Pipeline{{{Key: "$match", Value: bson.D{{Key: "manga", Value: objectID}}}}}, pipeline...)

	}

	return pipeline, nil
}

func (chapterRepository *ChapterRepository) CreateChapter(ctx context.Context, chapters *models.Chapter) (*models.Chapter, error) {
	chapters.CreatedAt = time.Now()
	chapters.UpdatedAt = time.Now()
	result, err := chapterRepository.coll.InsertOne(ctx, chapters)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("chapter already exists")
		}
		return nil, err
	}

	var oid bson.ObjectID
	chapterResult, _ := json.Marshal(result.InsertedID)
	if err := json.Unmarshal(chapterResult, &oid); err != nil {
		return nil, err
	}

	chapters.ID = oid
	return chapters, nil
}

func (chapterRepository *ChapterRepository) GetAllChaptersByMangaID(ctx context.Context, mangaId string) (*[]models.ChapterWithMangaAndGenres, error) {
	pipeline, err := chapterPipeline(mangaId)
	if err != nil {
		return nil, err
	}
	cursor, err := chapterRepository.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate failed: %v", err)
	}
	defer cursor.Close(ctx)

	var chapters []models.ChapterWithMangaAndGenres
	if err := cursor.All(ctx, &chapters); err != nil {
		return nil, err
	}

	if len(chapters) == 0 {
		return &[]models.ChapterWithMangaAndGenres{}, nil
	}

	return &chapters, nil
}

func (chapterRepository *ChapterRepository) DeleteChapterByID(ctx context.Context, mangaId string) (string, error) {
	return "", nil
}
