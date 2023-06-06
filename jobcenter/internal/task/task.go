package task

import (
	"github.com/go-co-op/gocron"
	"jobcenter/internal/kline"
	"jobcenter/internal/svc"
	"time"
)

type Task struct {
	s   *gocron.Scheduler
	ctx *svc.ServiceContext
}

func NewTask(ctx *svc.ServiceContext) *Task {
	return &Task{
		s:   gocron.NewScheduler(time.UTC),
		ctx: ctx,
	}
}

func (t *Task) Run() {
	t.s.Every(2).Second().Do(func() {
		kline.NewKline(t.ctx.Config.Okx).Do("1m")
	}) //代表每两秒运行一次

	t.s.Every(2).Hour().Do(func() {
		kline.NewKline(t.ctx.Config.Okx).Do("1H")
	}) //代表每两秒运行一次
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}

func (t *Task) Stop() {
	t.s.Stop()
}
