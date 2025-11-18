package webrtc

import (
	"github.com/pion/webrtc/v3"
	"sync"
)

type Peer struct {
	pc *webrtc.PeerConnection
}

type Room struct {
	peers map[string]*Peer
	mu    sync.RWMutex
}

type SFU struct {
	rooms map[string]*Room
	mu    sync.RWMutex
}

func NewSFU() *SFU { return &SFU{rooms: make(map[string]*Room)} }

func (s *SFU) CreateRoom(id string) *Room {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rooms[id] = &Room{peers: make(map[string]*Peer)}
	return s.rooms[id]
}

func (s *SFU) AddPeer(roomID, peerID string, pc *webrtc.PeerConnection) {
	room := s.getRoom(roomID)
	room.mu.Lock()
	room.peers[peerID] = &Peer{pc: pc}
	room.mu.Unlock()
}

func (s *SFU) ForwardTrack(roomID string, track *webrtc.TrackRemote, senderID string) {
	room := s.getRoom(roomID)
	room.mu.RLock()
	defer room.mu.RUnlock()
	for id, peer := range room.peers {
		if id == senderID { continue }
		peer.pc.AddTrack(track)
	}
}
