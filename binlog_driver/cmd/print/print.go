package main

import (
	"flag"

	"github.com/Shopify/sarama"
	"github.com/ngaut/log"
	"github.com/pingcap/tidb-tools/binlog_driver/reader"
)

var (
	offset    = flag.Int64("offset", sarama.OffsetNewest, "offset")
	commitTS  = flag.Int64("commitTS", 0, "commitTS")
	clusterID = flag.String("clusterID", "6561373978432450126", "clusterID")
)

func main() {
	flag.Parse()

	cfg := &reader.Config{
		KafakaAddr: []string{"127.0.0.1:9092"},
		Offset:     *offset,
		CommitTS:   *commitTS,
		ClusterID:  *clusterID,
	}

	breader, err := reader.NewReader(cfg)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-breader.Messages():
			log.Debug("recv: ", msg.Binlog.String())
		}
	}
}