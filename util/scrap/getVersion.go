package scrap

import (
	"fmt"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
)

type Version struct {
	Hostname string
	Platform string
	Version  string
	Uptime   string
}

func GetVersion() ([]map[string]interface{}, error) {
	targets := getHost()

	var results []map[string]interface{}

	for _, target := range targets {
		p, err := platform.NewPlatform(
			target.Platform,
			target.Hostname,
			options.WithAuthNoStrictKey(),
			options.WithAuthUsername(target.Username),
			options.WithAuthPassword(target.Password),
			// options.WithSSHConfigFile("ssh_config"),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to create platform for %s: %+v\n", target.Hostname, err)
		}

		d, err := p.GetNetworkDriver()
		if err != nil {
			return nil, fmt.Errorf("failed to create driver for %s: %+v\n", target.Hostname, err)
		}

		err = d.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open driver for %s: %+v\n", target.Hostname, err)
		}
		defer d.Close()

		response, err := d.SendCommand("show version")
		if err != nil {
			return nil, fmt.Errorf("failed to send command for %s: %+v\n", target.Hostname, err)
		}

		// return response.Result, nil

		parsedOut, err := response.TextFsmParse(TEMPLATES_PATH + "cisco_nxos_show_version.textfsm")
		if err != nil {
			return nil, fmt.Errorf("failed to parse command for %s: %+v\n", target.Hostname, err)
		}
		fmt.Println(parsedOut)
		parsedOut[0]["HOSTNAME"] = target.Hostname

		results = append(results, parsedOut...)
	}

	return results, nil
}
