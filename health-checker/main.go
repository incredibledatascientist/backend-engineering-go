package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func checkStatus(domain, port string) string {
	addr := net.JoinHostPort(domain, port)
	timeout := 5 * time.Second

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return fmt.Sprintf("Dest: %s - Not Reachable (%v)", addr, err)
	}

	defer conn.Close()

	return fmt.Sprintf(
		"Source: %s -> Dest: %s - Reachable",
		conn.LocalAddr().String(),
		conn.RemoteAddr().String(),
	)
}

func main() {
	// go run .\main.go --port=80 --domain=pixelstat.com
	app := &cli.App{
		Name:  "HealthChecker",
		Usage: "Health checker for checking the health of a domain",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "domain",
				Value: "google.com",
				Usage: "Domain name to check status",
			},
			&cli.StringFlag{
				Name:  "port",
				Value: "80",
				Usage: "Port to check",
			},
		},
		Action: func(c *cli.Context) error {
			domain := c.String("domain")
			port := c.String("port")

			result := checkStatus(domain, port)

			fmt.Printf("Domain: %s\n%s\n", domain, result)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
