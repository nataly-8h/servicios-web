package repository

import (
	"context"
	"log"

	"crud/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TaskRepository interface {
	Save(Task *entity.Task) (*entity.Task, error)
	FindAll() ([]entity.Task, error)
}

type repo struct{}

//NewTaskRepository creates a new repo
func NewTaskRepository() TaskRepository {
	return &repo{}
}

const (
	projectId      string = "websql-1f451"
	collectionName string = "tasks"
)

func (*repo) Save(Task *entity.Task) (*entity.Task, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":      Task.ID,
		"Name":    Task.Name,
		"Content": Task.Content,
	})
	if err != nil {
		log.Fatalf("Failed adding a new Task: %v", err)
		return nil, err
	}
	return Task, nil
}

func (*repo) FindAll() ([]entity.Task, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var Tasks []entity.Task
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of Tasks: %v", err)
			return nil, err
		}
		Task := entity.Task{
			ID:      doc.Data()["ID"].(int64),
			Name:    doc.Data()["Name"].(string),
			Content: doc.Data()["Content"].(string),
		}
		Tasks = append(Tasks, Task)
	}
	return Tasks, nil
}
