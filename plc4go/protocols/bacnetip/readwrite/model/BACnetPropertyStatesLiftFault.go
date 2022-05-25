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

// BACnetPropertyStatesLiftFault is the data-structure of this message
type BACnetPropertyStatesLiftFault struct {
	*BACnetPropertyStates
	LiftFault *BACnetLiftFaultTagged
}

// IBACnetPropertyStatesLiftFault is the corresponding interface of BACnetPropertyStatesLiftFault
type IBACnetPropertyStatesLiftFault interface {
	IBACnetPropertyStates
	// GetLiftFault returns LiftFault (property field)
	GetLiftFault() *BACnetLiftFaultTagged
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

func (m *BACnetPropertyStatesLiftFault) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesLiftFault) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesLiftFault) GetLiftFault() *BACnetLiftFaultTagged {
	return m.LiftFault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftFault factory function for BACnetPropertyStatesLiftFault
func NewBACnetPropertyStatesLiftFault(liftFault *BACnetLiftFaultTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesLiftFault {
	_result := &BACnetPropertyStatesLiftFault{
		LiftFault:            liftFault,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesLiftFault(structType interface{}) *BACnetPropertyStatesLiftFault {
	if casted, ok := structType.(BACnetPropertyStatesLiftFault); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftFault); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLiftFault(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLiftFault(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesLiftFault) GetTypeName() string {
	return "BACnetPropertyStatesLiftFault"
}

func (m *BACnetPropertyStatesLiftFault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesLiftFault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftFault)
	lengthInBits += m.LiftFault.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesLiftFault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftFaultParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesLiftFault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftFault"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftFault)
	if pullErr := readBuffer.PullContext("liftFault"); pullErr != nil {
		return nil, pullErr
	}
	_liftFault, _liftFaultErr := BACnetLiftFaultTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftFaultErr != nil {
		return nil, errors.Wrap(_liftFaultErr, "Error parsing 'liftFault' field")
	}
	liftFault := CastBACnetLiftFaultTagged(_liftFault)
	if closeErr := readBuffer.CloseContext("liftFault"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftFault"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesLiftFault{
		LiftFault:            CastBACnetLiftFaultTagged(liftFault),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesLiftFault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftFault"); pushErr != nil {
			return pushErr
		}

		// Simple Field (liftFault)
		if pushErr := writeBuffer.PushContext("liftFault"); pushErr != nil {
			return pushErr
		}
		_liftFaultErr := m.LiftFault.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("liftFault"); popErr != nil {
			return popErr
		}
		if _liftFaultErr != nil {
			return errors.Wrap(_liftFaultErr, "Error serializing 'liftFault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftFault"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesLiftFault) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}