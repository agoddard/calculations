package calculations

import ()
import "os"

var (
	mongoUrl = os.Getenv("MONGOHQ_URL")
	collection = "calculations"
)
