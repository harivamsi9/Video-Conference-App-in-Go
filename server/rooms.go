package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in the HashMap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap is the main HashMap [roomID string] -> [[]Participant ]
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init initialises the RoomMap Struct
func (r *RoomMap) Init() {
	// r.Mutex.Lock()
	// defer r.Mutex.Unlock()

	r.Map = make(map[string][]Participant)
}

// Get will return the array of participants in the room
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]

}

// CreateRoom generate a unique ID -> insert it to the Hashmap
func (r *RoomMap) CreateRoom() string {
	// generate a unique ID -> insert it to the Hashmap
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID

}

// InsertIntoRoom will create a participant and add it in the RoomMap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{Host: host, Conn: conn}
	log.Println("INSERTING INTO ROOM WITH roomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)

}

// DeleteRoom will delete the roomID
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)

}
