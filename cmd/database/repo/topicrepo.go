package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type TopicRepo interface {
	GetTopic(context.Context, *GetCondition) ([]*model.Topic, error)
	CreateTopic(context.Context, *model.Topic) error
	UpdateTopic(context.Context, *model.Topic) error
}
