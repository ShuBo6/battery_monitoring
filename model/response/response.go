package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR       = 7
	NotExist    = 6
	Unverified  = 3
	Authority   = 4
	SUCCESS     = 0
	ExportLimit = 9
)

type Response struct {
	Code        int         `json:"code"`
	MessageCode int         `json:"-"`
	Data        interface{} `json:"data"`
	Error       error       `json:"-"`
	Err         string      `json:"err,omitempty"`
	Message     string      `json:"msg"`
}

type (
	empty   struct{}
	Handler struct{}
)

type handler func(c *gin.Context) *Response

func (h *Handler) Handler() func(handler handler) gin.HandlerFunc {
	return func(handler handler) gin.HandlerFunc {
		return func(context *gin.Context) {
			response := handler(context)
			if response == nil {
				return
			}
			if response.Data == nil {
				response.Data = empty{}
			}
			if response.Error != nil {
				response.Code = ERROR
				response.Err = response.Error.Error()
				response.Message = response.Err
				_ = context.Error(response.Error)
			}
			if response.Code == ERROR && context.Err() == nil {
				_ = context.Error(errors.New(response.Message))
			}
			context.JSON(http.StatusOK, response)
		}
	}
}
