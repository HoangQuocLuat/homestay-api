// logic/homestay/get.go

package homestay

import (
	"back-end/cmd/database/repo"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"
	"database/sql"
)

type HomestayGetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayGetLogic {
	return HomestayGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayGetLogic) HomestayGetAll(input *types.HomestayGetAllRequest) (*types.HomestayGetAllResponse, error) {
	condition := &repo.HomestayCondition{
		Location:  input.Location,
		HostID:    input.HostID,
		ServiceID: input.ServiceID,
		Status:    input.Status,
	}

	homestays, total, err := l.svcCtx.HomestayRepo.GetHomestays(l.ctx, condition, input.Page, input.PageSize)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	var dataResponse []types.Homestay
	for _, h := range homestays {
		dataResponse = append(dataResponse, types.Homestay{
			Id:            h.Id,
			ServiceID:     h.ServiceID,
			HostID:        h.HostID,
			Name:          h.Name,
			Description:   h.Description,
			Location:      h.Location,
			Address:       h.Address,
			CoverImageURL: h.CoverImageURL,
			GalleryImages: h.GalleryImages,
			Status:        h.Status,
		})
	}

	totalPages := (total + input.PageSize - 1) / input.PageSize

	return &types.HomestayGetAllResponse{
		Raw: dataResponse,
		PaginationResponse: types.PaginationResponse{
			Page:       input.Page,
			PageSize:   input.PageSize,
			TotalItems: total,
			TotalPages: totalPages,
		},
	}, nil
}
