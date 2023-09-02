package repository

import (
	"database/sql"
	"net/http"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/lib/pq"
)

type SegmentsUserRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
}

func NewSegmentUserRepository(dbHandler *sql.DB) *SegmentsUserRepository {
	return &SegmentsUserRepository{
		dbHandler: dbHandler,
	}
}

func (sur SegmentsUserRepository) CreateUserSegment(segmentsUsers *model.SegmentsUsers) (*model.SegmentsUsers, *model.ResponseError) {
	query := `
        INSERT INTO SEGMENTS_USERS(USER_ID, SEGMENT_ID)
        VALUES ($1, $2) RETURNING USER_ID, SEGMENT_ID;
    `

	var (
		userId, segmentUserId int
	)

	err := sur.dbHandler.QueryRow(query, segmentsUsers.UserID, segmentsUsers.SegmentID).Scan(&userId, &segmentUserId)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &model.SegmentsUsers{
		UserID:    userId,
		SegmentID: segmentUserId,
	}, nil
}

func (sur SegmentsUserRepository) GetSegmentNamesByUserID(userId int) ([]string, *model.ResponseError) {
	query := `SELECT S.SEGMENT_NAME FROM 
	SEGMENTS S 
	JOIN SEGMENTS_USERS SU ON SU.SEGMENT_ID = S.ID WHERE SU.USER_ID = $1;`

	var segmentNames []string

	if err := sur.dbHandler.QueryRow(query, userId).Scan(pq.Array(&segmentNames)); err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return segmentNames, nil
}

func (sur SegmentsUserRepository) AddSegmentsToUser(toAdd []string, userId int) *model.ResponseError {

	query := `INSERT INTO SEGMENTS_USERS(SEGMENT_ID, USER_ID) VALUES($1,$2);`

	for i := 0; i < len(toAdd); i++ {
		res, err := sur.dbHandler.Exec(query, toAdd[i], userId)

		if err != nil {
			return &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		rowsAff, err := res.RowsAffected()
		if err != nil {
			return &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		if rowsAff == 0 {
			return &model.ResponseError{
				Message: "User not found",
				Status:  http.StatusNotFound,
			}
		}
	}

	return nil
}

func (sur SegmentsUserRepository) RemoveSegmentsUsers(toDel []string, userId int) *model.ResponseError {

	query := `DELETE FROM SEGMENTS_USERS WHERE USER_ID = $2 AND SEGMENT_ID = (
		SELECT ID FROM SEGMENTS
		WHERE SEGMENT_NAME = $1
	)`

	for i := 0; i < len(toDel); i++ {
		res, err := sur.dbHandler.Exec(query, toDel[i], userId)

		if err != nil {
			return &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		rowsAff, err := res.RowsAffected()
		if err != nil {
			return &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		if rowsAff == 0 {
			return &model.ResponseError{
				Message: "User not found",
				Status:  http.StatusNotFound,
			}
		}
	}
	return nil
}
