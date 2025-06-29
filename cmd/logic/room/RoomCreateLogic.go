package topic

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HomestayCreateLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewTopicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) HomestayCreateLogic {
	return HomestayCreateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *HomestayCreateLogic) HomestayCreate(input *types.C) error {
	l.logHelper.Infof("Start process get topic")

	topic := &model.Topic{
		TopicName:      input.TopicName,
		GroupStudentID: input.GroupStudentID,
		LecturerID:     input.LecturerID,
		StartDay:       input.StartDay,
		EndDay:         input.EndDay,
		Allowance:      input.Allowance,
		TopicStatus:    input.TopicStatus,
	}

	err := l.svcCtx.TopicRepo.CreateTopic(l.ctx, topic)
	if err != nil {
		l.logHelper.Errorf("Failed while insert topic, error: %s", err.Error())
		return err
	}
	return nil

}
