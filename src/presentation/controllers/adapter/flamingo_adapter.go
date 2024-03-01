package presentation_adapter

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

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
	query := parseQuery(r.Request().URL.Query())

	params := r.Params
	request := presentation.Request[T]{
		Headers: headers,
		Body:    body,
		Params:  params,
		Query:   query,
	}
	return request, nil
}

func parseQuery(args url.Values) map[string][]string {
	fmt.Println(args)
	query := map[string][]string{}
	for k, v := range args {
		v = strings.FieldsFunc(v[0], func(r rune) bool {
			return r == ','
		})
		query[k] = v
	}
	return query
}
