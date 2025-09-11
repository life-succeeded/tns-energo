package cron

import (
	"context"
	"fmt"
	"time"
	"tns-energo/config"
	"tns-energo/service/analytics"

	"github.com/go-co-op/gocron/v2"
	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type Service struct {
	scheduler        gocron.Scheduler
	settings         config.Settings
	analyticsService *analytics.Service
}

func NewService(settings config.Settings, analyticsService *analytics.Service) *Service {
	return &Service{
		settings:         settings,
		analyticsService: analyticsService,
	}
}

func (s *Service) LaunchJobs(ctx context.Context, log golog.Logger) error {
	log = log.WithTags("cron")

	dailyReportTime, err := time.Parse("15:04", s.settings.Cron.DailyReportTime)

	s.scheduler, err = gocron.NewScheduler(gocron.WithLocation(time.UTC), gocron.WithLogger(logger{log}))
	if err != nil {
		return fmt.Errorf("could not start scheduler: %w", err)
	}

	reportJob, err := s.scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(
			gocron.NewAtTime(uint(dailyReportTime.Hour()), uint(dailyReportTime.Minute()), 0),
		)),
		gocron.NewTask(s.dailyReportTask, ctx, log),
	)
	if err != nil {
		return fmt.Errorf("could not start reporting job: %w", err)
	}

	s.scheduler.Start()
	log.Debugf("started reporting job: %s", reportJob.ID())

	return nil
}

func (s *Service) Shutdown() error {
	if err := s.scheduler.Shutdown(); err != nil {
		return fmt.Errorf("could not shutdown cron scheduler: %w", err)
	}

	return nil
}

func (s *Service) dailyReportTask(ctx context.Context, log golog.Logger) {
	taskCtx, cancelTaskCtx := goctx.Wrap(ctx).WithCancel()
	defer cancelTaskCtx()

	log = log.WithTags("daily report task")

	report, err := s.analyticsService.GenerateDailyReport(taskCtx, log, time.Now())
	if err != nil {
		log.Errorf("could not generate daily report: %v", err)
		return
	}

	log.Debugf("daily report '%s' generated at %v", report.File.Name, report.CreatedAt)
}
