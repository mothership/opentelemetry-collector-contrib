// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "internal/data_generator/main.go". DO NOT EDIT.
// To regenerate this file run "go run internal/data_generator/main.go".

package data

import (
	"testing"
	
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
	"github.com/stretchr/testify/assert"
)

func TestSpanEvent(t *testing.T) {
	ms := NewSpanEvent()
	assert.EqualValues(t, newSpanEvent(&otlptrace.Span_Event{}), ms)

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())

	assert.EqualValues(t, NewAttributeMap(nil), ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	ms.SetAttributes(testValAttributes)
	assert.EqualValues(t, testValAttributes, ms.Attributes())

	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())

	assert.EqualValues(t, generateTestSpanEvent(), ms)
}

func TestSpanStatus(t *testing.T) {
	ms := NewSpanStatus()
	assert.EqualValues(t, newSpanStatus(&otlptrace.Status{}), ms)

	assert.EqualValues(t, StatusCode(0), ms.Code())
	testValCode := StatusCode(1)
	ms.SetCode(testValCode)
	assert.EqualValues(t, testValCode, ms.Code())

	assert.EqualValues(t, "", ms.Message())
	testValMessage := "cancelled"
	ms.SetMessage(testValMessage)
	assert.EqualValues(t, testValMessage, ms.Message())

	assert.EqualValues(t, generateTestSpanStatus(), ms)
}

func generateTestSpanEvent() SpanEvent {
	tv := NewSpanEvent()
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetName("test_name")
	tv.SetAttributes(generateTestAttributeMap())
	tv.SetDroppedAttributesCount(uint32(17))
	return tv
}

func generateTestSpanStatus() SpanStatus {
	tv := NewSpanStatus()
	tv.SetCode(StatusCode(1))
	tv.SetMessage("cancelled")
	return tv
}