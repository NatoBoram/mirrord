package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

	// Snapshots is the path where snapshots are going to be created.
	Snapshots string `json:"snapshots"`

	// Snapshot is the name of the latest snapshot in the snapshots folder.
	Snapshot int64 `json:"snapshot"`

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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (m *Mirror) snapshot() error {
	err := os.MkdirAll(m.Snapshots, osPrivateDirectory)
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()

	cmd := exec.Command("btrfs", "subvolume", "snapshot", m.Path, m.Snapshots+string(os.PathSeparator)+strconv.FormatInt(timestamp, 10))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	m.Snapshot = timestamp
	return m.save()
}

func (m *Mirror) ipfs() error {
	cmd := exec.Command(
		"ipfs", "add", "--recursive", "--hidden", "--quieter", "--progress", "--chunker=rabin", "--nocopy", "--cid-version=1",
		m.Snapshots+string(os.PathSeparator)+strconv.FormatInt(m.Snapshot, 10),
	)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	cmd = exec.Command("ipfs", "pin", "rm", m.IPFS)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()

	m.IPFS = strings.TrimSpace(string(out))
	return m.save()
}

func (m *Mirror) key() error {
	if m.Key != "" {
		return nil
	}

	cmd := exec.Command("ipfs", "key", "gen", m.Name)
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

	fmt.Println("Creating a snapshot...")
	err = m.snapshot()
	if err != nil {
		return
	}

	fmt.Println("Adding to IPFS...")
	err = m.ipfs()
	if err != nil {
		return
	}

	fmt.Println("Checking the IPNS key...")
	err = m.key()
	if err != nil {
		return
	}

	fmt.Println("Publishing on IPNS...")
	err = m.ipns()
	return
}
