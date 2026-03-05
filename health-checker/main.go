package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli"
)

func status(domain, port string) string {
	addr := domain + ":" + port
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", addr, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("Source: %s & Dest: %s - Is Not.", conn.LocalAddr().String(), conn.RemoteAddr())
	} else {
		status = fmt.Sprintf("Source: %s & Dest: %s - Is Reachable.", conn.LocalAddr().String(), conn.RemoteAddr())
	}
	defer conn.Close()
	return status
}

func main() {
	app := cli.NewApp()
	app.Name = "HealthChecker"
	app.Usage = "Health checker for checking the health of domain."
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "domain", Value: "google.com", Usage: "Domain name to check status"},
	}

	app.Action = func(c *cli.Context) error {
		port := c.String("port")
		if port == "" {
			port = "80"
		}
		domain := c.String("domain")
		status := status(domain, port)

		fmt.Printf("Domain-%s, %s\n", domain, status)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
