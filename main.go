package main

import (
	"context"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
)

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	host := "access.mainnet.nodes.onflow.org:9000"
	height := uint64(65264619)

	flowClient, err := grpc.NewClient(host)
	noErr(err)

	filters := flow.EventFilter{
		EventTypes: []string{
			"A.6f6702697b205c18.HWGaragePMV2.AdminMintCard",
			"A.6f6702697b205c18.HWGarageCardV2.DepositEvent",
		},
	}

	data, errChan, err := flowClient.SubscribeEventsByBlockHeight(ctx, height, filters)
	noErr(err)

	reconnect := func(height uint64) {
		fmt.Printf("Reconnecting at block %d\n", height)

		var err error
		flowClient, err = grpc.NewClient(host)
		noErr(err)

		data, errChan, err = flowClient.SubscribeEventsByBlockHeight(ctx, height, filters)
		noErr(err)
	}

	lastHeight := height
	for {
		select {
		case <-ctx.Done():
			return

		case eventData, ok := <-data:
			if !ok {
				if ctx.Err() != nil {
					return // graceful shutdown
				}
				// unexpected close
				reconnect(lastHeight + 1)
				continue
			}

			fmt.Printf("~~~ Height: %d: %d ~~~\n", eventData.Height, len(eventData.Events))
			//printEvents(eventData.Events)

			lastHeight = eventData.Height

		case err, ok := <-errChan:
			if !ok {
				if ctx.Err() != nil {
					return // graceful shutdown
				}
				// unexpected close
				reconnect(lastHeight + 1)
				continue
			}

			fmt.Printf("~~~ ERROR: %s ~~~\n", err.Error())
			reconnect(lastHeight + 1)
			continue
		}
	}
}

func printEvents(events []flow.Event) {
	for _, event := range events {
		fmt.Printf("\nType: %s\n", event.Type)
		fmt.Printf("Values: %v\n", event.Value)
		fmt.Printf("Transaction ID: %s\n", event.TransactionID)
	}
}
