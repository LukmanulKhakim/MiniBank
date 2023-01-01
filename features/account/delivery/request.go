package delivery

import "minibank/features/account/domain"

type AddFormat struct {
	Name    string  `json:"name" form:"name"`
	Balance float64 `json:"balance" form:"balance"`
}

func ToDomain(i interface{}) domain.AccountCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.AccountCore{Name: cnv.Name, Balance: cnv.Balance}
	}
	return domain.AccountCore{}
}
