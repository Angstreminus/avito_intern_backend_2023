package repository

import (
	"context"
	"database/sql"
)

func BeginTransaction(userSegmentsRep *UserSegmentRepository) error {
	ctx := context.Background()
	transaction, err := userSegmentsRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	userSegmentsRep.transaction = transaction
	return nil
}

func RollBackTransaction(UserSegmentRep *UserSegmentRepository) error {
	transaction := UserSegmentRep.transaction
	UserSegmentRep.transaction = nil
	return transaction.Rollback()
}

func CommitTransaction(UserSegmentRep *UserSegmentRepository) error {
	transaction := UserSegmentRep.transaction
	UserSegmentRep.transaction = nil
	return transaction.Commit()
}
