package scheduler

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	cron *cron.Cron
	log  *logrus.Logger
}

func NewScheduler(log *logrus.Logger) *Scheduler {
	location, _ := time.LoadLocation("Asia/Jakarta")

	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)

	s := cron.New(
		cron.WithParser(parser),
		cron.WithLocation(location),
	)

	return &Scheduler{
		cron: s,
		log:  log,
	}
}

func (s *Scheduler) AddJob(cronExpr string, handler func()) error {
	_, err := s.cron.AddFunc(cronExpr, handler)
	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
	s.log.Info("Scheduler Started")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	s.log.Info("Scheduler stopped")
}
