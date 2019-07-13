package chanstore

// central channel store
// manipulate the central store via AddChannel.

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

// Messenger stores channels of connections and possibly metadata.
type Messenger struct {
	Channel chan string
	IPAddr  string
}

func (m *Messenger) String() string {
	return fmt.Sprintf("%v", m.IPAddr)
}

// createMessenger creates new messenger.
func createMessenger() *Messenger {
	newChannel := make(chan string)
	newMessenger := new(Messenger)
	newMessenger.Channel = newChannel

	return newMessenger
}

type Store struct {
	Chans map[string]*Messenger
}

var storeInitialized bool
var channelStore *Store

// InitStore allocates memory to the global channel store.
func InitStore() {
	// Non Critical Section.
	messengerMap := make(map[string]*Messenger)
	channelStore = new(Store)
	channelStore.Chans = messengerMap

	// Critical Section: accesing storeInitialized.
	mutex.Lock()
	storeInitialized = true
	mutex.Unlock()
}

// IsStoreInitialized checks if the store is initialized.
func IsStoreInitialized() bool {
	// Critical Section: accessing storeInitialized.
	mutex.Lock()
	initialized := storeInitialized
	mutex.Unlock()

	return initialized
}

// GetStore returns the store instance.
func GetStore() *Store {
	// IsStoreInitialized locks the mutex.
	if IsStoreInitialized() {
		// Critical Section: accessing channelStore
		mutex.Lock()
		chanstore := channelStore
		mutex.Unlock()
		return chanstore
	}
	return nil
}

// GetChans returns the list of channels.
func GetChans() map[string]*Messenger {
	// Critical Section: accessing channelStore.
	mutex.Lock()
	chans := channelStore.Chans
	mutex.Unlock()
	return chans
}

// AddChannel creates a new messenger for the IPAddress.
func AddChannel(IPAddr string) *Messenger {
	newMessenger := createMessenger()
	newMessenger.IPAddr = IPAddr

	// Critical Section: accessing channelStore.
	mutex.Lock()
	channelStore.Chans[IPAddr] = newMessenger
	mutex.Unlock()

	return newMessenger
}

// RemoveChannel removes a mapping of an IPAddress to a channel
func RemoveChannel(IPAddr string) {
	channelStore.Chans[IPAddr] = nil
	delete(channelStore.Chans, IPAddr)
}

func GetChannel(IPAddr string) *Messenger {
	// Critical Section: accessing channelStore.
	mutex.Lock()
	selectedMessenger := channelStore.Chans[IPAddr]
	mutex.Unlock()

	return selectedMessenger
}
