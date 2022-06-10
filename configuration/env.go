package configuration

import "fmt"

type EnvConfiguration struct {
	Listen string
}

func (c *EnvConfiguration) Configure() error {
	fmt.Println("getting data from env file or from env and validate it")
	c.Listen = ":8080"
	return nil
}
