package casino

import (
	"Lab3/internal/cresponse"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Account struct {
	Id         int
	Money      int
	RealNumber int
	Client     http.Client
}

const (
	Lcg       = "Lcg"
	Mt        = "Mt"
	BetterMt  = "BetterMt"
	CreateAcc = "/createacc"
	Play      = "/play"
	Id        = "id"
	money     = "money"
	baseUrl   = "http://95.217.177.249/casino"
)

func (account *Account) MakeAccount(id int) error {
	params := fmt.Sprintf("%s%s?%s=%d", baseUrl, CreateAcc, Id, id)

	req, err := http.NewRequest(http.MethodGet, params, nil)
	if err != nil {
		return errors.New("[casino] -> Error when generating create account req" + err.Error())
	}
	resp, err := account.Client.Do(req)
	if err != nil {
		return errors.New("[casino] -> Error when receiving response from casino" + err.Error())
	}

	return account.createAccountParser(resp)
}
func (account *Account) createAccountParser(resp *http.Response) error {
	bodyI, err := cresponse.BodeReader(resp)
	if err != nil {
		return err
	}
	body, isOk := bodyI.(map[string]interface{})
	if isOk != true {

	}

	idStr, isOk := body[Id].(string)
	if isOk != true {

	}

	account.Id, err = strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	balance, isOk := body[money].(float64)
	if isOk != true {

	}
	account.Money = int(balance)

	return nil
}
