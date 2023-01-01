package delivery

import "minibank/features/transaction/domain"

type AddFormat struct {
	Credit_account uint    `json:"credit_account" form:"credit_account"`
	Debit_account  uint    `json:"debit_account" form:"debit_account"`
	Amount         float64 `json:"amount" form:"amount"`
}

type FilterFormat struct {
	Filter []AddFormat `json:"filter" form:"filter"`
}

func ToDomain(i interface{}) domain.TransactionCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.TransactionCore{
			Credit_account: cnv.Credit_account,
			Debit_account:  cnv.Debit_account,
			Amount:         cnv.Amount,
		}
	}
	return domain.TransactionCore{}
}

func ToDomainArray(i interface{}) []domain.TransactionCore {
	switch i.(type) {
	case FilterFormat:
		cnv := i.(FilterFormat)
		var detail []domain.TransactionCore
		for x := 0; x < len(cnv.Filter); x++ {
			detail = append(detail, domain.TransactionCore{Credit_account: cnv.Filter[x].Credit_account, Debit_account: cnv.Filter[x].Debit_account, Amount: cnv.Filter[x].Amount})
		}
		return detail
	}
	return []domain.TransactionCore{}
}
