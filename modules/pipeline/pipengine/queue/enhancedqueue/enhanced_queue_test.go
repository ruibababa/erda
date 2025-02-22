// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package enhancedqueue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnhancedQueue_InProcessing(t *testing.T) {
	mq := NewEnhancedQueue(20)

	inProcessing := mq.InProcessing("k1")
	assert.False(t, inProcessing, "no item now, nothing can be processing")

	mq.Add("k1", 1, time.Time{})
	inProcessing = mq.InProcessing("k1")
	assert.False(t, inProcessing, "k1 is pending")

	poppedKey := mq.PopPending()
	assert.Equal(t, "k1", poppedKey)
	inProcessing = mq.InProcessing("k1")
	assert.True(t, inProcessing, "k1 is pushed to processing")
}

func TestEnhancedQueue_InPending(t *testing.T) {
	mq := NewEnhancedQueue(20)

	inPending := mq.InPending("k1")
	assert.False(t, inPending, "no item now, nothing pending")

	mq.Add("k1", 1, time.Time{})
	inPending = mq.InPending("k1")
	assert.True(t, inPending, "k1 is pending")

	poppedKey := mq.PopPending()
	assert.Equal(t, "k1", poppedKey)
	inPending = mq.InPending("k1")
	assert.False(t, inPending, "k1 is pushed to processing")
}

func TestEnhancedQueue_InQueue(t *testing.T) {
	mq := NewEnhancedQueue(20)

	inQueue := mq.InQueue("k1")
	assert.False(t, inQueue, "no item now, nothing in queue")

	mq.Add("k1", 1, time.Time{})
	inQueue = mq.InQueue("k1")
	assert.True(t, inQueue, "k1 in queue")
}

func TestEnhancedQueue_Add(t *testing.T) {
	mq := NewEnhancedQueue(20)

	mq.Add("k1", 1, time.Time{})
	get := mq.pending.Get("k1")
	assert.NotNil(t, get, "k1 added to pending")
	assert.Equal(t, int64(1), get.Priority())

	mq.Add("k2", 2, time.Time{})
	get = mq.pending.Get("k2")
	assert.NotNil(t, get, "k2 added to pending")
	assert.Equal(t, int64(2), get.Priority())

	mq.Add("k2", 3, time.Time{})
	get = mq.pending.Get("k2")
	assert.NotNil(t, get, "k2 added to pending again, so update")
	assert.Equal(t, int64(3), get.Priority(), "k2's priority updated to 3")
}

func TestEnhancedQueue_PopPending(t *testing.T) {
	mq := NewEnhancedQueue(20)
	mq.SetProcessingWindow(1)

	mq.Add("k1", 1, time.Time{})
	mq.Add("k2", 2, time.Time{})
	popped := mq.PopPending()
	assert.NotNil(t, popped, "k2 should be popped")
	assert.Equal(t, "k2", popped, "k2's priority is higher, so popped")
	popped = mq.PopPending()
	assert.Empty(t, popped, "window size is 1 and k2 already been popped")

	mq.SetProcessingWindow(2)
	popped = mq.PopPending()
	assert.NotNil(t, popped, "window size updated to 2, so k1 can pop")
	assert.Equal(t, popped, "k1", "k1 popped")
}

func TestEnhancedQueue_SetProcessingWindow(t *testing.T) {
	mq := NewEnhancedQueue(20)
	assert.Equal(t, int64(20), mq.processingWindow)

	mq.SetProcessingWindow(10)
	assert.Equal(t, int64(10), mq.processingWindow)
}

func TestEnhancedQueue_ProcessingWindow(t *testing.T) {
	mq := NewEnhancedQueue(10)
	assert.Equal(t, int64(10), mq.ProcessingWindow())

	mq.SetProcessingWindow(20)
	assert.Equal(t, int64(20), mq.ProcessingWindow())
}
