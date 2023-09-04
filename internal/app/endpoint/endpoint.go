package endpoint

type Endpoint struct {
}

func New() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	return nil
}
