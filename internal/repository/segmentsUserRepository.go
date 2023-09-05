package repository

import (
	"database/sql"

	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
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

func (sur SegmentsUserRepository) CreateUserSegment(segmentsUsers *model.SegmentsUsers) (*model.SegmentsUsers, apperrors.AppError) {
	query := `
        INSERT INTO SEGMENTS_USERS(USER_ID, SEGMENT_ID)
        VALUES ($1, $2) RETURNING USER_ID, SEGMENT_ID;
    `

	var (
		userId, segmentUserId int
	)

	err := sur.dbHandler.QueryRow(query, segmentsUsers.UserID, segmentsUsers.SegmentID).Scan(&userId, &segmentUserId)
	if err != nil {
		return nil, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}

	return &model.SegmentsUsers{
		UserID:    userId,
		SegmentID: segmentUserId,
	}, nil
}

func (sur SegmentsUserRepository) GetSegmentNamesByUserID(userId int) ([]string, apperrors.AppError) {
	query := `SELECT S.SEGMENT_NAME FROM 
	SEGMENTS S 
	JOIN SEGMENTS_USERS SU ON SU.SEGMENT_ID = S.ID WHERE SU.USER_ID = $1;`

	var segmentNames []string

	if err := sur.dbHandler.QueryRow(query, userId).Scan(pq.Array(&segmentNames)); err != nil {
		return nil, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}

	return segmentNames, nil
}

func (sur SegmentsUserRepository) AddSegmentsToUser(toAdd []string, userId int) apperrors.AppError {

	query := `INSERT INTO SEGMENTS_USERS(SEGMENT_ID, USER_ID) VALUES($1,$2);`

	for i := 0; i < len(toAdd); i++ {
		res, err := sur.dbHandler.Exec(query, toAdd[i], userId)

		if err != nil {
			return &apperrors.DBoperationErr{
				Message: err.Error(),
			}
		}

		rowsAff, err := res.RowsAffected()
		if err != nil {
			return &apperrors.DBoperationErr{
				Message: err.Error(),
			}
		}

		if rowsAff == 0 {
			return &apperrors.DBoperationErr{
				Message: sql.ErrNoRows.Error(),
			}
		}
	}

	return nil
}

func (sur SegmentsUserRepository) RemoveSegmentsUsers(toDel []string, userId int) apperrors.AppError {

	query := `DELETE FROM SEGMENTS_USERS WHERE USER_ID = $2 AND SEGMENT_ID = (
		SELECT ID FROM SEGMENTS
		WHERE SEGMENT_NAME = $1
	)`

	for i := 0; i < len(toDel); i++ {
		res, err := sur.dbHandler.Exec(query, toDel[i], userId)

		if err != nil {
			return &apperrors.DBoperationErr{
				Message: err.Error(),
			}
		}

		rowsAff, err := res.RowsAffected()
		if err != nil {
			return &apperrors.DBoperationErr{
				Message: err.Error(),
			}
		}

		if rowsAff == 0 {
			return &apperrors.DBoperationErr{
				Message: sql.ErrNoRows.Error(),
			}
		}
	}
	return nil
}
