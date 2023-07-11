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
	//根据Minute,hour 对应的设置的时间来运行
	t.s.Every(1).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1m")
	}) //代表每两秒运行一次
	t.s.Every(3).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("3m")
	}) //代表每两秒运行一次
	t.s.Every(5).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("5m")
	}) //代表每两秒运行一次
	t.s.Every(15).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("15m")
	}) //代表每两秒运行一次
	t.s.Every(30).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("30m")
	}) //每三十分分钟运行

	t.s.Every(1).Hour().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1H")
	})
	t.s.Every(2).Hour().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("2H")
	})
	t.s.Every(4).Hour().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("4H")
	})
	t.s.Every(1).Day().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1D")
	})
	t.s.Every(1).Week().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1W")
	})
	t.s.Every(1).Month().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1M")
	})
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}

func (t *Task) Stop() {
	t.s.Stop()
}
