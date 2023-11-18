package raft

import (
	"github.com/hashicorp/raft"
	"log"
	"net"
	"os"
	"time"
)

func NewRaftNode(id, raftDir, raftBind string) *raft.Raft {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(id)

	store, err := raftboltdb.NewBoltStore(raftDir + "/raft.db")
	if err != nil {
		log.Fatal(err)
	}

	snapshotStore, err := raft.NewFileSnapshotStore(raftDir, 1, os.Stderr)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", raftBind)
	if err != nil {
		log.Fatal(err)
	}

	transport, err := raft.NewTCPTransport(raftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		log.Fatal(err)
	}

	ra, err := raft.NewRaft(config, nil, store, store, snapshotStore, transport)
	if err != nil {
		log.Fatal(err)
	}

	return ra
}
