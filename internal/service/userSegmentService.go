package service

import (
	"net/http"
	"time"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/repository"
)

type UserSegmentService struct {
	UserSegmentRepository *repository.UserSegmentRepository
}

func NewUserSegmentService(usrSegRepo *repository.UserSegmentRepository) *UserSegmentService {
	return &UserSegmentService{
		UserSegmentRepository: usrSegRepo,
	}
}

func (us UserSegmentService) CreateUserSegment(usrSeg *model.UserSegments) (*model.UserSegments, *model.ResponseError) {
	respErr := validateUserSegment(usrSeg)
	if respErr != nil {
		return nil, respErr
	}
	err := repository.BeginTransaction(us.UserSegmentRepository)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	usrSegment, respErr := us.UserSegmentRepository.CreateUserSegment(usrSeg)
	if respErr != nil {
		repository.RollBackTransaction(us.UserSegmentRepository)
		return nil, respErr
	}

	err = repository.CommitTransaction(us.UserSegmentRepository)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return usrSegment, nil
}

func (us UserSegmentService) EditUserSegment(addStr []string, removeStr []string, userId int) *model.ResponseError {
	respErr := checkEmptyEditFields(addStr, removeStr)
	if respErr != nil {
		return respErr
	}

	err := repository.BeginTransaction(us.UserSegmentRepository)
	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	curentUserSegNames, respErr := us.UserSegmentRepository.GetSegmentNamesByUserID(userId)
	if respErr != nil {
		repository.RollBackTransaction(us.UserSegmentRepository)
		return respErr
	}

	if len(addStr) != 0 {
		uniqSegments := uniqueSegments(curentUserSegNames, addStr)

		respErr := us.UserSegmentRepository.AddUserSegments(uniqSegments, userId)
		if respErr != nil {
			repository.RollBackTransaction(us.UserSegmentRepository)
			return respErr
		}
	}
	if len(removeStr) != 0 {
		respErr := us.UserSegmentRepository.RemoveUserSegments(removeStr, userId)
		if respErr != nil {
			repository.RollBackTransaction(us.UserSegmentRepository)
			return respErr
		}
	}
	err = repository.CommitTransaction(us.UserSegmentRepository)
	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (us UserSegmentService) GetSegmentNamesByUserId(userId int) ([]string, *model.ResponseError) {
	return us.UserSegmentRepository.GetSegmentNamesByUserID(userId)
}

func validateUserSegment(usrSegment *model.UserSegments) *model.ResponseError {
	if usrSegment.UserID < 0 {
		return &model.ResponseError{
			Message: "Invalid user id value",
			Status:  http.StatusBadRequest,
		}
	}

	if usrSegment.ExpireDate.Before(time.Now()) {
		return &model.ResponseError{
			Message: "Invalid time value",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func checkEmptyEditFields(addSeg []string, removeSeg []string) *model.ResponseError {
	if (len(addSeg) == 0) && (len(removeSeg) == 0) {
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
