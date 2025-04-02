package scylla

import (
	"errors"

	"strings"
	"time"
)

type Config struct {
	IPAddress 		string			`yaml:"ip_address"`
	Keyspace 		string			`yaml:"keyspace"`

	//Tables
	ProductTable			string		`yaml:"products_table"`

	//Settings
	Port 			int				`yaml:"port"`
	MaxRetries  	int				`yaml:"max_retries"`
	NumCons			int				`yaml:"num_cons"`
	Timeout 		time.Duration	`yaml:"timeout"`
	IntervalCons 	time.Duration  	`yaml:"interval_cons"`
	QueryTimeout 	time.Duration	`yaml:"query_timeout"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil mysql config")}

	if c.Keyspace 			= strings.TrimSpace(c.Keyspace); 	c.Keyspace 	== "" 	{return errors.New("empty Keyspace string")}
	if c.IPAddress 			= strings.TrimSpace(c.IPAddress); 	c.IPAddress == "" 	{return errors.New("empty IPAddress string")}

	if c.ProductTable		= strings.TrimSpace(c.ProductTable); 			c.ProductTable 		== "" 	{return errors.New("empty UserTable string")}

	if c.Port 				== 0 	{return errors.New("empty Port string")}
	if c.Timeout.Seconds() 	< 1 	{return errors.New("timeout for scylla must be greaten than 1s")}

	return
}
