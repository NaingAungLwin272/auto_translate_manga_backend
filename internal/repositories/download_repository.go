package repositories

import (
	"auto_translate_manga_backend/internal/configs"
	"auto_translate_manga_backend/internal/models"
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DownloadRepository struct {
	coll *mongo.Collection
}

func NewDownloadRepository(client *mongo.Client, cfg *configs.Config) *DownloadRepository {
	return &DownloadRepository{
		coll: client.Database(cfg.DBName).Collection("downloads"),
	}
}

func (downloadRepository *DownloadRepository) CreateDownloadMangaByChapterId(ctx context.Context, download *models.Download) (*models.Download, error) {
	download.CreatedAt = time.Now()
	result, err := downloadRepository.coll.InsertOne(ctx, download)
	if err != nil {
		return nil, err
	}
	var oid bson.ObjectID
	downloadResult, _ := json.Marshal(result.InsertedID)
	if err := json.Unmarshal(downloadResult, &oid); err != nil {
		return nil, err
	}

	download.ID = oid
	return download, nil
}

func (downloadRespository *DownloadRepository) CreateDownloadMangaByMangaId(ctx context.Context, download *models.Download) (*models.Download, error) {
	return nil, nil
}
