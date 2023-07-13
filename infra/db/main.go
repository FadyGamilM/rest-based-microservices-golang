package db

import (
	"time"
)

// define the connection pool of postgres db to be utilized via the customer-repo-postgres adapter
// var Db_Conn_Pool *sql.DB

// define the db timeout connection
const DbTimeOut = 3 * time.Second
