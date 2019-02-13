// The entrypoint for the CLI client
// This is a simple helper tool for interacting with the Packet API
// It's a first step in what would be needed to make a feature-rich
// cli tool for Packet. Enjoy!
package main

import (
	"flag"
	"os"
	"reflect"

	"github.com/golang/glog"
	"github.com/urfave/cli"

	"github.com/packethost/packngo"

	"pkg/packet"
)

// Returns the cli Flags for a user to specify a New Device Request
// Note
//   Not all attributes are exposed (i.e. lists) but can be added later.
//   This depends on reflection which is horribly difficult to maintain,
//   error check, and is inefficient. Unfortunately Golang as language
//   doesn't provide many options here for dynamically listing struct fields.
//   This is a good indicator that Go may not be the best language for this
//   specific use case.
func newDeviceRequestOptions() []cli.Flag {

	val := reflect.Indirect(reflect.ValueOf(packngo.DeviceCreateRequest{}))
	flags := []cli.Flag{}

	for i := 0; i < val.NumField(); i++ {

		fieldType := val.Field(i).Type().Name()
		fieldName := val.Type().Field(i).Name
		var flag cli.Flag

		switch fieldType {
		case "string":
			flag = &cli.StringFlag{
				Name:  fieldName,
				Usage: "",
			}
		case "bool":
			flag = &cli.BoolFlag{
				Name:  fieldName,
				Usage: "boolean",
			}
		case "float32":
			flag = &cli.Float64Flag{
				Name: fieldName,
			}
		case "float64":
			flag = &cli.Float64Flag{
				Name: fieldName,
			}
		default:
			// TODO: Add options for other flag types
			continue
		}
		flags = append(flags, flag)
	}

	return flags
}

// Converts a cli context into a New Device Request
func newDeviceRequest(ctx *cli.Context) *packngo.DeviceCreateRequest {
	device := &packngo.DeviceCreateRequest{}

	for _, field := range ctx.FlagNames() {

		val := reflect.ValueOf(device).Elem().FieldByName(field)
		fieldType := val.Type().Name()

		switch fieldType {
		case "string":
			val.SetString(ctx.String(field))

		case "bool":
			val.SetBool(ctx.Bool(field))

		case "float32":
			val.SetFloat(ctx.Float64(field))

		case "float64":
			val.SetFloat(ctx.Float64(field))
		default:
			// TODO: Add fields for other option types
		}
	}

	return device
}

func run() error {
	app := cli.NewApp()
	app.Name = "Packet CLI Demo"
	app.Usage = "Let's you query, create, and remove machines!"

	// All subcommands that operate on `device` API
	deviceSubcommands := []cli.Command{
		{
			Name:  "create",
			Usage: "create a device",
			Flags: newDeviceRequestOptions(),
			Action: func(ctx *cli.Context) error {
				return packet.CreateDevice(newDeviceRequest(ctx))
			},
		},
		{
			Name:  "delete",
			Usage: "delete a device",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Device UUID",
				},
				cli.BoolFlag{
					Name:  "force_delete",
					Usage: "Force the deletion of the device, by detaching any storage volume still active.",
				},
			},
			Action: func(ctx *cli.Context) error {
				return packet.DeleteDevice(ctx.String("id"))
			},
		},
		{
			Name: "list",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Project UUID",
				},
				cli.IntFlag{
					Name:  "page",
					Usage: "page to display, default to 1, max 100_000",
				},
				cli.IntFlag{
					Name:  "per_page",
					Usage: "items per page, default to 10, max 1_000",
				},
			},
			Usage: "Provides a collection of devices for a given project.",
			Action: func(ctx *cli.Context) error {
				opts := &packngo.ListOptions{
					Page:    ctx.Int("page"),
					PerPage: ctx.Int("per_page"),
					// TODO: Support more options
				}
				return packet.ListDevices(ctx.String("id"), opts)
			},
		},
	}

	// The CLI commands
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "lists the available projects",
			// the action, or code that will be executed when
			// we execute our `list` command
			Action: func(ctx *cli.Context) error {
				return packet.ListProjects()
			},
		},
		{
			Name:  "device",
			Usage: "options for device operations",
			// the subcommands, or code that will be executed when
			// we execute our `device` command
			Subcommands: deviceSubcommands,
		},
	}

	if err := app.Run(os.Args); err != nil {
		glog.Error(err)
	}

	return nil
}

func main() {
	flag.Parse()
	glog.Info("started")

	if err := run(); err != nil {
		glog.Fatal(err)
	}

	glog.Info("finished")
}
