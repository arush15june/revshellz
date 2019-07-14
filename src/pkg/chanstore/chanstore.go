package chanstore

// central channel store
// manipulate the central store via AddChannel.

import (
	"errors"
	"fmt"
	"sync"
)

var mutex = &sync.RWMutex{}

// Messenger stores channels of connections and possibly metadata.
type Messenger struct {
	RChannel chan []byte
	WChannel chan []byte
	IPAddr   string
}

func (m *Messenger) String() string {
	return fmt.Sprintf("%v", m.IPAddr)
}

// createMessenger creates new messenger.
func createMessenger() *Messenger {
	newRChannel := make(chan []byte)
	newWChannel := make(chan []byte)
	newMessenger := new(Messenger)
	newMessenger.RChannel = newRChannel
	newMessenger.WChannel = newWChannel

	return newMessenger
}

// Store for channel mappings.
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
	mutex.RLock()
	initialized := storeInitialized
	mutex.RUnlock()

	return initialized
}

// GetStore returns the store instance.
func GetStore() *Store {
	// IsStoreInitialized locks the mutex.
	if IsStoreInitialized() {
		// Critical Section: accessing channelStore
		mutex.RLock()
		chanstore := channelStore
		mutex.RUnlock()
		return chanstore
	}
	return nil
}

// GetChans returns the list of channels.
func GetChans() map[string]*Messenger {
	// Critical Section: accessing channelStore.
	mutex.RLock()
	chans := channelStore.Chans
	mutex.RUnlock()
	return chans
}

// AddChannel creates a new messenger for the IPAddress.
func AddChannel(IPAddr string) *Messenger {
	newMessenger := createMessenger()
	newMessenger.IPAddr = IPAddr

	// Critical Section: accessing channelStore.
	mutex.RLock()
	channelStore.Chans[IPAddr] = newMessenger
	mutex.RUnlock()

	return newMessenger
}

// RemoveChannel removes a mapping of an IPAddress to a channel
func RemoveChannel(IPAddr string) {
	// Critical Section: Accessing channelStore
	mutex.Lock()
	channelStore.Chans[IPAddr] = nil
	delete(channelStore.Chans, IPAddr)
	mutex.Unlock()
}

// GetChannel retrieves a specfic channel from the channelStore.
func GetChannel(IPAddr string) *Messenger {
	// Critical Section: accessing channelStore.
	mutex.RLock()
	selectedMessenger := channelStore.Chans[IPAddr]
	mutex.RUnlock()

	return selectedMessenger
}

// WriteChannel writes a byte string to a []byte channel.
func WriteChannel(ch chan []byte, data []byte) {
	select {
	case ch <- data:
	}
}

// ReadChannel reads []byte's from a []byte channel.
func ReadChannel(ch chan []byte) ([]byte, error) {
	var readMsg []byte
	var err error
	select {
	case readMsg = <-ch:
		err = nil
	default:
		err = errors.New("nothing to read.")
	}

	return readMsg, err
}
