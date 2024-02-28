package presentation_adapter

import (
	"encoding/json"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
)

func AdapterRequest[T any](r *web.Request) (presentation.Request[T], error) {
	var body T
	if r.Request().Method == "POST" || r.Request().Method == "PUT" {
		err := json.NewDecoder(r.Request().Body).Decode(&body)
		if err != nil {
			return presentation.Request[T]{}, err
		}
	}
	headers := r.Request().Header
	query := r.Request().URL.Query()
	request := presentation.Request[T]{
		Headers: headers,
		Body:    body,
		Params:  r.Params,
		Query:   query,
	}
	return request, nil
}
