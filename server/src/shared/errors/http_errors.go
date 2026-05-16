package shared_errors

const (
	TitleRequired    = "title é obrigatório."
	AmountRequired   = "amount deve ser maior que zero."
	CategoryRequired = "category é obrigatório."
	TypeInvalid      = "type deve ser income ou expense."
	IdInvalid        = "id inválido."
)

const (
	MandatoryFieldMessage      = "Existem campos inválidos na requisição."
	InternalServerErrorMessage = "Erro interno ao processar a requisição."
	TransactionNotFoundMessage = "Transação não encontrada."
	InvalidRequestMessage      = "Requisição inválida."
)
