package personService

import (
	"app/lesson4/models"
	"app/lesson4/pkg/logger"
	"app/lesson4/pkg/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func Add(user *models.UserInfo) error {

	result, err := mongo.MongoDB.Collection("user_info").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	logger.Logger.Infof("insert user success,id=%s", result.InsertedID)
	return nil
}

func GetByName(name string) (*models.UserInfo, error) {

	one := models.UserInfo{}
	err := mongo.MongoDB.Collection("user_info").Find(context.Background(), bson.M{"name": name}).One(&one)
	return &one, err
}

func GetAllByName(name string) ([]*models.UserInfo, error) {

	var all []*models.UserInfo
	err := mongo.MongoDB.Collection("user_info").Find(context.Background(), bson.M{"name": name}).All(&all)
	return all, err
}

func DeleteByName(name string) (int64, error) {
	result, err := mongo.MongoDB.Collection("user_info").RemoveAll(context.Background(), bson.M{"name": name})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
