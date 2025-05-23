package postgres
 
import (
	"fmt"
	"sync"
	
    "database/sql"
    _ "github.com/lib/pq"
)


type PGConn struct {
	Db 		*sql.DB
	InUse 	bool
	Index 	int
	connStr string;
}


type PGDbConnPool struct {
	conns 		[]*PGConn;
	lastIndex 	int;
	count 		int;
	mu       	sync.Mutex;
	Name		string;
	IsDefault 	bool;
}

type PGDbMod struct {
	Pools 		[]*PGDbConnPool;
}


var pgMod PGDbMod;

func ( pool * PGDbConnPool) GetConn() * PGConn {
	pool.mu.Lock();
	for i := 0; i < pool.count; i++ { 
		if pool.conns[i].InUse == false {
			pool.conns[i].InUse = true;
			pool.mu.Unlock();
			return pool.conns[i];
		}
	}
	pool.mu.Unlock();
	return nil;
}

func (conn * PGConn) ReleaseConn() {
	if conn != nil {
		conn.InUse = false;
	}	
}

type Config struct {
	Name            string `json:"Name"`
	Port            int    `json:"Port"`
	PGConnStr       string `json:"PGConnStr"`
	PGPoolSize      int    `json:"PGPoolSize"`
	MaxIdleDuration int    `json:"MaxIdeldur"`
	IsDefault       bool   `json:"IsDefault"`
}


func CreatePool(conf Config) {
	//func CreatePool(name string, connstr string, poolsize int, maxideldur int, isdefault bool) {
	var pool *PGDbConnPool = new(PGDbConnPool)
	pool.Name = conf.Name
	pool.lastIndex = 0
	pool.count = conf.PGPoolSize
	pool.IsDefault = conf.IsDefault

	for i := 0; i < pool.count; i++ {
		var conn *PGConn = new(PGConn)
		var err error

		conn.Index = i
		conn.connStr = conf.PGConnStr
		conn.InUse = false
		conn.Db, err = sql.Open("postgres", conn.connStr)

		if err != nil {
			fmt.Printf("Open connection failed for %s - %s: %v+\n", pool.Name, conn.connStr, err.Error())
		}

		pool.conns = append(pool.conns, conn)
	}

	pgMod.Pools = append(pgMod.Pools, pool)
}

func GetDbConnPool( name string) * PGDbConnPool {
	for i := 0; i < len( pgMod.Pools); i++ { 
		if pgMod.Pools[i].Name == name {
			return pgMod.Pools[i];
		}
	}
	return nil;
}


func GetDefaultDbConnPool() * PGDbConnPool {
	for i := 0; i < len( pgMod.Pools); i++ { 
		if pgMod.Pools[i].IsDefault == true {
			return pgMod.Pools[i];
		}
	}
	return nil;
}

func GetPGConnection() * PGConn {
	pool := GetDefaultDbConnPool();
	if pool != nil {
		return pool.GetConn();
	}
	return nil
}

func PutPGConnection( conn * PGConn) {
	if conn != nil {
		conn.ReleaseConn();
	}
}


