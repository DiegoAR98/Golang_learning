package vehicle

import "time"

// base struct all others embed
type Vehicle struct {
	Brand string
	Model string
	Year  int
	Color string
}
