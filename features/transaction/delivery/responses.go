package delivery

import (
	"minibank/features/transaction/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type AddResponse struct {
	Transaction_id uuid.UUID `json:"transaction_id"`
	Credit_account uint      `json:"credit_account"`
	Debit_account  uint      `json:"debit_account"`
	Amount         float64   `json:"amount"`
	Created_at     time.Time `json:"created_at"`
}

type GetResponse struct {
	Transaction_id uuid.UUID `json:"transaction_id"`
	Credit_account uint      `json:"credit_account"`
	Debit_account  uint      `json:"debit_account"`
	Amount         float64   `json:"amount"`
	Created_at     time.Time `json:"created_at"`
}

type ListResponse struct {
	Amount     float64   `json:"amount"`
	Created_at time.Time `json:"created_at"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := basic.(domain.TransactionCore)
		res = AddResponse{
			Transaction_id: cnv.Transaction_id,
			Credit_account: cnv.Credit_account,
			Debit_account:  cnv.Debit_account,
			Amount:         cnv.Amount,
			Created_at:     cnv.Created_at,
		}
	case "all":
		var arr []GetResponse
		cnv := basic.([]domain.TransactionCore)
		for _, val := range cnv {
			arr = append(arr, GetResponse{
				Transaction_id: val.Transaction_id,
				Credit_account: val.Credit_account,
				Debit_account:  val.Debit_account,
				Amount:         val.Amount,
				Created_at:     val.Created_at,
			})
		}
		res = arr
	case "list":
		var arr2 []ListResponse
		cnv2 := basic.([]domain.TransactionCore)
		for _, val2 := range cnv2 {
			arr2 = append(arr2, ListResponse{
				Amount:     val2.Amount,
				Created_at: val2.Created_at,
			})
		}
		res = arr2
	}
	return res
}
