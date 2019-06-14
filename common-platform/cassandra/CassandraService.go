package cassandra

import "github.com/gocql/gocql"

//connect to the cassandra cluster and return session
func ConnectToCluster() *gocql.Session {
	cluster := gocql.NewCluster("localhost")
	cluster.CQLVersion = "3.11.3"
	cluster.Keyspace = "article_db"
	cluster.Consistency = gocql.Quorum
	session, error := cluster.CreateSession()
	if error == nil {
		return session
	} else {
		panic("Cassandra cluster is not available")
	}
}

//
func CloseSession(session *gocql.Session) {
	session.Close()
}
