package command

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

func CmdStatus(c *cli.Context) error {
	// Write your code here
	// conf := GetConfig(c)

	return nil
}

const (
	eventSubtypeInsert = "compute.instances.insert"
	eventSubtypeDelete = "compute.instances.delete"
)

func Status() {

	ch := make(chan *LogEntry)
	chErr := make(chan error)

	go GetLogEntries("jkawamoto-ppls",
		"jsonPayload.event_type = \"GCE_OPERATION_DONE\"", ch, chErr)

	runnings := make(map[string]bool)

loop:
	for {
		select {
		case entry := <-ch:

			if entry == nil {
				break loop
			}

			if payload, err := getActivityPayload(entry); err == nil {

				switch payload.EventSubtype {
				case eventSubtypeInsert:
					runnings[payload.Resource.Name] = true
				case eventSubtypeDelete:
					runnings[payload.Resource.Name] = false
				}

			} else {
				log.Println(chalk.Red.Color(err.Error()))
			}

		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

	fmt.Println("Instances:")
	for name, status := range runnings {
		if status {
			fmt.Printf("  %s running\n", name)
		} else {
			fmt.Printf("  %s stop\n", name)
		}

	}

}

func getActivityPayload(entry *LogEntry) (*ActivityPayload, error) {

	var res ActivityPayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type ActivityPayload struct {
	EventTimestampUs string `mapstructure:"event_timestamp_us"`
	EventType        string `mapstructure:"vent_type"`
	TraceID          string `mapstructure:"trace_id"`
	Actor            struct {
		User string
	}
	Resource struct {
		Zone string
		Type string
		ID   string
		Name string
	}
	Version      string
	EventSubtype string `mapstructure:"event_subtype"`
	Operation    struct {
		Zone string
		Type string
		ID   string
		Name string
	}
}

// {
//   metadata: {
//     severity: "INFO"
//     projectId: "jkawamoto-ppls"
//     serviceName: "compute.googleapis.com"
//     zone: "us-central1-b"
//     labels: {
//       compute.googleapis.com/resource_zone: "us-central1-b"
//       compute.googleapis.com/resource_name: "mbp-2-test-20160622015847"
//       compute.googleapis.com/resource_type: "instance"
//       compute.googleapis.com/resource_id: "2770689903246678007"
//     }
//     timestamp: "2016-06-22T05:59:10.174129Z"
//     projectNumber: "1027771359384"
//   }
//   insertId: "rr2road991t"
//   log: "compute.googleapis.com/activity_log"
//   structPayload: {
//     event_timestamp_us: "1466575150174129"
//     event_type: "GCE_OPERATION_DONE"
//     trace_id: "operation-1466575128307-535d7a18f2538-ca003c59-c6fc4d23"
//     actor: {
// *     user: "kawamoto.junpei@gmail.com"
//     }
//     resource: {
//       zone: "us-central1-b"
//       type: "instance"
//       id: "2770689903246678007"
// *     name: "mbp-2-test-20160622015847"
//     }
//     version: "1.2"
// *   event_subtype: "compute.instances.insert"
//     operation: {
//       zone: "us-central1-b"
//       type: "operation"
//       id: "1906738067512203254"
//       name: "operation-1466575128307-535d7a18f2538-ca003c59-c6fc4d23"
//     }
//   }
// }
//
//
//
// {
//   metadata: {
//     severity: "INFO"
//     projectId: "jkawamoto-ppls"
//     serviceName: "compute.googleapis.com"
//     zone: "us-central1-b"
//     labels: {
//       compute.googleapis.com/resource_zone: "us-central1-b"
//       compute.googleapis.com/resource_name: "mbp-2-test-20160622015347"
//       compute.googleapis.com/resource_type: "instance"
//       compute.googleapis.com/resource_id: "7336954138552932610"
//     }
//     timestamp: "2016-06-22T05:56:54.141478Z"
//     projectNumber: "1027771359384"
//   }
//   insertId: "1bq7woqdn46e"
//   log: "compute.googleapis.com/activity_log"
//   structPayload: {
//     event_timestamp_us: "1466575014141478"
//     event_type: "GCE_OPERATION_DONE"
//     trace_id: "operation-1466574964969-535d797d2ce29-93022b89-aebe13a2"
//     actor: {
//       user: "1027771359384-uunducldfql6bsf9df8sjq57gefeng5d@developer.gserviceaccount.com"
//     }
//     resource: {
//       zone: "us-central1-b"
//       type: "instance"
//       id: "7336954138552932610"
//       name: "mbp-2-test-20160622015347"
//     }
//     version: "1.2"
//     event_subtype: "compute.instances.delete"
//     operation: {
//       zone: "us-central1-b"
//       type: "operation"
//       id: "3084257846255654554"
//       name: "operation-1466574964969-535d797d2ce29-93022b89-aebe13a2"
//     }
//   }
// }
