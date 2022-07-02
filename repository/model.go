package repository

import (
	"encoding/json"
	"nails/entity"
)

type Records []entity.Record

func (r *Records) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &r)
}

func (r *Records) MarshalBinary() (data []byte, err error) {
	return json.Marshal(data)
}
