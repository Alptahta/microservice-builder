package main

import "flag"

// InitPath is the path for where to create microservice.
const InitPath = "../"

// DefaultServiceName is the default value for service name if service name is not provided by user
const DefaultServiceName = "default"

// ServiceName is the service name given by the user.
var ServiceName string

func init() {
	flag.StringVar(&ServiceName, "service", DefaultServiceName, "service name")
}

func main() {

}
