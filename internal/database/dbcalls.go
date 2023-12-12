package database

import (
	"context"

	"github.com/codepnw/go-short-url/internal/constant"
	"github.com/codepnw/go-short-url/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *manager) Insert(data any, collectionName string) (any, error) {
	inst := m.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)

	return result.InsertedID, err
}

func (m *manager) GetUrlFromCode(code, collectionName string) (resp models.UrlDB, err error) {
	inst := m.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	
	return resp, err
}
