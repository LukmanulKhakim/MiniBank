package delivery

import (
	"minibank/features/account/domain"
	"time"
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
	Account_no uint      `json:"account_no"`
	Name       string    `json:"name"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

type GetResponse struct {
	Account_no uint      `json:"account_no"`
	Name       string    `json:"name"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

type GetNoAccount struct {
	Account_no uint `json:"account_no"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := basic.(domain.AccountCore)
		res = AddResponse{
			Account_no: cnv.Account_no,
			Name:       cnv.Name,
			Balance:    cnv.Balance,
			Created_at: cnv.Created_at,
		}
	case "no":
		var arr2 []GetNoAccount
		cnv2 := basic.([]domain.AccountCore)
		for _, val2 := range cnv2 {
			arr2 = append(arr2, GetNoAccount{
				Account_no: val2.Account_no,
			})
		}
		res = arr2

	case "all":
		var arr []GetResponse
		cnv := basic.([]domain.AccountCore)
		for _, val := range cnv {
			arr = append(arr, GetResponse{
				Account_no: val.Account_no,
				Name:       val.Name,
				Balance:    val.Balance,
				Created_at: val.Created_at,
			})
		}
		res = arr
	}
	return res
}
