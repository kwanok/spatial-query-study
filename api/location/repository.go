package location

import (
	"database/sql"
	"github.com/kwanok/spatial-query-study/api/db"
	"log"
)

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
	rows, err := db.Conn.Query(`
	SELECT 
    	id as 'id', 
    	name as 'name', 
    	ST_X(coordinates) AS 'x', 
    	ST_Y(coordinates) AS 'y' 
	FROM locations
	WHERE ST_Distance_Sphere(Point(?, ?), coordinates) < ?;
	`, query.X, query.Y, query.Km*1000)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func getRowsBySpatialIndex(query *NearQuery) *sql.Rows {
	rows, err := db.Conn.Query(`
	SELECT 
    	id as 'id', 
    	name as 'name', 
    	ST_X(coordinates) AS 'x', 
    	ST_Y(coordinates) AS 'y' 
	FROM locations
	WHERE ST_Contains(ST_Buffer(Point(?, ?), ?), coordinates);
	`, query.X, query.Y, query.Km*1000)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func FetchPolygonLocationsV1(query *PolygonQuery) []Location {
	return scanLocations(getRowsByPolygonV1(query))
}

func FetchPolygonLocationsV2(query *PolygonQuery) []Location {
	return scanLocations(getRowsByPolygonV2(query))
}

func getRowsByPolygonV1(query *PolygonQuery) *sql.Rows {
	rows, err := db.Conn.Query(`
	SELECT 
		id as 'id', 
		name as 'name', 
		ST_X(coordinates) AS 'x', 
		ST_Y(coordinates) AS 'y' 
	FROM locations
	WHERE ST_X(coordinates) BETWEEN ? AND ? AND ST_Y(coordinates) BETWEEN ? AND ?;
	`, query.X1, query.X2, query.Y1, query.Y2)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func getRowsByPolygonV2(query *PolygonQuery) *sql.Rows {
	rows, err := db.Conn.Query(`
	SELECT 
		id as 'id', 
		name as 'name', 
		ST_X(coordinates) AS 'x', 
		ST_Y(coordinates) AS 'y' 
	FROM locations
	WHERE ST_Contains(ST_PolyFromText(?), coordinates);
	`, query.ConvertPolygon())
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
