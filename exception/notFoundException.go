package exception

type NotFoundException struct {
    Error error
    Message string
    Code int
}

func NewNotFoundException(message string) NotFoundException {
  exception := NotFoundException{}
  exception.Message = message
  exception.Code = 404
}
