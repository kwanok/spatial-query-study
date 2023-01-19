package location

import (
	"database/sql"
	"github.com/kwanok/spatial-query-study/api/db"
	"log"
)

type Repository struct {
	Conn *sql.DB
}

var r Repository

func init() {
	r.Conn = db.Conn
}

func FetchNearLocationsV1(query *NearQuery) []Location {
	return scanLocations(getRowsByStDistanceSphere(query))
}

func FetchNearLocationsV2(query *NearQuery) []Location {
	return scanLocations(getRowsBySpatialIndex(query))
}

func scanLocations(rows *sql.Rows) []Location {
	var result []Location

	for rows.Next() {
		var l Location
		err := rows.Scan(&l.Id, &l.Name, &l.Point.X, &l.Point.Y)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, l)
	}

	return result
}

func getRowsByStDistanceSphere(query *NearQuery) *sql.Rows {
	rows, err := r.Conn.Query(`
	SELECT 
    	id as 'id', 
    	name as 'name', 
    	ST_X(coords) AS 'x', 
    	ST_Y(coords) AS 'y' 
	FROM places
	WHERE ST_Distance_Sphere(Point(?, ?), coords) < ?
	`, query.X, query.Y, query.Km)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func getRowsBySpatialIndex(query *NearQuery) *sql.Rows {
	rows, err := r.Conn.Query(`
	SELECT 
    	id as 'id', 
    	name as 'name', 
    	ST_X(coords) AS 'x', 
    	ST_Y(coords) AS 'y' 
	FROM places
	WHERE ST_Contains(ST_Buffer(Point(?, ?), ?, 'KILOMETERS'), coords);
	`, query.X, query.Y, query.Km)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
