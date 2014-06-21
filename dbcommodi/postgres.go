package dbcommodi

import (
	"database/sql"
	"strings"

	"github.com/miie/msjuvi/logger"
	//"fundare/logger"
	_ "github.com/lib/pq"
)

// ConnVars is used as variables for postgres connection
type ConnVars struct {
	dbname   string
	user     string
	password string
	host     string
	port     string
	sslmode  string
}

// New initiates ConnVars with default values
func New() *ConnVars {
	//TODO: get default values from config
	return &Vars{"dbname", "user", "password", "host", "5432", "disable"}
}

var connvars = New()

// SetVars sets the connection variables
func (cv *ConnVars) SetVars(dbname, user, password, host, port, sslmode) {
	v.dbname = dbname
	v.user = user
	v.password = password
	v.host = host
	v.port = port
	v.sslmode = sslmode
}

// GetConn gets a PostgreSQL connection using ConnVars
func GetConn() (*sql.DB, err) {

	connparameters := []string{"user", "=", connvars.username, " ", "password", "=", connvars.password, " ", "dbname", "=", connvars.dbname, " ", "host", "=", connvars.host, " ", "port", "=", connvars.port, " ", "sslmode", "=", connvars.sslmode}
	conn, err := sql.Open("postgres", strings.Join(connparameters, ""))
	if err != nil {
		logger.LogWarning("error when getting portgresql connection. err: ", err)
		return nil, err
	}
	return conn, err
}
