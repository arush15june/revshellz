package chanstore

// central channel store
// manipulate the central store via AddChannel.

import (
	"errors"
	"fmt"
	"sync"
)

// Messenger stores channels of connections and possibly metadata.
type Messenger struct {
	RChLock  *sync.Mutex
	RChannel chan []byte
	WChLock  *sync.Mutex
	WChannel chan []byte
	IPAddr   string
}

func (m *Messenger) String() string {
	return fmt.Sprintf("%v", m.IPAddr)
}

// WriteChannel writes data to the write channel.
func (m *Messenger) WriteChannel(data []byte) {
	m.WChLock.Lock()
	select {
	case m.WChannel <- data:
	}
	m.WChLock.Unlock()
}

// ReadChannel reads data, if available, from the messengers Read channel.
func (m *Messenger) ReadChannel() ([]byte, error) {
	var readMsg []byte
	var err error
	m.RChLock.Lock()
	select {
	case readMsg = <-m.RChannel:
		err = nil
	default:
		err = errors.New("nothing to read.")
	}
	m.RChLock.Unlock()

	return readMsg, err
}

// ChannelBroker fans out data from a single channel to multiple allocated channel.
type MessengerBroker struct {
	RChannels map[*Messenger][]chan []byte
	WChannels map[*Messenger][]chan []byte
}

// createMessenger creates new messenger.
func createMessenger() *Messenger {
	newMessenger := new(Messenger)
	newMessenger.RChannel = make(chan []byte)
	newMessenger.WChannel = make(chan []byte)
	newMessenger.RChLock = new(sync.Mutex)
	newMessenger.WChLock = new(sync.Mutex)

	return newMessenger
}

// Store for channel mappings.
type Store struct {
	Chans map[string]*Messenger
	mutex *sync.RWMutex
}

var storeInitialized bool
var channelStore *Store

// InitStore allocates memory to the global channel store.
func InitStore() {
	// Non Critical Section.
	channelStore = new(Store)
	channelStore.Chans = make(map[string]*Messenger)
	channelStore.mutex = new(sync.RWMutex)

	// Critical Section: accesing storeInitialized.
	channelStore.mutex.Lock()
	storeInitialized = true
	channelStore.mutex.Unlock()
}

// IsStoreInitialized checks if the store is initialized.
func IsStoreInitialized() bool {
	// Critical Section: accessing storeInitialized.
	channelStore.mutex.RLock()
	initialized := storeInitialized
	channelStore.mutex.RUnlock()

	return initialized
}

// GetStore returns the store instance.
func GetStore() *Store {
	// IsStoreInitialized requires unlocked mutex.
	if IsStoreInitialized() {
		// Critical Section: accessing channelStore
		channelStore.mutex.RLock()
		chanstore := channelStore
		channelStore.mutex.RUnlock()
		return chanstore
	}
	return nil
}

// GetChans returns the list of channels.
func GetChans() map[string]*Messenger {
	// Critical Section: accessing channelStore.
	channelStore.mutex.RLock()
	chans := channelStore.Chans
	channelStore.mutex.RUnlock()
	return chans
}

// AddChannel creates a new messenger for the IPAddress.
func AddChannel(IPAddr string) *Messenger {
	newMessenger := createMessenger()
	newMessenger.IPAddr = IPAddr

	// Critical Section: accessing channelStore.
	channelStore.mutex.RLock()
	channelStore.Chans[IPAddr] = newMessenger
	channelStore.mutex.RUnlock()

	return newMessenger
}

// RemoveChannel removes a mapping of an IPAddress to a channel
func RemoveChannel(IPAddr string) {
	// Critical Section: Accessing channelStore
	channelStore.mutex.Lock()
	channelStore.Chans[IPAddr] = nil
	delete(channelStore.Chans, IPAddr)
	channelStore.mutex.Unlock()
}

// GetChannel retrieves a specfic channel from the channelStore.
func GetChannel(IPAddr string) *Messenger {
	// Critical Section: accessing channelStore.
	channelStore.mutex.RLock()
	selectedMessenger := channelStore.Chans[IPAddr]
	channelStore.mutex.RUnlock()

	return selectedMessenger
}
