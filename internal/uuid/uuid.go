package uuid

import "github.com/google/uuid"

func Generate() string {

	newID := uuid.New()
	return newID.String()

}
