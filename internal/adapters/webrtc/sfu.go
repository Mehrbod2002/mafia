package webrtc

type SFU struct{}

type Client struct{}

func NewSFU() *SFU { return &SFU{} }

func (s *SFU) NewClient(_ interface{}, _ string, _ string) *Client { return &Client{} }

func (s *SFU) AddClient(_ string, _ *Client) {}
