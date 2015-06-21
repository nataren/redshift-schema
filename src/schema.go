package main

import (
	"log"

	rs "github.com/aws/aws-sdk-go/service/redshift"
)

func main() {
	log.Printf("%+v", rs.Redshift{})
}
