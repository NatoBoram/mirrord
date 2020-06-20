# Mirrord

[![pipeline status](https://gitlab.com/NatoBoram/mirrord/badges/master/pipeline.svg)](https://gitlab.com/NatoBoram/mirrord/-/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/NatoBoram/mirrord)](https://goreportcard.com/report/gitlab.com/NatoBoram/mirrord)
[![GoDoc](https://godoc.org/gitlab.com/NatoBoram/mirrord?status.svg)](https://godoc.org/gitlab.com/NatoBoram/mirrord)
[![StackShare](https://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/NatoBoram/mirrord)

Handy tool to create and manage IPFS mirrors.

## Install

```sh
go get -u -v gitlab.com/NatoBoram/mirrord
cd ~/go/src/gitlab.com/NatoBoram/mirrord
go install
```

### Dependencies

- Go <https://golang.org/dl/>
- Btrfs (optional) <https://btrfs.wiki.kernel.org/index.php/Main_Page>

## Usage

Configuration file examples can be found in the configs folder.

Add a configuration file to `~/.config/mirrord/config.json`. This file contains paths to scripts to be ran before and after the update process. Example : [`config.json`](configs/config.json)

```json
{
 "before_script": "",
 "after_script": ""
}
```

`before_script` and `after_script` are paths to scripts to be ran before and after the update process.

Add mirrors to `~/.config/mirrord/mirrors/*.json`. Every file in this directory will be considered mirrors and will be used. Example : [`ubuntu.json`](configs/ubuntu.json).

```json
{
 "name": "ubuntu",
 "update": "/home/ubuntu/Bash/rsync_ubuntu.sh",
 "path": "/mnt/Seagate/mirrors/ubuntu",
 "snapshots": "/mnt/Seagate/snapshots/ubuntu"
}
```

`name` is simply the name of the mirror. It will be used to generate an IPNS key. `update` is the path to your update script. `path` is where the mirror is located; it will be added to IPFS. `snapshots` is optional; only set it if you want to use `--nocopy` and Btrfs snapshots.

Mirrord has two modes : One using `--nocopy` with Btrfs snapshots and one without both. It's impossible to mix them because of limitations from IPFS.

This is currently being used to create an IPFS mirror of the Ubuntu Archives. [Read more](https://www.reddit.com/r/ipfs/comments/hc9aqd/).

```list
deb http://localhost:8080/ipns/QmRzYWabKciZNiRxnyPZGbcY8XWDBkqUwNdjXpsm1q2v7F/ubuntu focal           main restricted universe multiverse # IPNS
deb http://localhost:8080/ipns/QmRzYWabKciZNiRxnyPZGbcY8XWDBkqUwNdjXpsm1q2v7F/ubuntu focal-updates   main restricted universe multiverse # IPNS
deb http://localhost:8080/ipns/QmRzYWabKciZNiRxnyPZGbcY8XWDBkqUwNdjXpsm1q2v7F/ubuntu focal-backports main restricted universe multiverse # IPNS
deb http://localhost:8080/ipns/QmRzYWabKciZNiRxnyPZGbcY8XWDBkqUwNdjXpsm1q2v7F/ubuntu focal-security  main restricted universe multiverse # IPNS
deb http://localhost:8080/ipns/QmRzYWabKciZNiRxnyPZGbcY8XWDBkqUwNdjXpsm1q2v7F/ubuntu focal-proposed  main restricted universe multiverse # IPNS
```

## License

[GNU GPLv3](LICENSE.md) 🄯 Nato Boram
