package er

import (
	"fmt"
	"runtime"
)

const (
	//NO ERROR
	NoError = 0

	//BAD REQUEST - 400 (10-49)
	ErrorValueEmpty = 10
	ErrorValueWrong = 11

	//NOT FOUND - 404 (50-89)
	ErrorRequestNotFound = 50

	//INTERNAL SERVER ERROR - 500 (90-99)
	ErrorServer = 99
)

var errorDescription = map[int]string{
	NoError: "No error",

	ErrorValueEmpty: "At least one of required value is empty",
	ErrorValueWrong: "At least one of provided value is incorrect",

	ErrorRequestNotFound: "Request does not exist",

	ErrorServer: "Server error, try again",
}

var errorTextPl = map[int]string{
	NoError: "",

	ErrorValueEmpty: "Co najmniej jedna z wymaganych wartości jest pusta",
	ErrorValueWrong: "Co najmniej jedna z podanych wartości jest nieprawidłowa",

	ErrorRequestNotFound: "Żądanie nie istnieje",

	ErrorServer: "Błąd serwera, spróbuj ponownie",
}

type ErrResponse struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

type Err struct {
	code     int
	response ErrResponse
	status   int
}

func NewErr(code int) Err {
	var status int
	if code >= 10 && code <= 49 {
		status = 400
	} else if code >= 50 && code <= 89 {
		status = 404
	} else if code >= 90 && code <= 99 {
		status = 500
	}

	return Err{
		code: code,
		response: ErrResponse{
			Description: errorDescription[code],
			Text:        errorTextPl[code],
		},
		status: status,
	}
}

// func (e *Err) Check() bool {
// 	if e.code > 0 {
// 		return true
// 	}
// 	return false
// }

func (e *Err) GetStatus() int {
	return e.status
}

func (e *Err) GetResponse() ErrResponse {
	return e.response
}

func (e *Err) Print() {
	if e.code > 0 {
		fmt.Println(e.response.Description)
	}
}

// Check returns true and prints an error in the console if the error exists
func Check(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Print("\n#################### ERROR ####################\n\n")
		fmt.Println("IN", file, "AT", line)
		fmt.Println(err.Error())
		fmt.Print("\n##################### END #####################\n\n")
		return true
	}
	return false
}
