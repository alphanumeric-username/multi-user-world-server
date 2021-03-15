package user

import (
	"github.com/google/uuid"
)

//User holds data about a client user
type User struct {
	Id   uuid.UUID
	Name string
	IP   string
	X    float64
	Y    float64
}

//NewUser creates and return a user object based provided parameters
func NewUser(id uuid.UUID, name, ip string, x, y float64) User {
	var user User
	user.Id = id
	user.Name = name
	user.IP = ip
	user.X = x
	user.Y = y

	return user
}

// //GenUserUDPPacket generates a byte array to be sent via an UDP packet to the clients
// func (u User) ByteData() ([32]byte, error) {
// 	var packet [32]byte
// 	uuidBytes, err := u.Id.MarshalBinary()
// 	if err != nil {
// 		return packet, err
// 	}

// 	for idx, b := range uuidBytes {
// 		packet[idx] = b
// 	}

// 	xBits, yBits := math.Float64bits(u.X), math.Float64bits(u.Y)

// 	for i := 16; i < 24; i++ {
// 		packet[i] = byte(xBits >> (8 * (i - 16)))
// 	}
// 	for i := 24; i < 32; i++ {
// 		packet[i] = byte(yBits >> (8 * (i - 24)))
// 	}

// 	return packet, nil
// }

// func UpdateFromByteData(data [32]byte) (uuid.UUID, float64, float64, error) {
// 	id, err := uuid.FromBytes(data[:16])
// 	if err != nil {
// 		return id, 0, 0, err
// 	}

// 	x, err := byte8ToFloat64(data[16:24])
// 	if err != nil {
// 		return id, x, 0, err
// 	}

// 	y, err := byte8ToFloat64(data[24:32])

// 	return id, x, y, err
// }

// func byte8ToFloat64(bytes []byte) (float64, error) {
// 	if len(bytes) != 8 {
// 		return 0, fmt.Errorf("expected byte slice of length 8, but provided byte slice of length '%d'", len(bytes))
// 	}

// 	var n uint64 = 0
// 	for i, b := range bytes {
// 		n += (uint64(b) << (i * 8))
// 	}
// 	return math.Float64frombits(n), nil
// }
