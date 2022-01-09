package service

import (
	"gin/entity"
	"gin/repository"
)

type videoService struct {
	videoRepository repository.VideoRepository
}

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{repo}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.videoRepository.Save(video)
	return video
}

func (s *videoService) Update(video entity.Video) {
	s.videoRepository.Update(video)
}

func (s *videoService) Delete(video entity.Video) {
	s.videoRepository.Delete(video)
}

func (s *videoService) FindAll() []entity.Video {
	return s.videoRepository.FindAll()
}
