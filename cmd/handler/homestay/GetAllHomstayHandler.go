package homestay

import (
	"back-end/cmd/logic/homestay"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"back-end/core/http_request"
	"back-end/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllHomestayHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := c.Request.Context()

		// New object logic (all logic code will implement in this object)
		getAllLogic := homestay.NewHomestayGetAllLogic(ctx, svcCtx)

		request := &types.HomestayGetAllRequest{}
		err := http_request.BindQueryString(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		result, err := getAllLogic.HomestayGetAll(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
