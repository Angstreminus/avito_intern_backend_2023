package service

import (
	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
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

func (ss SegmentService) CreateSegment(segment *model.Segments) (*model.Segments, apperrors.AppError) {
	err := validateSegment(segment)
	if err != nil {
		return nil, err
	}

	return ss.SegmentRepository.CreateSegment(segment)
}

func (ss SegmentService) DeleteSegment(segmentId int) apperrors.AppError {
	return ss.SegmentRepository.DeleteSegment(segmentId)
}

func validateSegment(segment *model.Segments) apperrors.AppError {
	if segment.SegmentName == "" {
		return &apperrors.InvalidDataErr{
			Message: "Empty segment data",
		}
	}
	return nil
}
