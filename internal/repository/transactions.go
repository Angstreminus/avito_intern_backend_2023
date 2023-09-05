package repository

import (
	"context"
	"database/sql"

	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
)

func BeginTransaction(segmentsUserRep *SegmentsUserRepository) apperrors.AppError {
	ctx := context.Background()
	transaction, err := segmentsUserRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return &apperrors.TransactionErr{
			Message: "FAILED TO BEGIN TRANSACTION",
		}
	}
	segmentsUserRep.Transaction = transaction
	return nil
}

func RollBackTransaction(segmentsUserRep *SegmentsUserRepository) error {
	transaction := segmentsUserRep.Transaction
	segmentsUserRep.Transaction = nil
	return transaction.Rollback()
}

func CommitTransaction(segmentsUserRep *SegmentsUserRepository) error {
	transaction := segmentsUserRep.Transaction
	segmentsUserRep.Transaction = nil
	return transaction.Commit()
}
