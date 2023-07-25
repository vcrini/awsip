package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vcrini/go-utils"
)

func main() {
	host := flag.String("host", "*", "Name of hosts or wildcard")
	h := flag.Bool("h", false, "Display help")
	instanceStateName := flag.String("instance-state-name", "running", "The state of the instance (pending | running | shutting-down | terminated | stopping | stopped | all ).")
	// here all flags
	flag.Parse()
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
