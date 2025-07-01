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

func GetHomestayHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Parse query param (?id=...)
		request := &types.HomestayGetRequest{}
		err := http_request.BindUri(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		// Call logic
		getLogic := homestay.NewHomestayGetLogic(ctx, svcCtx)
		result, err := getLogic.HomestayGet(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
