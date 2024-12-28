package endpoint

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

type HasAccess func(e Endpoint, r *http.Request) error
type Parser func(r *http.Request, key string) (string, error)

type Value struct {
	Key      string
	Required bool

	HasAccess HasAccess
	Parser    Parser
}

func (v Value) GetValue(r *http.Request) string {
	if v.Parser == nil {
		v.Parser = CombinedParser(HeaderParser, PathParser, QueryParamParser, CookieParser)
	}
	data, err := v.Parser(r, v.Key)
	if err != nil {
		return ""
	}
	return data
}

func (v Value) IsValid(e Endpoint, r *http.Request) bool {
	if v.Parser == nil {
		v.Parser = CombinedParser(HeaderParser, PathParser, QueryParamParser, CookieParser)
	}
	data, err := v.Parser(r, v.Key)
	if (err != nil || len(data) == 0) && v.Required {
		return false
	}
	if v.HasAccess != nil {
		if err = v.HasAccess(e, r); err != nil {
			return false
		}
	}
	return true
}

var (
	_ Parser = HeaderParser
	_ Parser = PathParser
	_ Parser = QueryParamParser
	_ Parser = CookieParser
	_ Parser = CombinedParser(HeaderParser, PathParser, QueryParamParser, CookieParser)
)

// HeaderParser retrieves the value of a header from the HTTP request by key.
func HeaderParser(r *http.Request, key string) (string, error) {
	// Ensure the key is formatted as a proper header key (Title-Case)
	c := cases.Title(language.English)
	formattedKey := c.String(strings.ReplaceAll(strings.ToLower(key), "_", "-"))
	value := r.Header.Get(formattedKey)
	if value == "" {
		return "", errors.New("header not found or empty")
	}
	return value, nil
}

func QueryParamParser(r *http.Request, key string) (string, error) {
	params := r.URL.Query()
	formattedKey := strings.ReplaceAll(strings.ToLower(key), "-", "_")
	value := params.Get(formattedKey)
	if value == "" {
		return "", errors.New("query parameter not found or empty")
	}
	return value, nil
}

func PathParser(r *http.Request, key string) (string, error) {
	formattedKey := strings.ReplaceAll(strings.ToLower(key), "-", "_")
	v := r.PathValue(formattedKey)
	if v == "" {
		return "", errors.New("path parameter not found or empty")
	}
	return v, nil
}

func CookieParser(r *http.Request, key string) (string, error) {
	formattedKey := strings.ReplaceAll(strings.ToLower(key), "-", "_")
	cookie, err := r.Cookie(formattedKey)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", errors.New("cookie not found")
		}
		return "", err
	}
	return cookie.Value, nil
}

func CombinedParser(list ...Parser) Parser {
	return func(r *http.Request, key string) (string, error) {
		for _, parser := range list {
			parsed, err := parser(r, key)
			if err == nil && len(parsed) > 0 {
				return parsed, nil
			}
		}
		return "", fmt.Errorf("no parser found for key %s", key)
	}
}
