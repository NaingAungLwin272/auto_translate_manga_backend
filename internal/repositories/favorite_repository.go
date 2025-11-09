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

type FavoriteRepo struct {
	coll *mongo.Collection
}

func NewFavoriteRepo(client *mongo.Client, cfg *configs.Config) *FavoriteRepo {
	return &FavoriteRepo{
		coll: client.Database(cfg.DBName).Collection("favorites"),
	}
}

func (favoriteRepo *FavoriteRepo) CreateFavoriteManga(ctx context.Context, favorite *models.Favorite) (*models.Favorite, error) {
	favorite.CreatedAt = time.Now()
	result, err := favoriteRepo.coll.InsertOne(ctx, favorite)
	if err != nil {
		return nil, err
	}

	var oid bson.ObjectID
	genreResult, _ := json.Marshal(result.InsertedID)
	if err := json.Unmarshal(genreResult, &oid); err != nil {
		return nil, err
	}

	favorite.ID = oid
	return favorite, nil
}

func (favoriteRepo *FavoriteRepo) GetAllFavoriteMangaByUserId(ctx context.Context, id string) (*models.FavoriteWithUsersMangaDetail, error) {
	userObjID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %v", err)
	}

	// pipeline for get favorite manga detail
	pipeline := mongo.Pipeline{
		// match the specific user and exclude deleted favorites
		{{
			Key: "$match", Value: bson.D{
				{Key: "user", Value: userObjID},
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$exists", Value: false}}}},
					bson.D{{Key: "deleted_at", Value: nil}},
				}},
			},
		}},
		// lookup manga details
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "mangas"},
				{Key: "localField", Value: "manga"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "manga"},
			},
		}},
		{{Key: "$unwind", Value: "$manga"}}, // unwind manga
		// lookup genres for each manga
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "genres"},
				{Key: "localField", Value: "manga.genres"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "manga.genres"},
			},
		}},
		// lookup user details
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
			},
		}},
		{{Key: "$unwind", Value: "$user"}}, // unwind user
		// group all favorites for this user
		{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$user._id"},
				{Key: "user", Value: bson.D{{Key: "$first", Value: "$user"}}},
				{Key: "favorites", Value: bson.D{{Key: "$push", Value: bson.D{
					{Key: "manga", Value: "$manga"},
					{Key: "created_at", Value: "$created_at"},
				}}}},
			},
		}},
	}

	cursor, err := favoriteRepo.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate failed: %v", err)
	}
	defer cursor.Close(ctx)

	var result models.FavoriteWithUsersMangaDetail
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
		return &result, nil
	}

	return nil, fmt.Errorf("no favorites found for this user")
}

func (favoriteRepo *FavoriteRepo) RemoveFavoriteMangaByUserId(ctx context.Context, userId string, mangaIds []string) (string, error) {
	userObjectId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return "", fmt.Errorf("invalid userId: %v", err)
	}

	var mangaObjectIds []bson.ObjectID
	for _, id := range mangaIds {
		objId, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return "", fmt.Errorf("invalid mangaId: %v", err)
		}
		mangaObjectIds = append(mangaObjectIds, objId)

	}

	filter := bson.M{
		"user":  userObjectId,
		"manga": bson.M{"$in": mangaObjectIds},
	}

	result, err := favoriteRepo.coll.DeleteMany(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("failed to remove favorites: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Sprintf("%d favorites remove unsuccessfully"), nil
	}

	return fmt.Sprintf("%d favorites removed successfully", result.DeletedCount), nil
}
