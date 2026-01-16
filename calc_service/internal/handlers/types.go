package handlers

type TwoNumberRequest struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type DivideRequest struct {
	Dividend *int `json:"dividend"`
	Divisor  *int `json:"divisor"`
}

type SumRequest []int

type ResultResponse struct {
	Result int `json:"result"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}
