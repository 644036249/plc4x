/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesReliability is the data-structure of this message
type BACnetPropertyStatesReliability struct {
	*BACnetPropertyStates
	Reliability *BACnetReliabilityTagged
}

// IBACnetPropertyStatesReliability is the corresponding interface of BACnetPropertyStatesReliability
type IBACnetPropertyStatesReliability interface {
	IBACnetPropertyStates
	// GetReliability returns Reliability (property field)
	GetReliability() *BACnetReliabilityTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesReliability) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesReliability) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesReliability) GetReliability() *BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesReliability factory function for BACnetPropertyStatesReliability
func NewBACnetPropertyStatesReliability(reliability *BACnetReliabilityTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesReliability {
	_result := &BACnetPropertyStatesReliability{
		Reliability:          reliability,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesReliability(structType interface{}) *BACnetPropertyStatesReliability {
	if casted, ok := structType.(BACnetPropertyStatesReliability); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesReliability); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesReliability(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesReliability(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesReliability) GetTypeName() string {
	return "BACnetPropertyStatesReliability"
}

func (m *BACnetPropertyStatesReliability) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesReliability) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesReliability) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesReliabilityParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesReliability, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesReliability"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reliability)
	if pullErr := readBuffer.PullContext("reliability"); pullErr != nil {
		return nil, pullErr
	}
	_reliability, _reliabilityErr := BACnetReliabilityTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reliabilityErr != nil {
		return nil, errors.Wrap(_reliabilityErr, "Error parsing 'reliability' field")
	}
	reliability := CastBACnetReliabilityTagged(_reliability)
	if closeErr := readBuffer.CloseContext("reliability"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesReliability"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesReliability{
		Reliability:          CastBACnetReliabilityTagged(reliability),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesReliability) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesReliability"); pushErr != nil {
			return pushErr
		}

		// Simple Field (reliability)
		if pushErr := writeBuffer.PushContext("reliability"); pushErr != nil {
			return pushErr
		}
		_reliabilityErr := m.Reliability.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("reliability"); popErr != nil {
			return popErr
		}
		if _reliabilityErr != nil {
			return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesReliability"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}