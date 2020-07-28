package request

type (
	CustomError struct {
		Message 	string 		`json:"message"`
		Code 		int 		`json:"code"`
	}

)
