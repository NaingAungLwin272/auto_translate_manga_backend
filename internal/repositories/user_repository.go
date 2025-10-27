package repositories

import (
	"auto_translate_manga_backend/internal/configs"
	"auto_translate_manga_backend/internal/models"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	coll *mongo.Collection
}

func NewUserRepo(client *mongo.Client, cfg *configs.Config) *UserRepo {
	return &UserRepo{
		coll: client.Database(cfg.DBName).Collection("users"),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := r.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	var oid primitive.ObjectID
	b, _ := json.Marshal(res.InsertedID)
	if err := json.Unmarshal(b, &oid); err != nil {
		return nil, err
	}

	user.ID = oid
	return user, nil
}

func (r *UserRepo) GetAllUsers(ctx context.Context) ([]models.User, error) {
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
