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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetObjectIdentifierOrUnsignedInteger is the data-structure of this message
type BACnetObjectIdentifierOrUnsignedInteger struct {
	PeekedTagHeader    *BACnetTagHeader
	ObjectIdentifier   *BACnetApplicationTagObjectIdentifier
	FallbackIdentifier *BACnetApplicationTagUnsignedInteger
	NullValue          *BACnetApplicationTagNull
}

// IBACnetObjectIdentifierOrUnsignedInteger is the corresponding interface of BACnetObjectIdentifierOrUnsignedInteger
type IBACnetObjectIdentifierOrUnsignedInteger interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetApplicationTagObjectIdentifier
	// GetFallbackIdentifier returns FallbackIdentifier (property field)
	GetFallbackIdentifier() *BACnetApplicationTagUnsignedInteger
	// GetNullValue returns NullValue (property field)
	GetNullValue() *BACnetApplicationTagNull
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetObjectIdentifier() *BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetFallbackIdentifier() *BACnetApplicationTagUnsignedInteger {
	return m.FallbackIdentifier
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetNullValue() *BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetPeekedTagNumber() uint8 {
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	fallbackIdentifier := m.FallbackIdentifier
	_ = fallbackIdentifier
	nullValue := m.NullValue
	_ = nullValue
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetObjectIdentifierOrUnsignedInteger factory function for BACnetObjectIdentifierOrUnsignedInteger
func NewBACnetObjectIdentifierOrUnsignedInteger(peekedTagHeader *BACnetTagHeader, objectIdentifier *BACnetApplicationTagObjectIdentifier, fallbackIdentifier *BACnetApplicationTagUnsignedInteger, nullValue *BACnetApplicationTagNull) *BACnetObjectIdentifierOrUnsignedInteger {
	return &BACnetObjectIdentifierOrUnsignedInteger{PeekedTagHeader: peekedTagHeader, ObjectIdentifier: objectIdentifier, FallbackIdentifier: fallbackIdentifier, NullValue: nullValue}
}

func CastBACnetObjectIdentifierOrUnsignedInteger(structType interface{}) *BACnetObjectIdentifierOrUnsignedInteger {
	if casted, ok := structType.(BACnetObjectIdentifierOrUnsignedInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetObjectIdentifierOrUnsignedInteger); ok {
		return casted
	}
	return nil
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetTypeName() string {
	return "BACnetObjectIdentifierOrUnsignedInteger"
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += (*m.ObjectIdentifier).GetLengthInBits()
	}

	// Optional Field (fallbackIdentifier)
	if m.FallbackIdentifier != nil {
		lengthInBits += (*m.FallbackIdentifier).GetLengthInBits()
	}

	// Optional Field (nullValue)
	if m.NullValue != nil {
		lengthInBits += (*m.NullValue).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetObjectIdentifierOrUnsignedIntegerParse(readBuffer utils.ReadBuffer) (*BACnetObjectIdentifierOrUnsignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectIdentifierOrUnsignedInteger"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Optional Field (objectIdentifier) (Can be skipped, if a given expression evaluates to false)
	var objectIdentifier *BACnetApplicationTagObjectIdentifier = nil
	if bool((peekedTagNumber) == (0xC)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'objectIdentifier' field")
		default:
			objectIdentifier = CastBACnetApplicationTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (fallbackIdentifier) (Can be skipped, if a given expression evaluates to false)
	var fallbackIdentifier *BACnetApplicationTagUnsignedInteger = nil
	if bool((peekedTagNumber) == (0x2)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("fallbackIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'fallbackIdentifier' field")
		default:
			fallbackIdentifier = CastBACnetApplicationTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("fallbackIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (nullValue) (Can be skipped, if a given expression evaluates to false)
	var nullValue *BACnetApplicationTagNull = nil
	if bool(bool((objectIdentifier) == (nil))) && bool(bool((fallbackIdentifier) == (nil))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'nullValue' field")
		default:
			nullValue = CastBACnetApplicationTagNull(_val)
			if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetObjectIdentifierOrUnsignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetObjectIdentifierOrUnsignedInteger(peekedTagHeader, objectIdentifier, fallbackIdentifier, nullValue), nil
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetObjectIdentifierOrUnsignedInteger"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Optional Field (objectIdentifier) (Can be skipped, if the value is null)
	var objectIdentifier *BACnetApplicationTagObjectIdentifier = nil
	if m.ObjectIdentifier != nil {
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return pushErr
		}
		objectIdentifier = m.ObjectIdentifier
		_objectIdentifierErr := objectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return popErr
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}
	}

	// Optional Field (fallbackIdentifier) (Can be skipped, if the value is null)
	var fallbackIdentifier *BACnetApplicationTagUnsignedInteger = nil
	if m.FallbackIdentifier != nil {
		if pushErr := writeBuffer.PushContext("fallbackIdentifier"); pushErr != nil {
			return pushErr
		}
		fallbackIdentifier = m.FallbackIdentifier
		_fallbackIdentifierErr := fallbackIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fallbackIdentifier"); popErr != nil {
			return popErr
		}
		if _fallbackIdentifierErr != nil {
			return errors.Wrap(_fallbackIdentifierErr, "Error serializing 'fallbackIdentifier' field")
		}
	}

	// Optional Field (nullValue) (Can be skipped, if the value is null)
	var nullValue *BACnetApplicationTagNull = nil
	if m.NullValue != nil {
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return pushErr
		}
		nullValue = m.NullValue
		_nullValueErr := nullValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return popErr
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetObjectIdentifierOrUnsignedInteger"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetObjectIdentifierOrUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}