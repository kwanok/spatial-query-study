package location

import "testing"

func TestFetchNearLocationsV1(t *testing.T) {
	defer r.Conn.Close()
	FetchNearLocationsV1(&NearQuery{
		X:  1,
		Y:  2,
		Km: 10,
	})
}

func TestFetchNearLocationsV2(t *testing.T) {
	defer r.Conn.Close()
	FetchNearLocationsV2(&NearQuery{
		X:  1,
		Y:  2,
		Km: 10,
	})
}
