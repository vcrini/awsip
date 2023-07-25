package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vcrini/go-utils"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func main() {
	h := flag.Bool("h", false, "Display help")
	host := flag.String("host", "*", "Name of hosts or wildcard")
	instanceStateName := flag.String("instance-state-name", "running", "The state of the instance (pending | running | shutting-down | terminated | stopping | stopped | all ).")
	v := flag.Bool("v", false, "if true, print version and exit")
	// here all flags
	flag.Parse()
	if *v {
		fmt.Printf("Build on %s from sha1 %s\n", buildTime, sha1ver)
		os.Exit(0)
	}
	if *h {
		flag.Usage()
		os.Exit(0)
	}
	var filter string
	if *instanceStateName == "all" {
		filter = fmt.Sprintf(`[{"Name": "tag:Name","Values":["%s"]}]`, *host)
	} else {
		filter = fmt.Sprintf(`[{"Name": "tag:Name","Values":["%s"]},{"Name":"instance-state-name","Values":["%s"]}]`, *host, *instanceStateName)
	}
	buildCommand := []string{"aws", "ec2", "describe-instances", "--filters", filter, "--query", "Reservations[].Instances[].{id:InstanceId,name:Tags[?Key == 'Name'].Value | [0],ip:PrivateIpAddress,az:Placement.AvailabilityZone}"}
	fmt.Println(utils.Exe(buildCommand))
}
