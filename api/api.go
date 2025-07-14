package api

import(
	"encoding/json"
	"net/http"
)
//coin balance params
type CoinBalanceParams struct{
	Username string
}
type CoinBalanceResponse struct{
	//successcode
	Code int
	//Account balance
	Balance int64
}


