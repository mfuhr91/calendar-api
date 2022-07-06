package repositories

import (
	"calendar-api/database"
	"calendar-api/models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type eventRepository struct{}

type EventRepository interface {
	GetAll() ([]models.Event, error)
	GetById(id string) (models.Event, error)
	Save(event models.Event) (models.Event, error)
	Update(event models.Event) (models.Event, error)
	Delete(id string) error
}

const (
	databaseName   = "calendardb"
	collectionName = "events"
)

func NewEventRepository() EventRepository {
	return &eventRepository{}
}

func (f *eventRepository) GetAll() ([]models.Event, error) {
	client, err := database.MongoConnect()
	if err != nil {
		return []models.Event{}, err
	}

	collection := client.Database(databaseName).Collection(collectionName)

	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.Event{}, err
	}
	var events []models.Event

	err = cursor.All(context.TODO(), &events)
	if err != nil {
		return []models.Event{}, err
	}

	return events, nil
}

func (f *eventRepository) Save(event models.Event) (models.Event, error) {
	client, err := database.MongoConnect()
	if err != nil {
		return models.Event{}, err
	}

	collection := client.Database(databaseName).Collection(collectionName)

	_, err = collection.InsertOne(context.TODO(),
		bson.M{
			"Title": event.Title,
			"Date":  event.Date,
			"End":   event.End,
			"Color": event.Color,
		})
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (f *eventRepository) Update(event models.Event) (models.Event, error) {
	client, err := database.MongoConnect()
	if err != nil {
		return models.Event{}, err
	}

	collection := client.Database(databaseName).Collection(collectionName)

	id, _ := primitive.ObjectIDFromHex(event.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.M{
		"Title": event.Title,
		"Date":  event.Date,
		"End":   event.End,
		"Color": event.Color,
	}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (f *eventRepository) GetById(id string) (models.Event, error) {
	var event models.Event

	client, err := database.MongoConnect()
	if err != nil {
		return models.Event{}, err
	}

	collection := client.Database(databaseName).Collection(collectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (f *eventRepository) Delete(id string) error {
	client, err := database.MongoConnect()
	if err != nil {
		return err
	}

	collection := client.Database(databaseName).Collection(collectionName)
	objID, _ := primitive.ObjectIDFromHex(id)

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("document id: %s not found", id))
	}

	return nil
}
