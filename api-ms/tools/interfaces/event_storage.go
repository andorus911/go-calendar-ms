package interfaces

import (
	"context"
	"github.com/andorus911/go-calendar-ms/api-ms/tools/domain/models"
	"time"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event models.Event) (int64, error)
	DeleteEventById(ctx context.Context, id int64) error
	GetEventById(ctx context.Context, id int64) (*models.Event, error)
	GetEventsByOwnerStartTime(ctx context.Context, owner int64, startTime time.Time) ([]models.Event, error)
}
