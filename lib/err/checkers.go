package err

import (
	"context"
	"errors"
)

// AsError возвращает ошибку в указанном типе или nil, если привести её к этом типу не удалось.
// Является оберткой для сокращения кода при использовании стандартного метода errors.As.
//
// Пример:
//
//	if e, ok := sw.AsError[CustomError](err); ok {
//	  log.Println(e.CustomMethod())
//	  response.Status(http.StatusNotFound)
//	} else if e, ok := sw.AsError[ValidationError](err); ok {
//
//	  response.Status(e.HttpStatus())
//	}
func AsError[E error](err error) (E, bool) {
	var castError E
	return castError, errors.As(err, &castError)
}

// IsContextError
// Проверяет, является ли ошибка связанной с работой контекстов context.Context / swctx.Context
func IsContextError(err error) bool {
	return errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)
}

// IsCriticalError
// Проверяет, является ли ошибка критичной (обязательной для логирования, например).
// Под критичные ошибки не попадают: предупреждения Warning, ошибки NotFoundError и ошибки контекста, проверяемые через IsContextError
func IsCriticalError(err error) bool {
	_, isWarn := AsError[Warning](err)
	if isWarn {
		return false
	}
	_, isNotFound := AsError[NotFoundError](err)
	if isNotFound {
		return false
	}

	return !IsContextError(err)
}
