package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateSimulation(strategySchema map[collection.StrategyName]int, initialMoney float64) (simulation collection.Simulation, err error) {
	simulation = collection.Simulation{
		GamblerInitialMoney: initialMoney,
		StrategySchema:      strategySchema,
	}

	res, dbErr := c.Simulation.InsertOne(nil, simulation)
	if dbErr != nil {
		log.Error("fail to insert simulation: ", dbErr.Error())
		err = dbErr
		return
	}
	id := res.InsertedID.(primitive.ObjectID)
	simulation.Id = &id
	return
}
