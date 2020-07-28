package respond

type (
	InitialUser struct {
		Authorized bool `json:"authorized"`
		ID string `json:"id"`
		Role string `json:"role"`
		Company string `json:"company"`
		Exp int64 `json:"exp"`
	}
)
