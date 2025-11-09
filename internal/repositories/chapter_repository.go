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
