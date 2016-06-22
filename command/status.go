package command

import "github.com/urfave/cli"

func CmdStatus(c *cli.Context) error {
	// Write your code here
	return nil
}

// // RoadieLogEntry is a structure of one log entry.
// type RoadieLogEntry struct {
// 	Timestamp time.Time
// 	Payload   struct {
// 		Username     string
// 		Stream       string
// 		Log          string
// 		ContainerID  string `mapstructure:"container_id"`
// 		InstanceName string `mapstructure:"instance_name"`
// 	}
// }
//
//
// func Logging(project, name string) {
//
// 	ch := make(chan *RoadieLogEntry)
// 	chErr := make(chan error)
// 	go getLogEntries(project, name, ch, chErr)
//
// 	fmt.Println("Current Logs:")
// loop:
// 	for {
// 		select {
// 		case entry := <-ch:
// 			if entry == nil {
// 				break loop
// 			}
// 			fmt.Printf("  %v: %s\n", entry.Timestamp.Format("2006/01/02 15:04:05"), entry.Payload.Log)
// 		case err := <-chErr:
// 			fmt.Println(err.Error())
// 			break loop
// 		}
// 	}
//
// }
//
// func getLogEntries(project, name string, ch chan *RoadieLogEntry, chErr chan error) {
//
// 	client, err := google.DefaultClient(context.Background(), logging.CloudPlatformReadOnlyScope)
// 	if err != nil {
// 		chErr <- err
// 		return
// 	}
//
// 	service, err := logging.New(client)
// 	if err != nil {
// 		chErr <- err
// 		return
// 	}
//
// 	pageToken := ""
// 	for {
//
// 		res, err := service.Entries.List(&logging.ListLogEntriesRequest{
// 			ProjectIds: []string{project},
// 			Filter: fmt.Sprintf(
// 				"resource.type = \"gce_instance\" AND logName = \"projects/%s/logs/%s\"",
// 				project, name),
// 			PageToken: pageToken,
// 		}).Do()
// 		if err != nil {
// 			chErr <- err
// 			return
// 		}
//
// 		for _, v := range res.Entries {
// 			if v.JsonPayload != nil {
//
// 				timestamp, err := time.Parse("2006-01-02T15:04:05.000Z", v.Timestamp)
// 				if err != nil {
// 					log.Println(chalk.Red.Color(err.Error()))
// 					continue
// 				}
//
// 				entry := &RoadieLogEntry{Timestamp: timestamp}
// 				if err := mapstructure.Decode(v.JsonPayload, &entry.Payload); err != nil {
// 					log.Println(chalk.Red.Color(err.Error()))
// 					continue
// 				}
//
// 				entry.Payload.Log = strings.TrimRight(entry.Payload.Log, "\n")
// 				ch <- entry
//
// 			}
// 		}
//
// 		pageToken = res.NextPageToken
// 		if pageToken == "" {
// 			break
// 		}
//
// 	}
//
// 	ch <- nil
//
// }
