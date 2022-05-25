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

// BACnetServiceAckVTOpen is the data-structure of this message
type BACnetServiceAckVTOpen struct {
	*BACnetServiceAck
	RemoteVtSessionIdentifier *BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceAckLength uint16
}

// IBACnetServiceAckVTOpen is the corresponding interface of BACnetServiceAckVTOpen
type IBACnetServiceAckVTOpen interface {
	IBACnetServiceAck
	// GetRemoteVtSessionIdentifier returns RemoteVtSessionIdentifier (property field)
	GetRemoteVtSessionIdentifier() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetServiceAckVTOpen) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_OPEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckVTOpen) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckVTOpen) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckVTOpen) GetRemoteVtSessionIdentifier() *BACnetApplicationTagUnsignedInteger {
	return m.RemoteVtSessionIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckVTOpen factory function for BACnetServiceAckVTOpen
func NewBACnetServiceAckVTOpen(remoteVtSessionIdentifier *BACnetApplicationTagUnsignedInteger, serviceAckLength uint16) *BACnetServiceAckVTOpen {
	_result := &BACnetServiceAckVTOpen{
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
		BACnetServiceAck:          NewBACnetServiceAck(serviceAckLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckVTOpen(structType interface{}) *BACnetServiceAckVTOpen {
	if casted, ok := structType.(BACnetServiceAckVTOpen); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckVTOpen(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckVTOpen(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *BACnetServiceAckVTOpen) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckVTOpen) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (remoteVtSessionIdentifier)
	lengthInBits += m.RemoteVtSessionIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetServiceAckVTOpen) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckVTOpenParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (*BACnetServiceAckVTOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTOpen"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (remoteVtSessionIdentifier)
	if pullErr := readBuffer.PullContext("remoteVtSessionIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_remoteVtSessionIdentifier, _remoteVtSessionIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _remoteVtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_remoteVtSessionIdentifierErr, "Error parsing 'remoteVtSessionIdentifier' field")
	}
	remoteVtSessionIdentifier := CastBACnetApplicationTagUnsignedInteger(_remoteVtSessionIdentifier)
	if closeErr := readBuffer.CloseContext("remoteVtSessionIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTOpen"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckVTOpen{
		RemoteVtSessionIdentifier: CastBACnetApplicationTagUnsignedInteger(remoteVtSessionIdentifier),
		BACnetServiceAck:          &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckVTOpen) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTOpen"); pushErr != nil {
			return pushErr
		}

		// Simple Field (remoteVtSessionIdentifier)
		if pushErr := writeBuffer.PushContext("remoteVtSessionIdentifier"); pushErr != nil {
			return pushErr
		}
		_remoteVtSessionIdentifierErr := m.RemoteVtSessionIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("remoteVtSessionIdentifier"); popErr != nil {
			return popErr
		}
		if _remoteVtSessionIdentifierErr != nil {
			return errors.Wrap(_remoteVtSessionIdentifierErr, "Error serializing 'remoteVtSessionIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTOpen"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}