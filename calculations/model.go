package calculations

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"strings"
)

type Calculation struct {
	Calculation, OS, Language string
	Id                        bson.ObjectId "_id"
	Answer                    int
	Instance				  string
	Time                      time.Time
}

func NewCalculation(calculation string) *Calculation {
	return &Calculation{
		Calculation: calculation,
		OS:          "",
		Language:    "",
		Answer:      0,
		Instance:    "",
		Id:          bson.NewObjectId(),
		Time:        bson.Now(),
	}
}

func Get(id string, collection *mgo.Collection) (c *Calculation, err error) {
	err = collection.FindId(bson.ObjectIdHex(id)).One(&c)
	return c, err
}

func (c *Calculation) Save(collection *mgo.Collection) (err error) {
	err = collection.Insert(c)
	return err
}

func (c *Calculation) Escaped() string {
	r := strings.NewReplacer("+", "add", "*", "mul", "-", "sub", "/", "div")
	return r.Replace(c.Calculation)
}
