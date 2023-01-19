package location

import (
	"github.com/kwanok/spatial-query-study/api/db"
	"testing"
)

func TestFetchNearLocationsV1(t *testing.T) {
	db.Start()
	defer db.Conn.Close()
	FetchNearLocationsV1(&NearQuery{
		X:  1,
		Y:  2,
		Km: 10,
	})
}

func TestFetchNearLocationsV2(t *testing.T) {
	defer db.Conn.Close()
	FetchNearLocationsV2(&NearQuery{
		X:  1,
		Y:  2,
		Km: 10,
	})
}
