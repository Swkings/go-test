package main

import (
	"context"
	"sync"

	"gitlab.uisee.ai/gkl10385/scheduler-general-mq-protocol/mqaction"
	"gitlab.uisee.ai/gkl10385/scheduler-general-mq-protocol/mqcommon"
	"gitlab.uisee.ai/gkl10385/scheduler-general-mq-protocol/mqprotocol"

	"gitlab.uisee.ai/gkl10385/utc-utils/amq"
	"gitlab.uisee.ai/gkl10385/utc-utils/amqp"
	"gitlab.uisee.ai/gkl10385/utc-utils/logger"
)

var (
	conn                  *amqp.Connection
	ctx                   context.Context
	cancel                context.CancelFunc
	log                   *logger.Logger
	orderID, batchOrderID int64
	delay                 int32
	wg                    sync.WaitGroup
)

func init() {
	logger.SetDebug(false)
	logger.SetLevel(logger.TraceLevel)

	ctx, cancel = context.WithCancel(context.Background())
	log = logger.NewLogger("test")
	if c, err := amqp.NewConnection(ctx, "amqp://admin:admin@localhost:5672/", log); err != nil {
		log.Panicf("amqp connection initialize failed, %s", err.Error())
	} else {
		conn = c
	}

	orderID = 1
	batchOrderID = 2
	delay = 100

	initDelayConsumer()

	initBroadcastConsumer()
}

func initBroadcastConsumer() {
	broadcastConsumer, err := amq.NewBroadcastConsumer(string(mqcommon.Modules.Biz), conn, log, 1)
	if err != nil {
		log.Panicf("broadcast consumer initialize failed, %s", err.Error())
	}

	broadcastConsumer.Listen(func(b []byte) {
		payload := mqprotocol.Actions.Unmarshal(b)

		if !mqprotocol.Actions.Biz.Is(payload.Module) {
			return
		}

		switch {
		case mqprotocol.Actions.Biz.AddCronJob.Is(payload.Action):
			payload, _ := mqprotocol.Actions.Biz.AddCronJob.Unmarshal(b) //nolint: errcheck
			act, _ := payload.(*mqaction.PayloadCronAddJob)              //nolint: errcheck
			if act.IsValid() {
				log.Infof("broadcast consumer recv, %+v", act)
				wg.Done()
			}
		}
	})
}

func initDelayConsumer() {
	delayConsumer, err := amq.NewDelayConsumer(string(mqcommon.Modules.Order), conn, log, 1)
	if err != nil {
		log.Panicf("delay consumer initialize failed, %s", err.Error())
	}

	delayConsumer.Listen(func(b []byte) {
		payload := mqprotocol.Actions.Unmarshal(b)

		if !mqprotocol.Actions.Order.Is(payload.Module) {
			return
		}

		switch {
		case mqprotocol.Actions.Order.AppointBatchOrder.Is(payload.Action):
			payload, _ := mqprotocol.Actions.Order.AppointBatchOrder.Unmarshal(b) //nolint: errcheck
			act, _ := payload.(*mqaction.PayloadAppointBatchOrder)                //nolint: errcheck
			if act.IsValid() {
				log.Infof("delay consumer recv, %+v", act)

				cancel()
			}
		}
	})
}
