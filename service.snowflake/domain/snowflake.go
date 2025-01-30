package domain

import (
	"fmt"
	"sync"
	"time"
)

/* TODO:
- Add a helper function to generate the machine ID
- Need to ensure clock sync between instances to avoid possible collisions. Could use something like NTP to get more accurate times
- Might also consider adding more logging here, especially if implementing NTP
*/

// https://en.wikipedia.org/wiki/Snowflake_ID
const (
	epoch          int64 = 1735689601 // Jan 1st, 2025 00:00:01 GMT+0000
	machineIdBits  uint  = 10         // machine id bits
	sequenceBits   uint  = 12         // sequence bits
	machineIdShift uint  = sequenceBits
	timestampShift uint  = sequenceBits + machineIdBits
	maxMachineId   int64 = -1 ^ (-1 << machineIdBits)
	maxSequence    int64 = -1 ^ (-1 << sequenceBits)
)

type SnowflakeService struct {
	mu        sync.Mutex
	machineId int64
	lastTime  int64
	sequence  int64
}

func NewSnowflakeService(machineId int64) (*SnowflakeService, error) {
	if machineId < 0 || machineId > maxMachineId {
		return nil, fmt.Errorf("machine ID must be between 0 and %d", maxMachineId)
	}

	return &SnowflakeService{
		machineId: machineId,
		lastTime:  -1,
		sequence:  0,
	}, nil
}

func (s *SnowflakeService) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()
	if now == s.lastTime {
		// Increment sequence if in the same millisecond
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// Sequence overflow, wait for the next millisecond
			for now <= s.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		// Reset sequence for a new millisecond
		s.sequence = 0
	}
	s.lastTime = now

	// Generate the ID
	id := ((now - epoch) << timestampShift) |
		(s.machineId << machineIdShift) |
		s.sequence

	return id
}
