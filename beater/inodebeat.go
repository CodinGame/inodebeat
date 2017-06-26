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
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		for _, directory := range bt.config.Directories {
			inodes := getInodes(directory)

			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       b.Name,
				"inodes":     inodes,
			}
			bt.client.PublishEvent(event)
			logp.Info("Event sent")
		}
	}
}

func (bt *Inodebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func getInodes(directory string) common.MapStr {
	var stat syscall.Statfs_t

	err := syscall.Statfs(directory, &stat)

	if err != nil {
		errorMessage := err.Error()
		logp.Err(errorMessage)
		return common.MapStr{
			"directory": directory,
			"error":     errorMessage,
		}
	}

	inodesTotal := stat.Files
	inodesFree := stat.Ffree
	inodesUsed := inodesTotal - inodesFree
	inodesFreePercent := 100.0 * float64(inodesFree) / float64(inodesTotal)
	inodesUsedPercent := 100.0 * float64(inodesUsed) / float64(inodesTotal)

	inodes := common.MapStr{
		"directory": directory,
		"total":     inodesTotal,
		"free": common.MapStr{
			"count": inodesFree,
			"pct":   inodesFreePercent,
		},
		"used": common.MapStr{
			"count": inodesUsed,
			"pct":   inodesUsedPercent,
		},
	}

	return inodes
}
