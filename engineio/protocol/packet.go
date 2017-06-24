package protocol

const (
	// Frame types.
	FrameString = 1
	FrameBinary = 2
)

// PacketType is the type of Packet.
type PacketType byte

const (
	// Packet types.
	Open PacketType = iota
	Close
	Ping
	Pong
	Message
	Upgrade
	Noop
)

func (pt PacketType) String() string {
	switch pt {
	case Open:
		return "open"
	case Close:
		return "close"
	case Ping:
		return "ping"
	case Pong:
		return "pong"
	case Message:
		return "message"
	case Upgrade:
		return "upgrade"
	case Noop:
		return "noop"
	}
	return "unknown"
}

// StringByte converts a PacketType to byte.
func (pt PacketType) StringByte() byte {
	return byte(pt) + '0'
}

// BinaryByte converts a PacketType to byte.
func (pt PacketType) BinaryByte() byte {
	return byte(pt)
}

// Byte converts a PacketType to byte.
func (pt PacketType) Byte(frameType int) byte {
	if frameType == FrameString {
		return pt.StringByte()
	} else {
		return pt.BinaryByte()
	}
}

// ByteToPacketType converts a byte to PacketType.
func ByteToPacketType(b byte, frameType int) PacketType {
	if frameType == FrameString {
		b = b - '0'
	}
	return PacketType(b)
}

// Packet
//
// Example:
// 	<packet type id>[<data>]
//
//	4hello world
//	2
type Packet struct {
	Type PacketType
	Data []byte
}

// NewPacket creates a new instance of the Packet.
func NewPacket(t PacketType, data []byte) *Packet {
	return &Packet{Type: t, Data: data}
}
