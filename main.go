package main

import (
	"flag"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
)

func main() {
	version := 1
	number := 1
	domainUUID := "3c1a7f99-ebfd-46fc-ae4a-3652ec6d1949"
	domainName := "example.com"
	flag.IntVar(&version, "v", 4, "UUID version could be v4 and v3")
	flag.IntVar(&number, "n", 1, "Number of UUIDs to be generated")
	flag.StringVar(&domainUUID, "u", "3c1a7f99-ebfd-46fc-ae4a-3652ec6d1949", "Namespace uuid to be used for v3, v5")
	flag.StringVar(&domainName, "d", "example.com", "Domain Name to be used for v3, v5")
	flag.Parse()
	
	log.SetFlags(0)
	
	fromUUID, err := uuid.FromString(domainUUID)
	if err != nil {
		log.Fatal("Error: Provided namespace UUID is not a valid uuid")
	}
	
	if version != 4 {
		if number > 1 {
			fmt.Println("Based on type of UUID only 1 record will be generated")
		}
		switch version {
		case 1:
			fmt.Println("UUID based on current timestamp and MAC address.")
			fmt.Println(uuid.NewV1())
		case 2:
			fmt.Println("Person:", uuid.NewV2(uuid.DomainPerson))
			fmt.Println("Group :", uuid.NewV2(uuid.DomainGroup))
			fmt.Println("Org   :", uuid.NewV2(uuid.DomainOrg))
		case 3:
			fmt.Println("UUID based on MD5 hash of namespace UUID and name.")
			fmt.Println(uuid.NewV3(fromUUID, domainName))
		case 5:
			fmt.Println("UUID based on SHA-1 hash of namespace UUID and name.")
			fmt.Println(uuid.NewV5(fromUUID, domainName))
		}
	} else {
		for i := 1; i <= number; i++ {
			fmt.Println(uuid.NewV4())
		}
	}
}
