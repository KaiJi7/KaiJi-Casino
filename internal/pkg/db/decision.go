package db

import (
	"github.com/KaiJi7/common/structs"
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

func (c client) GetDecisions(strategyId *primitive.ObjectID, gambleId *primitive.ObjectID) (decisions []structs.Decision, err error) {
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

func (c client) SaveDecision(decision structs.Decision) (document structs.Decision, err error) {
	res, dbErr := c.Decision.InsertOne(nil, decision)
	if dbErr != nil {
		log.Error("fail to insert document: ", dbErr.Error())
		err = dbErr
		return
	}
	oId := res.InsertedID.(primitive.ObjectID)
	document = decision
	document.Id = &oId
	return
}

//func (c client) SaveDecisions(decisions []interface{}) (err error) {
//	_, err = c.Decision.InsertMany(nil, decisions)
//	return
//}
