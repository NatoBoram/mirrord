// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mirror, err := UnmarshalMirror(bytes)
//    bytes, err = mirror.Marshal()

package main

import "encoding/json"

// UnmarshalMirror unmarshals a mirror configuration.
func UnmarshalMirror(data []byte) (Mirror, error) {
	var r Mirror
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal turns a mirror configuration into a JSON file.
func (r *Mirror) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Mirror represents a single mirror that mirrord will manage.
type Mirror struct {

	// Update is a path to a script that will update the mirror's files.
	Update string `json:"update"`

	// Path is the path of the mirror.
	Path string `json:"path"`

	// IPFS is the IPFS hash of the mirror.
	IPFS string `json:"ipfs"`

	// Key is the mirror's IPNS key.
	Key string `json:"key"`

	// IPNS is the IPNS key of the mirror.
	IPNS string `json:"ipns"`
}
