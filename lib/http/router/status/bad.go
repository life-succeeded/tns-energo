package status

import (
	"net/http"
	"tns-energo/lib/apierr"
	"tns-energo/lib/http/router"
)

func BadRequestHandler(c router.Context) error {
	return c.WriteJson(http.StatusBadRequest,
		router.NewErrorResponse(apierr.CodeUnknownError, "Произошла неизвестная ошибка"))
}

func UnauthorizedHandler(c router.Context) error {
	return c.WriteJson(http.StatusUnauthorized,
		router.NewErrorResponse(apierr.CodeUnauthorized, "Авторизация не выполнена"))
}

func ForbiddenHandler(c router.Context) error {
	return c.WriteJson(http.StatusForbidden,
		router.NewErrorResponse(apierr.CodeForbidden, "Доступ запрещен"))
}
