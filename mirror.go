// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mirror, err := UnmarshalMirror(bytes)
//    bytes, err = mirror.Marshal()

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// UnmarshalMirror unmarshals a mirror configuration.
func UnmarshalMirror(data []byte) (Mirror, error) {
	var m Mirror
	err := json.Unmarshal(data, &m)
	return m, err
}

// Marshal turns a mirror configuration into a JSON file.
func (m *Mirror) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Mirror represents a single mirror that mirrord will manage.
type Mirror struct {

	// Name is the name of this mirror.
	Name string `json:"name"`

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

	path string
	info os.FileInfo
}

func (m *Mirror) update() error {
	cmd := exec.Command(m.Update)

	// Stream outputs to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}

func (m *Mirror) ipfs() error {
	cmd := exec.Command("ipfs", "add", "--recursive", "--hidden", "--quieter", "--wrap-with-directory", "--chunker=rabin", "--fscache", "--cid-version=1", m.Path)

	// Stream errors to the console
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	m.IPFS = strings.TrimSpace(string(out))
	return m.save()
}

func (m *Mirror) key() error {
	if m.Key != "" {
		return nil
	}

	cmd := exec.Command("ipfs", "key", "gen", m.Name)

	// Stream errors to the console
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	m.Key = strings.TrimSpace(string(out))
	return m.save()
}

func (m *Mirror) ipns() error {
	cmd := exec.Command("ipfs", "name", "publish", "--key="+m.Key, "--quieter", m.IPFS)

	// Stream errors to the console
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	m.IPNS = strings.TrimSpace(string(out))
	return m.save()
}

func (m *Mirror) save() error {
	bytes, err := m.Marshal()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(m.path, bytes, osPrivateFile)
}

func (m *Mirror) cycle() (err error) {
	fmt.Println("Updating...")
	err = m.update()
	if err != nil {
		return
	}

	fmt.Println("Adding to IPFS...")
	err = m.ipfs()
	if err != nil {
		return
	}

	err = m.key()
	if err != nil {
		return
	}

	fmt.Println("Publishing on IPNS...")
	err = m.ipns()
	return
}
