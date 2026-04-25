package errors

import "errors"

var (
	ErrTransactionNotFound = errors.New(TransactionNotFound)
)
