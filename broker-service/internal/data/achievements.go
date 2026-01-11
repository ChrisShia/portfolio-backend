package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Achievement struct {
	//ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	ID      string      `json:"id,omitempty"`
	Year    string      `json:"year" bson:"year"`
	Title   string      `json:"title" bson:"title"`
	Icon    string      `json:"icon" bson:"icon"`
	Content []Paragraph `json:"content,omitempty" bson:"content,omitempty"`
}

type Paragraph struct {
	Numbered bool     `json:"numbered,omitempty" bson:"numbered,omitempty"`
	Bulleted bool     `json:"bulleted,omitempty" bson:"bulleted,omitempty"`
	Body     []string `json:"body,omitempty" bson:"body,omitempty"`
}

func (p *Paragraph) UnmarshalJSON(b []byte) error {
	var unmarshalled struct {
		Numbered bool     `json:"numbered,omitempty"`
		Bulleted bool     `json:"bulleted,omitempty"`
		Body     []string `json:"body,omitempty"`
	}

	if err := json.Unmarshal(b, &unmarshalled); err != nil {
		log.Println("error Unmarshalling paragraph: ", err)
		return err
	}

	return nil
}

const achievementsCollName = "achievements"

func (a *Achievement) Insert(entry Achievement) error {
	collection := database_.Collection(achievementsCollName)

	_, err := collection.InsertOne(context.Background(), entry)
	if err != nil {
		return fmt.Errorf("%s: %s", errorInsertingAchievement, err)
	}

	return nil
}

func (a *Achievement) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(achievementsCollName)

	objectID, err := bson.ObjectIDFromHex(a.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorDecodingToObjectID, err)
	}

	result, err := collection.UpdateOne(ctx,
		bson.M{"_id": objectID},
		bson.D{
			{"$set", bson.D{
				{"title", a.Title},
				{"icon", a.Icon},
				{"content", a.Content},
				{"year", a.Year},
			}},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorUpdatingAchievement, err)
	}

	return result, nil
}

func (a *Achievement) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(achievementsCollName)
	if err := collection.Drop(ctx); err != nil {
		return fmt.Errorf("%s: %s", errorDroppingCollection, err)
	}

	return nil
}

func (a *Achievement) GetOne(title string) (*Achievement, error) {
	collection := database_.Collection(achievementsCollName)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"title": title}

	var achievement Achievement
	err := collection.FindOne(ctx, filter).Decode(&achievement)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %s", errorGettingAchievementByTitle, title, err)
	}

	return &achievement, nil
}

func (a *Achievement) All() ([]*Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection(achievementsCollName)

	opts := options.Find()

	sort := bson.D{{"year", -1}}
	opts.SetSort(sort)

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errorFindingAllAchievements, err)
	}
	defer cursor.Close(ctx)

	var achievements []*Achievement
	for cursor.Next(ctx) {
		var item Achievement
		if err := cursor.Decode(&item); err != nil {
			return nil, fmt.Errorf("%w: %s", errorDecodingAchievementsIntoSlice, err)
		} else {
			achievements = append(achievements, &item)
		}
	}

	return achievements, nil
}
