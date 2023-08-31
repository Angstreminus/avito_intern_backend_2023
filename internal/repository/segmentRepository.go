package repository

import (
	"database/sql"
	"net/http"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
)

type SegmentRepository struct {
	dbHandler *sql.DB
}

func NewSegmentRepository(dbHandler *sql.DB) *SegmentRepository {
	return &SegmentRepository{
		dbHandler: dbHandler,
	}
}

func (sr SegmentRepository) CreateSegment(segment *model.Segments) (*model.Segments, *model.ResponseError) {
	query := `INSERT INTO segments(segment_name) VALUES $1 RETURNING id;`

	rows, err := sr.dbHandler.Query(query, segment.SegmentName)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var segmentId int

	for rows.Next() {
		err = rows.Scan(segmentId)
		if err != nil {
			return nil, &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &model.Segments{
		Id:          segmentId,
		SegmentName: segment.SegmentName,
	}, nil
}

func (sr SegmentRepository) DeleteSegment(segmentId int) *model.ResponseError {
	query := `DELETE FROM segments WHERE ID = $1;`

	res, err := sr.dbHandler.Exec(query, segmentId)

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
			Message: "Segment not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}
