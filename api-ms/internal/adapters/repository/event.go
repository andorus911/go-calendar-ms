package repository

import (
	"context"
	"github.com/andorus911/go-calendar-ms/api-ms/tools/domain/models"
	"time"
)

// TODO
func SaveEvent(ctx context.Context, event *models.Event) error {
	return nil
}

func DeleteEvent(ctx context.Context, event *models.Event) error {
	return nil
}

func GetEventById(ctx context.Context, is string) (*models.Event, error) {
	return nil, nil
}

func GetEventByOwnerStartTime(ctx context.Context, owner string, startTime time.Time) ([]*models.Event, error) {
	return nil, nil
}
