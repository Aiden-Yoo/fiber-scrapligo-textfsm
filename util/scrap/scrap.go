package scrap

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Hostname  string `yaml:"hostname"`
	Platform  string `yaml:"platform"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	StrictKey bool   `yaml:"strictkey"`
}

type Inventory struct {
	Targets []Target `yaml:"target"`
}

const TEMPLATES_PATH = "./util/templates/"

type data struct {
	host     string
	platform string
	version  string
	uptime   string
}

func Scrap() {
	src, err := os.Open("./util/hosts.yaml")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	d := yaml.NewDecoder(src)

	var inv Inventory
	err = d.Decode(&inv)
	if err != nil {
		panic(err)
	}

	// getVersion(v)
	for _, v := range inv.Targets {
		fmt.Println(getVersion(v))
	}
}
