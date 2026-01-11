package data

import (
	"context"
	"fmt"
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const codingSkillsCollName = "codingSkills"

type CodingSkill struct {
	//ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	ID    string `json:"id,omitempty"`
	Title string `json:"title" bson:"title"`
	Image string `json:"image" bson:"image"`
}

func (s *CodingSkill) Insert(entry CodingSkill) error {
	collection := database_.Collection(codingSkillsCollName)

	_, err := collection.InsertOne(context.Background(), entry)
	if err != nil {
		return fmt.Errorf("%w: %s", errorInsertingCodingSkill, err)
	}

	return nil
}

func (s *CodingSkill) All() ([]*CodingSkill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(codingSkillsCollName)

	opts := options.Find()

	sort := bson.D{{"_id", -1}}
	opts.SetSort(sort)

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errorFindingAllCodingSkills, err)
	}
	defer cursor.Close(ctx)

	var codingSkills []*CodingSkill
	for cursor.Next(ctx) {
		var item CodingSkill
		if err := cursor.Decode(&item); err != nil {
			return nil, fmt.Errorf("%w: %s", errorDecodingCodingSkillsIntoSlice, err)
		} else {
			codingSkills = append(codingSkills, &item)
		}
	}

	return codingSkills, nil
}

func (s *CodingSkill) GetOne(title string) (*CodingSkill, error) {
	collection := database_.Collection(codingSkillsCollName)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"title": title}

	var result CodingSkill
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %s", errorGettingCodingSkillByTitle, title, err)
	}

	return &result, nil
}

func (s *CodingSkill) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(codingSkillsCollName)
	if err := collection.Drop(ctx); err != nil {
		return fmt.Errorf("%s: %s", errorDroppingCollection, err)
	}

	return nil
}

func (s *CodingSkill) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(codingSkillsCollName)

	docID, err := bson.ObjectIDFromHex(s.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorDecodingToObjectID, err)
	}

	result, err := collection.UpdateOne(ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", s.Title},
				{"image", s.Image},
			}},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorUpdatingCodingSkill, err)
	}

	return result, nil
}
