package scheduler

import (
	"fmt"

	"github.com/faststackco/faststack/api/config"
	"github.com/termbox.bak/termbox/api/driver"
	redis "gopkg.in/redis.v5"
)

type LocalScheduler struct {
	redis         *redis.Client
	driverOptions *config.DriverOptions
}

func NewLocalScheduler(redis *redis.Client, options *config.DriverOptions) (Scheduler, error) {
	return &ConsulScheduler{
		redis:         redis,
		driverOptions: options,
	}, nil
}

func (c *LocalScheduler) Create(name, image, driverName string) error {
	driver, err := driver.NewDriver(name, *c.driverOptions)
	if err != nil {
		return err
	}

	if err := driver.Create(name, image); err != nil {
		return err
	}

	return c.redis.HMSet(fmt.Sprintf("machine:%s", name), map[string]string{
		"image":  image,
		"driver": driverName,
	}).Err()
}

func (c *LocalScheduler) Delete(name string) error {
	hash, err := c.redis.HGetAll(fmt.Sprintf("machine:%s", name)).Result()

	driverName, ok := hash["driver"]
	if !ok {
		return fmt.Errorf("machine '%s' does not exist", name)
	}

	driver, err := driver.NewDriver(driverName, *c.driverOptions)
	if err != nil {
		return err
	}

	if err := driver.Delete(name); err != nil {
		return err
	}

	return c.redis.HDel(fmt.Sprintf("machine:%s", name)).Err()
}
