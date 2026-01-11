package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Project struct {
	ID               string    `json:"id,omitempty"`
	Title            string    `json:"title" bson:"title"`
	GithubUrl        string    `json:"github_url" bson:"github_url"`
	Url              string    `json:"url" bson:"url"`
	Technologies     []string  `json:"technologies" bson:"technologies"`
	ShortDescription string    `json:"short_description" bson:"short_description"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

func (l *Project) Insert(entry Project) error {
	collection := database_.Collection("projects")

	_, err := collection.InsertOne(context.TODO(), Project{
		Title:            entry.Title,
		GithubUrl:        entry.GithubUrl,
		Url:              entry.Url,
		Technologies:     entry.Technologies,
		ShortDescription: entry.ShortDescription,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		return fmt.Errorf("%s: %s", errorInsertingProject, err)
	}

	return nil
}

func (l *Project) All() ([]*Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection("projects")

	opts := options.Find()

	sort := bson.D{{"technologies.length", 1}}
	opts.SetSort(sort)

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorFindingAllProjects, err)
	}
	defer cursor.Close(ctx)

	var projects []*Project
	for cursor.Next(ctx) {
		var item Project
		if err := cursor.Decode(&item); err != nil {
			return nil, fmt.Errorf("%s: %s", errorDecodingProjectsIntoSlice, err)
		} else {
			projects = append(projects, &item)
		}
	}

	return projects, nil
}

func (l *Project) GetOne(title string) (*Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection("projects")

	filter := bson.M{"title": title}

	var result Project
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %s", errorGettingProjectByTitle, title, err)
	}

	return &result, nil
}

func (l *Project) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection("projects")
	if err := collection.Drop(ctx); err != nil {
		return fmt.Errorf("%s: %s", errorDroppingCollection, err)
	}

	return nil
}

func (l *Project) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := database_.Collection("projects")

	docID, err := bson.ObjectIDFromHex(l.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorDecodingToObjectID, err)
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", l.Title},
				{"data", l.GithubUrl},
				{"url", l.Url},
				{"technologies", l.Technologies},
				{"short_description", l.ShortDescription},
				{"updated_at", time.Now()},
			}},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errorUpdatingProject, err)
	}

	return result, nil
}
