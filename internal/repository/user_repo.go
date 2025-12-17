package repository
import (
	"context"
	"binarykeeda-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	cursor, err := r.Collection.Find(ctx, map[string]interface{}{})
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
