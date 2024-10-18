
package pkg

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConn() (*pgxpool.Pool, error) {

	//new conenction
	//here you can use another libraries for database
	
	fmt.Println("success connection")

	return &pgxpool.Pool{}, nil
}
