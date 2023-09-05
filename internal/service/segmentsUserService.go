package service

import (
	"time"

	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/repository"
)

type SegmentsUserService struct {
	SegmentsUserRepository *repository.SegmentsUserRepository
}

func NewSegmentUserService(usrSegRepo *repository.SegmentsUserRepository) *SegmentsUserService {
	return &SegmentsUserService{
		SegmentsUserRepository: usrSegRepo,
	}
}

func (ss SegmentsUserService) CreateSegmentUser(usrSeg *model.SegmentsUsers) (*model.SegmentsUsers, apperrors.AppError) {

	err := repository.BeginTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return nil, err
	}
	usrSegment, err := ss.SegmentsUserRepository.CreateUserSegment(usrSeg)
	if err != nil {
		repository.RollBackTransaction(ss.SegmentsUserRepository)
		return nil, err
	}

	err = repository.CommitTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return nil, err
	}
	return usrSegment, nil
}

func (ss SegmentsUserService) EditSegment(toAdd []string, toDel []string, userId int) apperrors.AppError {
	respErr := checkEmptyEditFields(toDel, toAdd)
	if respErr != nil {
		return respErr
	}

	err := checkIsUniqueNames(toAdd, toDel)
	if err != nil {
		return err
	}

	err = repository.BeginTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return err
	}

	curentUserSegNames, err := ss.SegmentsUserRepository.GetSegmentNamesByUserID(userId)
	if err != nil {
		repository.RollBackTransaction(ss.SegmentsUserRepository)
		return err
	}

	if len(toAdd) != 0 {
		uniqSegments := uniqueSegments(curentUserSegNames, toAdd)

		respErr := ss.SegmentsUserRepository.AddSegmentsToUser(uniqSegments, userId)
		if respErr != nil {
			repository.RollBackTransaction(ss.SegmentsUserRepository)
			return respErr
		}
	}
	if len(toDel) != 0 {
		respErr := ss.SegmentsUserRepository.RemoveSegmentsUsers(toDel, userId)
		if respErr != nil {
			repository.RollBackTransaction(ss.SegmentsUserRepository)
			return respErr
		}
	}
	err = repository.CommitTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return err
	}

	return nil
}

func (ss SegmentsUserService) GetSegmentNamesByUserId(userId int) ([]string, apperrors.AppError) {
	return ss.SegmentsUserRepository.GetSegmentNamesByUserID(userId)
}

func validateUserSegment(user *model.User) apperrors.AppError {
	if user.ID < 0 {
		return &apperrors.InvalidDataErr{
			Message: "Invalid user id value",
		}
	}

	if user.Expire_date.Before(time.Now()) {
		return &apperrors.InvalidDataErr{
			Message: "Invalid time value",
		}
	}
	return nil
}

func checkIsUniqueNames(toAdd, toDel []string) apperrors.AppError {
	uniqueMap := make(map[string]bool)

	for _, segment := range toAdd {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
		} else {
			return &apperrors.InvalidDataErr{
				Message: "TO ADD AND TO DELETE NEAMES MUST BE UNIQUE",
			}
		}
	}

	for _, segment := range toDel {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
		} else {
			return &apperrors.InvalidDataErr{
				Message: "TO ADD AND TO DELETE NEAMES MUST BE UNIQUE",
			}
		}
	}

	return nil
}

func checkEmptyEditFields(toAdd []string, toDel []string) apperrors.AppError {
	if (len(toAdd) == 0) && (len(toDel) == 0) {
		return &apperrors.InvalidDataErr{
			Message: "Empty processing data fields",
		}
	}
	return nil
}

func uniqueSegments(currSegNames, incomingSegNames []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueNames := []string{}

	for _, segment := range currSegNames {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
			uniqueNames = append(uniqueNames, segment)
		}
	}

	for _, segment := range incomingSegNames {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
			uniqueNames = append(uniqueNames, segment)
		}
	}

	return uniqueNames
}
