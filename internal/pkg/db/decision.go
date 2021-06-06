package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getDecisionFilter struct {
	StrategyId *primitive.ObjectID `bson:"strategy_id,omitempty"`
	GambleId   *primitive.ObjectID `bson:"gamble_id,omitempty"`
}

func (g getDecisionFilter) toBson() (filter bson.M) {
	b, _ := bson.Marshal(g)
	_ = bson.Unmarshal(b, &filter)
	return
}

func (c client) GetDecisions(strategyId *primitive.ObjectID, gambleId *primitive.ObjectID) (decisions []collection.Decision, err error) {
	filter := getDecisionFilter{
		StrategyId: strategyId,
		GambleId:   gambleId,
	}.toBson()

	cursor, dbErr := c.Decision.Find(nil, filter)
	if dbErr != nil {
		log.Error("fail to get decision: ", dbErr.Error())
		err = dbErr
		return
	}
	if err = cursor.All(nil, decisions); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return
	}
	return
}

func (c client) SaveDecision(decision collection.Decision) (err error) {
	_, err = c.Decision.InsertOne(nil, decision)
	if err != nil {
		log.Error("fail to insert document: ", err.Error())
	}
	return
}

//func (c client) SaveDecisions(decisions []interface{}) (err error) {
//	_, err = c.Decision.InsertMany(nil, decisions)
//	return
//}