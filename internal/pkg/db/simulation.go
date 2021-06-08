package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateSimulation(simulation collection.Simulation) (document collection.Simulation, err error) {
	res, dbErr := c.Simulation.InsertOne(nil, simulation)
	if dbErr != nil {
		log.Error("fail to insert simulation: ", dbErr.Error())
		err = dbErr
		return
	}
	oId := res.InsertedID.(primitive.ObjectID)
	document = simulation
	document.Id = &oId
	return
}
