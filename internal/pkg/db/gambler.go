package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateGambler(simulationId *primitive.ObjectID, moneyBegin float64) (gambler collection.GamblerData, err error) {
	gambler = collection.GamblerData{
		SimulationId: simulationId,
		MoneyBegin:   moneyBegin,
		MoneyCurrent: moneyBegin,
	}
	res, err := c.Gambler.InsertOne(nil, gambler)
	if err != nil {
		log.Error("fail to insert gambler: ", err.Error())
		return
	}
	gambler.Id = res.InsertedID.(*primitive.ObjectID)
	return
}

func (c client) ListGambler(simulationId *primitive.ObjectID) (gamblers []collection.GamblerData, err error) {
	filter := bson.M{
		"simulationId": simulationId,
	}

	cursor, err := c.Gambler.Find(nil, filter)
	if err != nil {
		log.Error("fail to get gamblers: ", err.Error())
		return
	}

	if err := cursor.All(nil, gamblers); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return
	}
	return
}
