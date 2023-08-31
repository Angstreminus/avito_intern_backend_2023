package repository

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/lib/pq"
)

type UserSegmentRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserSegmentRepository(dbHandler *sql.DB) *UserSegmentRepository {
	return &UserSegmentRepository{
		dbHandler: dbHandler,
	}
}

func (ur UserSegmentRepository) CreateUserSegment(usegment *model.UserSegments) (*model.UserSegments, *model.ResponseError) {
	query := `
        INSERT INTO USER_SEGMENTS (USER_ID, SEGMENT_NAMES, EXPIRE_DATE)
        VALUES ($1, $2, $3)
        RETURNING ID, USER_ID, SEGMENT_NAMES, EXPIRE_DATE;
    `
	var (
		id              int
		segmentNamesStr string
		segmentNames    []string
	)

	segmentNamesStr = "{" + strings.Join(usegment.SegmentNames, ",") + "}"

	err := ur.dbHandler.QueryRow(query, usegment.UserID, segmentNamesStr).Scan(&id, &usegment.UserID, pq.Array(&segmentNames))
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	UserID := usegment.UserID

	return &model.UserSegments{
		ID:           id,
		UserID:       UserID,
		SegmentNames: segmentNames,
	}, nil
}

func (ur UserSegmentRepository) GetSegmentNamesByUserID(userId int) ([]string, *model.ResponseError) {
	sel := "SELECT SEGMENT_NAMES FROM USER_SEGMENTS WHERE user_id=$1"
	var segmentNames []string

	if err := ur.dbHandler.QueryRow(sel, userId).Scan(pq.Array(&segmentNames)); err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return segmentNames, nil
}

func (ur UserSegmentRepository) AddUserSegments(addSegm []string, userId int) *model.ResponseError {
	segmentNamesStr := "{" + strings.Join(addSegm, ",") + "}"

	query := `UPDATE USER_SEGMENTS SET SEGMENT_NAMES = array_add(SEGMENT_NAMES, $1);`

	res, err := ur.dbHandler.Exec(query, segmentNamesStr)

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

	return nil
}

func (ur UserSegmentRepository) RemoveUserSegments(removeSegm []string, userId int) *model.ResponseError {
	segmentNamesStr := "{" + strings.Join(removeSegm, ",") + "}"

	query := `UPDATE USER_SEGMENTS SET SEGMENT_NAMES = array_remove(SEGMENT_NAMES, $1);`

	res, err := ur.dbHandler.Exec(query, segmentNamesStr)

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

	return nil
}
