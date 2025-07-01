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

func CreateHomestayHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		hsLogic := homestay.NewHomestayCreateLogic(ctx, svcCtx)

		request := &types.CreateHomestayRequest{}
		err := http_request.BindBodyJson(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		result, err := hsLogic.HomestayCreate(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
