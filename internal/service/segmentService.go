package service

import (
	"net/http"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/repository"
)

type SegmentService struct {
	SegmentRepository *repository.SegmentRepository
}

func NewSegmentService(segRepo *repository.SegmentRepository) *SegmentService {
	return &SegmentService{
		SegmentRepository: segRepo,
	}
}

func (ss SegmentService) CreateSegment(segment *model.Segments) (*model.Segments, *model.ResponseError) {
	err := validateSegment(segment)
	if err != nil {
		return nil, err
	}

	return ss.SegmentRepository.CreateSegment(segment)
}

func (ss SegmentService) DeleteSegment(segmentId int) *model.ResponseError {
	return ss.SegmentRepository.DeleteSegment(segmentId)
}

func validateSegment(segment *model.Segments) *model.ResponseError {
	if segment.SegmentName == "" {
		return &model.ResponseError{
			Message: "Empty segment name",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
