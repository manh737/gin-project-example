package service

import "github.com/manh737/gin-project-example/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	GetAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) GetAll() []entity.Video {
	return service.videos
}
