package keccak

const (
	Rate     = 1088   // Rate in bits for SHA3-256
	Capacity = 512    // Capacity in bits
	BlockSize = Rate / 8 // Block size in bytes
)

// Keccak - Keccak hash function state
type Keccak struct {
	state [25]uint64
	rate  int
}

// NewKeccak - Initialize a new Keccak instance
func NewKeccak() *Keccak {
	return &Keccak{
		rate: Rate,
	}
}

// Absorb - Absorb input data into the state
func (k *Keccak) Absorb(data []byte) {
	for len(data) > 0 {
		toAbsorb := min(len(data), k.rate/8)
		block := make([]byte, k.rate/8)
		copy(block, data[:toAbsorb])
		data = data[toAbsorb:]
		for i := 0; i < toAbsorb; i++ {
			k.state[i/8] ^= uint64(block[i]) << (8 * (i % 8))
		}
		k.state = keccakF(k.state)
	}
}

// Squeeze - Squeeze output from the state
func (k *Keccak) Squeeze(length int) []byte {
	var output []byte
	for length > 0 {
		block := make([]byte, k.rate/8)
		for i, v := range k.state[:len(block)/8] {
			block[i] = byte(v)
		}
		if length < len(block) {
			output = append(output, block[:length]...)
			length = 0
		} else {
			output = append(output, block...)
			length -= len(block)
			k.state = keccakF(k.state)
		}
	}
	return output
}



// keccakF - Apply the Keccak-f permutation
func keccakF(state [25]uint64) [25]uint64 {
	// Implement the Keccak-f permutation rounds here
	for round := 0; round < 24; round++ {
		state = roundFunction(state, round)
	}
	return state
}

// roundFunction - Perform a single round of the Keccak-f permutation
func roundFunction(state [25]uint64, round int) [25]uint64 {
	// Constants for Keccak-f
	const (
		thetaRhoPiChiIota = 0 // Placeholder for actual constants
	)

	// θ (Theta)
	C := [5]uint64{}
	D := [5]uint64{}
	for x := 0; x < 5; x++ {
		C[x] = state[x] ^ state[x+5] ^ state[x+10] ^ state[x+15] ^ state[x+20]
	}
	for x := 0; x < 5; x++ {
		D[x] = C[(x+4)%5] ^ (C[(x+1)%5]<<1 | C[(x+1)%5]>>(64-1))
	}
	for x := 0; x < 25; x += 5 {
		for y := 0; y < 5; y++ {
			state[x+y] ^= D[y]
		}
	}

	// ρ (Rho) and π (Pi)
	var (
		x = 1
		y = 0
	)
	for t := 0; t < 24; t++ {
		state[x+5*y], state[(x+5*y+1)%25] = state[(x+5*y+1)%25], state[x+5*y]
		state[(x+5*y+1)%25], state[(x+5*y+2)%25] = state[(x+5*y+2)%25], state[(x+5*y+1)%25]
		state[(x+5*y+2)%25], state[(x+5*y+3)%25] = state[(x+5*y+3)%25], state[(x+5*y+2)%25]
		state[(x+5*y+3)%25], state[(x+5*y+4)%25] = state[(x+5*y+4)%25], state[(x+5*y+3)%25]
		state[(x+5*y+4)%25] = state[(x+5*y+4)%25]
	}

	// χ (Chi)
	for x := 0; x < 25; x += 5 {
		for y := 0; y < 5; y++ {
			state[x+y] ^= ^state[x+(y+1)%5] & state[x+(y+2)%5]
		}
	}

	// ι (Iota)
	state[0] ^= roundConstants[round]
	return state
}

// pad10star1 - Padding for the SHA3 function
func pad10star1(length, rate int) []byte {
	paddingLen := ((length + 1 + rate) / rate) * rate
	pad := make([]byte, paddingLen-length)
	pad[0] = 0x01
	pad[len(pad)-1] = 0x80
	return pad
}

// min - Utility function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var roundConstants = [24]uint64{
	0x00000001, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
}






