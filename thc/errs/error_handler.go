package errs

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type ServiceError struct {
	ErrCode    string `json:"err_code"`
	ErrMessage string `json:"err_message"`
}

func (se ServiceError) Error() string {
	return fmt.Sprintf("[%s]%s", se.ErrCode, se.ErrMessage)
}

func DefaultHttpErrHandler(err error, ctx echo.Context) {

	if _, ok := err.(*ServiceError); ok {
		ctx.JSON(200, err)
	} else if he, ok := err.(*echo.HTTPError); ok {
		ctx.JSON(he.Code, err)
	} else {
		ctx.JSON(http.StatusInternalServerError, "系统异常")
	}

}
