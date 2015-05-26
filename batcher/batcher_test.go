package batcher

import (
	"testing"
	"time"

	"github.com/influxdb/influxdb"
)

// TestBatch_Size ensures that a batcher generates a batch when the size threshold is reached.
func TestBatch_Size(t *testing.T) {
	batchSize := 5
	batcher := NewBatcher(batchSize, time.Hour)
	if batcher == nil {
		t.Fatal("failed to create batcher for size test")
	}

	in := make(chan influxdb.Point)
	out := make(chan []influxdb.Point)

	go batcher.Start(in, out)

	var p influxdb.Point
	for i := 0; i < batchSize; i++ {
		in <- p
	}
	batch := <-out
	if len(batch) != batchSize {
		t.Errorf("received batch has incorrect length exp %d, got %d", batchSize, len(batch))
	}
	checkBatcherStats(t, batcher, -1, batchSize, 1, 0)
}

// TestBatch_Size ensures that a batcher generates a batch when the timeout triggers.
func TestBatch_Timeout(t *testing.T) {
	batchSize := 5
	batcher := NewBatcher(batchSize+1, 100*time.Millisecond)
	if batcher == nil {
		t.Fatal("failed to create batcher for timeout test")
	}

	in := make(chan influxdb.Point)
	out := make(chan []influxdb.Point)

	go batcher.Start(in, out)

	var p influxdb.Point
	for i := 0; i < batchSize; i++ {
		in <- p
	}
	batch := <-out
	if len(batch) != batchSize {
		t.Errorf("received batch has incorrect length exp %d, got %d", batchSize, len(batch))
	}
	checkBatcherStats(t, batcher, -1, batchSize, 0, 1)
}

// TestBatch_MultipleBatches ensures that a batcher correctly processes multiple batches.
func TestBatch_MultipleBatches(t *testing.T) {
	batchSize := 2
	batcher := NewBatcher(batchSize, 100*time.Millisecond)
	if batcher == nil {
		t.Fatal("failed to create batcher for size test")
	}

	in := make(chan influxdb.Point)
	out := make(chan []influxdb.Point)

	go batcher.Start(in, out)

	var p influxdb.Point
	var b []influxdb.Point

	in <- p
	in <- p
	b = <-out // Batch threshold reached.
	if len(b) != batchSize {
		t.Errorf("received batch has incorrect length exp %d, got %d", batchSize, len(b))
	}

	in <- p
	b = <-out // Timeout triggered.
	if len(b) != 1 {
		t.Errorf("received batch has incorrect length exp %d, got %d", 1, len(b))
	}

	checkBatcherStats(t, batcher, -1, 3, 1, 1)
}

func checkBatcherStats(t *testing.T, b *Batcher, batchTotal, pointTotal, sizeTotal, timeoutTotal int) {
	stats := b.Stats()

	if batchTotal != -1 && stats.BatchTotal != uint64(batchTotal) {
		t.Errorf("batch total stat is incorrect: %d", stats.BatchTotal)
	}
	if pointTotal != -1 && stats.PointTotal != uint64(pointTotal) {
		t.Errorf("point total stat is incorrect: %d", stats.PointTotal)
	}
	if sizeTotal != -1 && stats.SizeTotal != uint64(sizeTotal) {
		t.Errorf("size total stat is incorrect: %d", stats.SizeTotal)
	}
	if timeoutTotal != -1 && stats.TimeoutTotal != uint64(timeoutTotal) {
		t.Errorf("timeout total stat is incorrect: %d", stats.TimeoutTotal)
	}
}
