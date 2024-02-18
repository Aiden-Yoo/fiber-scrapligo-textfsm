package scrap

import (
	"fmt"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
)

func getVersion(r Target) (data, error) {
	p, err := platform.NewPlatform(
		r.Platform,
		r.Hostname,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(r.Username),
		options.WithAuthPassword(r.Password),
		// options.WithSSHConfigFile("ssh_config"),
	)

	nil_data := data{}
	if err != nil {
		return nil_data, fmt.Errorf("failed to create platform for %s: %+v\n", r.Hostname, err)
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		return nil_data, fmt.Errorf("failed to create driver for %s: %+v\n", r.Hostname, err)
	}

	err = d.Open()
	if err != nil {
		return nil_data, fmt.Errorf("failed to open driver for %s: %+v\n", r.Hostname, err)
	}
	defer d.Close()

	response, err := d.SendCommand("show version")
	if err != nil {
		return nil_data, fmt.Errorf("failed to send command for %s: %+v\n", r.Hostname, err)
	}

	// return response.Result, nil

	parsedOut, err := response.TextFsmParse(TEMPLATES_PATH + "cisco_nxos_show_version.textfsm")
	if err != nil {
		return nil_data, fmt.Errorf("failed to parse command for %s: %+v\n", r.Hostname, err)
	}
	fmt.Println(parsedOut)

	dt := data{
		host:     r.Hostname,
		platform: fmt.Sprintf("%s", parsedOut[0]["PLATFORM"]),
		version:  fmt.Sprintf("%s", parsedOut[0]["OS"]),
		uptime:   fmt.Sprintf("%s", parsedOut[0]["UPTIME"]),
	}
	return dt, nil
}
