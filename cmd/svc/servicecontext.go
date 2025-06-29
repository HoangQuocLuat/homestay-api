package svc

import (
	"back-end/cmd/config"
	"back-end/cmd/database/repo"
)

type ServiceContext struct {
	Config       config.Config
	TopicRepo    repo.TopicRepo
	HomestayRepo repo.HomestayRepo
	RoomRepo     repo.RoomRepo
	ServiceRepo  repo.ServiceRepo
}

func NewServiceContext(
	c config.Config,
	topicRepo repo.TopicRepo,
	homestayRepo repo.HomestayRepo,
	roomRepo repo.RoomRepo,
	serviceRepo repo.ServiceRepo) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		TopicRepo:    topicRepo,
		HomestayRepo: homestayRepo,
		RoomRepo:     roomRepo,
		ServiceRepo:  serviceRepo,
	}
}
