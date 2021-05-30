package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateSimulation(strategySchema map[strategy.Name]int, initialMoney float64) (simulation collection.Simulation, err error) {
	simulation = collection.Simulation{
		GamblerInitialMoney: initialMoney,
		StrategySchema:      strategySchema,
	}

	res, err := c.Simulation.InsertOne(nil, simulation)
	if err != nil {
		log.Error("fail to insert simulation: ", err.Error())
		return
	}
	simulation.Id = res.InsertedID.(*primitive.ObjectID)
	return
}
