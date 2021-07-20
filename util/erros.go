package util
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
