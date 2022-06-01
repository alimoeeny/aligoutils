package aliUtils

import (
	"strconv"
	"time"
)

// used to generate sessionid s
// uses the time
// reverses it, to avoid problems like "hot hashes" and busy shards
// and over kill but a good practice
func NextId() string {
	our_zero := time.Date(1978, 7, 20, 0, 0, 0, 0, time.UTC)
	delta := time.Since(our_zero)
	newid := strconv.FormatInt(int64(delta), 10)
	// revese it so that the haskeys on dynamo (or similarly shares on a db cluster and such) get evenly distributed over shards
	// or primary keys get evenly distributed over any datastore
	return Reverse(newid)
}

// NextID is used to generate sessionid s
// uses the time
// reverses it, to avoid problems like "hot hashes" and busy shards
// an overkill but a good practice
func NextID() string {
	our_zero := time.Date(1978, 7, 20, 0, 0, 0, 0, time.UTC)
	delta := time.Since(our_zero)
	newid := strconv.FormatInt(int64(delta), 10)
	// revese it so that the haskeys on dynamo (or similarly shares on a db cluster and such) get evenly distributed over shards
	// or primary keys get evenly distributed over any datastore
	return Reverse(newid)
}
