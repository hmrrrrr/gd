// Code generated by /internal/packages/generator DO NOT EDIT
package ssl

import "github.com/readykit/gd"

type StreamPeerStatus int64

const (
	StreamPeerStatusDisconnected StreamPeerStatus = 0
	StreamPeerStatusHandshaking StreamPeerStatus = 1
	StreamPeerStatusConnected StreamPeerStatus = 2
	StreamPeerStatusError StreamPeerStatus = 3
	StreamPeerStatusErrorHostnameMismatch StreamPeerStatus = 4
)

type StreamPeer = gd.StreamPeerSSL
