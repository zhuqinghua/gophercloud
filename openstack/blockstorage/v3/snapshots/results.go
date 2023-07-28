package snapshots

import (
	"encoding/json"
	"time"

	"github.com/zhuqinghua/gophercloud.git/pagination"
)

type Snapshot struct {
	// Unique identifier for the snapshot.
	ID string `json:"id"`
	// Current status of the snapshot.
	Status string `json:"status"`
	// Size of the snapshot in GB.
	Size int `json:"size"`
	// The date when this snapshot was created.
	CreatedAt time.Time `json:"-"`
	// Human-readable display name for the snapshot.
	Name string `json:"name"`
}

func (r *Snapshot) UnmarshalJSON(b []byte) error {
	type tmp Snapshot
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Snapshot(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}

type SnapshotPage struct {
	pagination.SinglePageBase
}

// ExtractVolumes extracts and returns Volumes. It is used while iterating over a volumes.List call.
func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	var s []Snapshot
	err := ExtractSnapshotInto(r, &s)
	return s, err
}

func ExtractSnapshotInto(r pagination.Page, v interface{}) error {
	return r.(SnapshotPage).Result.ExtractIntoSlicePtr(v, "snapshots")
}
