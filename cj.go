package cj

import "time"

type CronJob struct {
	Name       string
	Interval   time.Duration
	ExpireTime time.Time
	Do         func()
}

func NewCronJob(name string, interval time.Duration, expireTime time.Time, do func()) *CronJob {
	return &CronJob{
		Name:       name,
		Interval:   interval,
		ExpireTime: expireTime,
		Do:         do,
	}
}

func (c *CronJob) Start() {
	tik := time.NewTicker(c.Interval)
	for true {
		select {
		case now := <-tik.C:
			if c.ExpireTime.Sub(now) > 0 {
				c.Do()
			} else {
				return
			}
		}
	}
}
