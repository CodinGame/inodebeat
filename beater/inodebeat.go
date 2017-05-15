package beater

import (
	"fmt"
	"syscall"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/codingame/inodebeat/config"
)

type Inodebeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Inodebeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Inodebeat) Run(b *beat.Beat) error {
	logp.Info("inodebeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		inodes := getInodes()

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
			"inodes":     inodes,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Inodebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func getInodes() common.MapStr {
	var stat syscall.Statfs_t

	err := syscall.Statfs("/", &stat)

	if err != nil {
		logp.Err(err.Error())
	}

	inodesTotal := stat.Files
	inodesFree := stat.Ffree
	inodesUsed := inodesTotal - inodesFree
	inodesFreePercent := 100.0 * float64(inodesFree) / float64(inodesTotal)
	inodesUsedPercent := 100.0 * float64(inodesUsed) / float64(inodesTotal)

	inodes := common.MapStr{
		"total": inodesTotal,
		"free": common.MapStr{
			"total": inodesFree,
			"pct":   inodesFreePercent,
		},
		"used": common.MapStr{
			"total": inodesUsed,
			"pct":   inodesUsedPercent,
		},
	}

	return inodes
}
