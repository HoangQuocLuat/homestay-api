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

func UpdateHomestayHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		homestayUpdateLogic := homestay.NewHomestayUpdateLogic(ctx, svcCtx)
		path := &types.UpdateHomestayPath{}
		err := http_request.BindUri(c, path)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		request := &types.UpdateHomestayRequest{}
		err = http_request.BindBodyJson(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		input := &types.UpdateHomestayInput{
			Path:    path,
			Request: request,
		}

		result, err := homestayUpdateLogic.HomestayUpdate(input)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
