package service

import (
	"fmt"

	"github.com/Ammce/pilot/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	FindById(Id *string) string
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

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func (service *videoService) FindById(Id *string) string {
	fmt.Println("AMKO")
	fmt.Println(service.videos)
	for i := 0; i < len(service.videos); i++ {
		fmt.Println(i)
	}
	return "Amel"
}
