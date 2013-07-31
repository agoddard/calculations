package calculations

import ()
import "os"

const (
	mongoUrl = os.Getenv("MONGOHQ_URL")
	collection = "calculations"
)

