package shared_errors

import "errors"

var (
	ErrTransactionNotFound = errors.New(TransactionNotFoundMessage)
)
