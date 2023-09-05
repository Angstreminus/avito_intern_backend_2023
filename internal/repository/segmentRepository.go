package repository

import (
	"database/sql"

	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
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

func (sr SegmentRepository) CreateSegment(segment *model.Segments) (*model.Segments, apperrors.AppError) {
	query := `INSERT INTO segments(segment_name) VALUES $1 RETURNING id;`

	var segmentId int

	err := sr.dbHandler.QueryRow(query, segment.SegmentName).Scan(segmentId)

	if err != nil {
		return nil, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}

	return &model.Segments{
		Id:          segmentId,
		SegmentName: segment.SegmentName,
	}, nil
}

func (sr SegmentRepository) DeleteSegment(segmentId int) apperrors.AppError {
	query := `DELETE FROM segments WHERE ID = $1;`

	res, err := sr.dbHandler.Exec(query, segmentId)

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

	return nil
}
