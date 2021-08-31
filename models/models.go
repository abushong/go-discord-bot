package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Health        int                `bson:"health,omitempty"`
	Items         []string           `bson:"items,omitempty"`
	Damage        int                `bson:"damage,omitempty"`
	IsPlaying     bool               `bson:"isPlaying,omitempty"`
	Form          string             `bson:"level,omitempty"`
	Stage         string             `bson:"stage,omitempty"`
	StageProgress int                `bson:"stageProgress,omitempty"`
	Cash          int                `bson:"cash,omitempty"`
}

type StoreItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int                `bson:"price,omitempty"`
	Damage      int                `bson:"damage,omitempty"`
	Healing     int                `bson:"healing,omitempty"`
	Quantity    int                `bson:"quantity,omitempty"`
}

type Store struct {
	Items []StoreItem `bson:"items,omitempty"`
	Cash  int         `bson:"cash,omitempty"`
}

type Boss struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Health  int                `bson:"health,omitempty"`
	Attacks []string           `bson:"attacks,omitempty"`
	Damage  int                `bson:"damage,omitempty"`
}
