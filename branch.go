package golib

import "time"

func RetryDelay(proc func() error, times int, interval time.Duration) error {
	var err error
	for retry := 0; retry < times; retry++ {
		if err = proc(); err == nil {
			return nil
		}
		time.Sleep(interval)
	}
	return err
}

/* 将func重试times次, 默认1s间隔 */
func Retry(proc func() error, times int) error {
	return RetryDelay(proc, times, time.Second)
}
