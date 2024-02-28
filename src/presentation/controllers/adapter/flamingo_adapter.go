package presentation_adapter

import (
	"encoding/json"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
)

func AdapterRequest[T any](r *web.Request) (presentation.Request[T], error) {
	headers := r.Request().Header
	var body T
	if r.Request().Method == "POST" || r.Request().Method == "PUT" {
		err := json.NewDecoder(r.Request().Body).Decode(&body)
		if err != nil {
			return presentation.Request[T]{}, err
		}
	}
	request := presentation.Request[T]{
		Headers: headers,
		Body:    body,
		Params:  r.Params,
	}
	return request, nil
}
