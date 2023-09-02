package service

import (
	"errors"
	"net/http"
	"time"

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

func (ss SegmentsUserService) CreateSegmentUser(usrSeg *model.SegmentsUsers) (*model.SegmentsUsers, *model.ResponseError) {

	err := repository.BeginTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	usrSegment, respErr := ss.SegmentsUserRepository.CreateUserSegment(usrSeg)
	if respErr != nil {
		repository.RollBackTransaction(ss.SegmentsUserRepository)
		return nil, respErr
	}

	err = repository.CommitTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return usrSegment, nil
}

func (ss SegmentsUserService) EditSegment(toAdd []string, toDel []string, userId int) *model.ResponseError {
	respErr := checkEmptyEditFields(toDel, toAdd)
	if respErr != nil {
		return respErr
	}

	err := checkIsUniqueNames(toAdd, toDel)
	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	err = repository.BeginTransaction(ss.SegmentsUserRepository)
	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	curentUserSegNames, respErr := ss.SegmentsUserRepository.GetSegmentNamesByUserID(userId)
	if respErr != nil {
		repository.RollBackTransaction(ss.SegmentsUserRepository)
		return respErr
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
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (ss SegmentsUserService) GetSegmentNamesByUserId(userId int) ([]string, *model.ResponseError) {
	return ss.SegmentsUserRepository.GetSegmentNamesByUserID(userId)
}

func validateUserSegment(user *model.User) *model.ResponseError {
	if user.ID < 0 {
		return &model.ResponseError{
			Message: "Invalid user id value",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Expire_date.Before(time.Now()) {
		return &model.ResponseError{
			Message: "Invalid time value",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func checkIsUniqueNames(toAdd, toDel []string) error {
	uniqueMap := make(map[string]bool)
	uniqErr := errors.New("CONFLICT OF NAMES")
	for _, segment := range toAdd {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
		} else {
			return uniqErr
		}
	}

	for _, segment := range toDel {
		if !uniqueMap[segment] {
			uniqueMap[segment] = true
		} else {
			return uniqErr
		}
	}

	return nil
}

func checkEmptyEditFields(toAdd []string, toDel []string) *model.ResponseError {
	if (len(toAdd) == 0) && (len(toDel) == 0) {
		return &model.ResponseError{
			Message: "Empty processing data fields",
			Status:  http.StatusBadRequest,
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
