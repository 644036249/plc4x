/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// BACnetLogRecordLogDatumNullValue is the data-structure of this message
type BACnetLogRecordLogDatumNullValue struct {
	*BACnetLogRecordLogDatum
	NullValue *BACnetContextTagNull

	// Arguments.
	TagNumber uint8
}

// IBACnetLogRecordLogDatumNullValue is the corresponding interface of BACnetLogRecordLogDatumNullValue
type IBACnetLogRecordLogDatumNullValue interface {
	IBACnetLogRecordLogDatum
	// GetNullValue returns NullValue (property field)
	GetNullValue() *BACnetContextTagNull
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

func (m *BACnetLogRecordLogDatumNullValue) InitializeParent(parent *BACnetLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetLogRecordLogDatum.OpeningTag = openingTag
	m.BACnetLogRecordLogDatum.PeekedTagHeader = peekedTagHeader
	m.BACnetLogRecordLogDatum.ClosingTag = closingTag
}

func (m *BACnetLogRecordLogDatumNullValue) GetParent() *BACnetLogRecordLogDatum {
	return m.BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogRecordLogDatumNullValue) GetNullValue() *BACnetContextTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumNullValue factory function for BACnetLogRecordLogDatumNullValue
func NewBACnetLogRecordLogDatumNullValue(nullValue *BACnetContextTagNull, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLogRecordLogDatumNullValue {
	_result := &BACnetLogRecordLogDatumNullValue{
		NullValue:               nullValue,
		BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogRecordLogDatumNullValue(structType interface{}) *BACnetLogRecordLogDatumNullValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumNullValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumNullValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogRecordLogDatum); ok {
		return CastBACnetLogRecordLogDatumNullValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatum); ok {
		return CastBACnetLogRecordLogDatumNullValue(casted.Child)
	}
	return nil
}

func (m *BACnetLogRecordLogDatumNullValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumNullValue"
}

func (m *BACnetLogRecordLogDatumNullValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogRecordLogDatumNullValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogRecordLogDatumNullValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumNullValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLogRecordLogDatumNullValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumNullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumNullValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(7)), BACnetDataType(BACnetDataType_NULL))
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field")
	}
	nullValue := CastBACnetContextTagNull(_nullValue)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumNullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumNullValue")
	}

	// Create a partially initialized instance
	_child := &BACnetLogRecordLogDatumNullValue{
		NullValue:               CastBACnetContextTagNull(nullValue),
		BACnetLogRecordLogDatum: &BACnetLogRecordLogDatum{},
	}
	_child.BACnetLogRecordLogDatum.Child = _child
	return _child, nil
}

func (m *BACnetLogRecordLogDatumNullValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumNullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumNullValue")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := writeBuffer.WriteSerializable(m.NullValue)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumNullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumNullValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogRecordLogDatumNullValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}