package service

import (
	"sync"
	"testing"
	"time"
)

func TestSnowflake_Generate(t *testing.T) {
	svc, err := NewSnowflake(0)
	if err != nil {
		t.Fatalf("failed starting snowflake service: %v", err)
	}

	t.Run("basic id generation", func(t *testing.T) {
		id := svc.Generate()
		if id <= 0 {
			t.Errorf("expected a positive id, got %v", id)
		}
	})

	t.Run("ids are unique", func(t *testing.T) {
		id1 := svc.Generate()
		id2 := svc.Generate()
		if id1 == id2 {
			t.Errorf("expected unique ids, got %v and %v", id1, id2)
		}
	})

	t.Run("sequence increment", func(t *testing.T) {
		svc.lastTime = time.Now().UnixMilli()
		svc.sequence = 0
		id1 := svc.Generate()
		id2 := svc.Generate()
		if id1 >= id2 {
			t.Errorf("expected id2 to be greater than id1, got %v and %v", id1, id2)
		}
	})

	t.Run("sequence reset", func(t *testing.T) {
		// mock new millisecond
		svc.lastTime = time.Now().UnixMilli() - 1
		// mock previous sequence overflow
		svc.sequence = maxSequence

		id := svc.Generate()

		if svc.sequence != 0 {
			t.Errorf("expected sequence to reset, got %v", svc.sequence)
		}

		if id <= 0 {
			t.Errorf("expected valid id, got %v", id)
		}
	})

	t.Run("wait for next millisecond", func(t *testing.T) {
		// mock sequence overflow
		svc.sequence = maxSequence
		svc.lastTime = time.Now().UnixMilli()

		start := time.Now()
		svc.Generate()
		elapsed := time.Since(start)

		if elapsed < time.Millisecond {
			t.Errorf("expected a wait of at least 1 millisecond, got %v", elapsed)
		}
	})
}

func TestGenerate_ConcurrentRequests(t *testing.T) {
	svc, err := NewSnowflake(0)
	if err != nil {
		t.Fatalf("failed starting snowflake service: %v", err)
	}

	var wg sync.WaitGroup
	const numRequests = 100
	ids := make(chan int64, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ids <- svc.Generate()
		}()
	}

	wg.Wait()
	close(ids)

	idSet := make(map[int64]struct{})
	for id := range ids {
		if _, exists := idSet[id]; exists {
			t.Fatalf("duplicate ID generated: %v", id)
		}
		idSet[id] = struct{}{}
	}
}
