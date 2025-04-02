package scylla

import (
	"fmt"
	"github.com/scylladb/gocqlx/v3"
	"snd/backend/store/models"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3/table"
)

type Adapter struct {
	config		*Config
	session		gocqlx.Session	// Use getter for get connection for private methods too
	tables 		*models.Tables
}

func NewAdapter(config *Config) (a *Adapter, err error) {
	var cluster_cfg 				= gocql.NewCluster(config.IPAddress)

	cluster_cfg.Port 				= config.Port
	cluster_cfg.Consistency 		= gocql.Quorum
	cluster_cfg.Keyspace 			= config.Keyspace
	cluster_cfg.Timeout 			= config.Timeout
	cluster_cfg.Timeout 			= config.QueryTimeout
	cluster_cfg.NumConns 			= config.NumCons
	cluster_cfg.ReconnectionPolicy 	= &gocql.ConstantReconnectionPolicy{
		MaxRetries	: config.MaxRetries,
		Interval	: config.IntervalCons,
	}

	a = &Adapter{config : config}
	a.tables = &models.Tables{
		ProductTable		: table.New(productMetadata),

	}
	if a.session, err = gocqlx.WrapSession(cluster_cfg.CreateSession()); err != nil {return}

	if err = a.syncTables(); err != nil {return}

	return
}
func(a *Adapter) GetSession() (session gocqlx.Session) {return a.session}

func(a *Adapter) syncTables() (err error) {
	if err = a.session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		uuid 				text PRIMARY KEY,
		class               text,
		name				text,
		price 			    float,
		img              	text,
		description 		text, 
		specs 				map<text, text>,
	)`, a.config.ProductTable)); err != nil {return}

	return
}

func(a *Adapter) GetTables() *models.Tables {
	if a == nil {return nil}
	return a.tables
}

