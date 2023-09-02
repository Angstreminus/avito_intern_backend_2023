package repository

import (
	"context"
	"database/sql"
)

func BeginTransaction(segmentsUserRep *SegmentsUserRepository) error {
	ctx := context.Background()
	transaction, err := segmentsUserRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
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
