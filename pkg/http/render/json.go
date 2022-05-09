package render

import (
	stdjson "encoding/json"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// JSONSerializer uses json-iterator.
type JSONSerializer struct{}

// WithJSONSerializer sets custom JSON serializer for the provided echo instance.
func WithJSONSerializer(e *echo.Echo) {
	e.JSONSerializer = JSONSerializer{}
}

// Serialize converts any value into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (JSONSerializer) Serialize(c echo.Context, val any, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}

	return enc.Encode(val)
}

// Deserialize reads a JSON from a request body and converts it into an any value.
func (JSONSerializer) Deserialize(c echo.Context, val any) error {
	err := json.NewDecoder(c.Request().Body).Decode(val)
	if ute, ok := err.(*stdjson.UnmarshalTypeError); ok { //nolint:errorlint // expect exactly the specified error
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("unmarshal type error: expected=%v, got=%v, field=%v, offset=%v",
				ute.Type, ute.Value, ute.Field, ute.Offset),
		).SetInternal(err)
	}

	if se, ok := err.(*stdjson.SyntaxError); ok { //nolint:errorlint // expect exactly the specified error
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("syntax error: offset=%v, error=%v", se.Offset, se.Error()),
		).SetInternal(err)
	}

	return err
}
