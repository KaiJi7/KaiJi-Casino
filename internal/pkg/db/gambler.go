package db

import (
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateGambler(simulationId *primitive.ObjectID, moneyBegin float64) (gambler structs.GamblerData, err error) {
	gambler = structs.GamblerData{
		SimulationId: simulationId,
		MoneyBegin:   moneyBegin,
		MoneyCurrent: moneyBegin,
	}
	res, err := c.Gambler.InsertOne(nil, gambler)
	if err != nil {
		log.Error("fail to insert gambler: ", err.Error())
		return
	}
	id := res.InsertedID.(primitive.ObjectID)
	gambler.Id = &id
	return
}

func (c client) ListGambler(simulationId *primitive.ObjectID) (gamblers []structs.GamblerData, err error) {
	filter := bson.M{
		"simulation_id": simulationId,
	}

	cursor, dbErr := c.Gambler.Find(nil, filter)
	if dbErr != nil {
		log.Error("fail to get gamblers: ", dbErr.Error())
		err = dbErr
		return
	}

	if err = cursor.All(nil, &gamblers); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return
	}
	return
}
