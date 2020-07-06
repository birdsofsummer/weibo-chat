// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalData(bytes)
//    bytes, err = topLevel.Marshal()

package batch_exists

import "encoding/json"

type Data []DataElement

func UnmarshalData(data []byte) (Data, error) {
	var r Data
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Data) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DataElement struct {
	ID     int64 `json:"id"`    
	Result int64 `json:"result"`
}
