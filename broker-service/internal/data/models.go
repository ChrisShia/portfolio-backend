package data

import "go.mongodb.org/mongo-driver/v2/mongo"

var client_ *mongo.Client
var database_ *mongo.Database

func New(mongo *mongo.Client) Models {
	client_ = mongo
	database_ = client_.Database("portfolio")

	return Models{
		Project:     &Project{},
		CodingSkill: &CodingSkill{},
		Achievement: &Achievement{},
	}
}

type Models struct {
	Project     model[Project]
	CodingSkill model[CodingSkill]
	Achievement model[Achievement]
}

type model[T any] interface {
	Insert(T) error
	Update() (*mongo.UpdateResult, error)
	DropCollection() error
	GetOne(title string) (*T, error)
	All() ([]*T, error)
}
