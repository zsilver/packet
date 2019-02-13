package packet

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/glog"

	"github.com/packethost/packngo"
)

func DeleteDevice(id string) error {
	c, err := packngo.NewClient()
	if err != nil {
		glog.Fatal(err)
	}

	res, err := c.Devices.Delete(id)
	if err != nil {
		return err
	}

	fmt.Printf("%q\n", res)
	return nil
}

func CreateDevice(request *packngo.DeviceCreateRequest) error {
	c, err := packngo.NewClient()
	if err != nil {
		glog.Fatal(err)
	}

	request.Facility = []string{"any"}

	dev, res, err := c.Devices.Create(request)
	if err != nil {
		glog.Fatal(err)
	}

	msg, err := json.MarshalIndent(dev, "", "  ")
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Printf("%s\n%q\nSuccess!\n", msg, res)

	return nil
}

func ListDevices(id string, options *packngo.ListOptions) error {
	if id == "" {
		return errors.New("`--id` must be set")
	}

	c, err := packngo.NewClient()
	if err != nil {
		glog.Fatal(err)
	}

	ds, _, err := c.Devices.List(id, options)
	if err != nil {
		glog.Fatal(err)
	}

	fmt.Printf("%s %9s %42s\n", "State", "ID", "Hostname")
	for _, d := range ds {
		fmt.Println(d.State, d.ID, d.Hostname)
	}

	if len(ds) == 0 {
		fmt.Println("No devices exist")
	}

	return nil
}

func ListProjects() error {
	c, err := packngo.NewClient()
	if err != nil {
		glog.Fatal(err)
	}

	ps, _, err := c.Projects.List(nil)
	if err != nil {
		glog.Fatal(err)
	}

	fmt.Printf("%s %38s\n", "ID", "Name")
	for _, p := range ps {
		fmt.Println(p.ID, p.Name)
	}
	return nil
}
