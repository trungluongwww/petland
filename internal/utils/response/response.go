package response

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func R200(c echo.Context, data interface{}) error {
	return response(c, http.StatusOK, data, "thành công")
}

func R400(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrBadRequest
	}

	res := FindError(error)

	return response(c, http.StatusBadRequest, data, res.Message)
}

func R401(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrUnauthorized
	}

	res := FindError(error)

	return response(c, http.StatusBadRequest, data, res.Message)
}

func R404(c echo.Context, data interface{}, error string) error {
	if error == "" {
		error = ErrNotFound
	}

	res := FindError(error)

	return response(c, http.StatusBadRequest, data, res.Message)
}

func RouterResponse(ctx echo.Context, err error) error {
	var (
		invalidValidationError *validator.InvalidValidationError
	)
	if !errors.As(err, &invalidValidationError) {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() != "" && err.Error() != "" {
				return response(ctx, http.StatusBadRequest, nil, fmt.Sprintf(errParamFormat, err.Field()))
			}
		}
	}

	return R400(ctx, nil, err.Error())
}

func response(c echo.Context, statusCode int, data interface{}, err string) error {
	return c.JSON(statusCode, echo.Map{"data": data, "message": err})
}
