package eventlogstore

import (
	ipfslog "berty.tech/go-ipfs-log"
	"berty.tech/go-orbit-db/iface"
	"sync"
)

type eventIndex struct {
	index ipfslog.Log
	lock  sync.RWMutex
}

func (i *eventIndex) Get(key string) interface{} {
	i.lock.RLock()
	idx := i.index
	i.lock.RUnlock()

	if idx == nil {
		return nil
	}

	return idx.Values().Slice()
}

func (i *eventIndex) UpdateIndex(log ipfslog.Log, _ []ipfslog.Entry) error {
	i.lock.Lock()
	i.index = log
	i.lock.Unlock()

	return nil
}

// NewEventIndex Creates a new index for an EventLog Store
func NewEventIndex(_ []byte) iface.StoreIndex {
	return &eventIndex{}
}

var _ iface.IndexConstructor = NewEventIndex
var _ iface.StoreIndex = &eventIndex{}
