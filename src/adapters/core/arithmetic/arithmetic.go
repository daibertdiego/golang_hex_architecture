package arithmetic

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Adapter struct {
	ID        primitive.ObjectID `bson:"_id" json:"id, omitempty"`
	Answer    int32              `bson:"answer" json:"answer"`
	Operation string             `bson:"operation" json:"operation"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith *Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith *Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith *Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith *Adapter) Division(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("Division by 0")
	}
	return a / b, nil
}
