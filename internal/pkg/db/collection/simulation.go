package collection

import (
	"KaiJi-Casino/internal/pkg/strategy"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Simulation struct {
	Id                  *primitive.ObjectID   `json:"id"`
	GamblerInitialMoney float64               `json:"gambler_initial_money"`
	StrategySchema      map[strategy.Name]int `json:"strategy_info"`
}

func (s Simulation) String() string {
	return fmt.Sprintf("Id: %s, Initial money: %f, Schema: %v", s.Id.Hex(), s.GamblerInitialMoney, s.StrategySchema)
}
