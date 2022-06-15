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

// BACnetConstructedDataStartTime is the data-structure of this message
type BACnetConstructedDataStartTime struct {
	*BACnetConstructedData
	StartTime *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataStartTime is the corresponding interface of BACnetConstructedDataStartTime
type IBACnetConstructedDataStartTime interface {
	IBACnetConstructedData
	// GetStartTime returns StartTime (property field)
	GetStartTime() *BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetDateTime
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

func (m *BACnetConstructedDataStartTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataStartTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_START_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataStartTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataStartTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataStartTime) GetStartTime() *BACnetDateTime {
	return m.StartTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataStartTime) GetActualValue() *BACnetDateTime {
	return CastBACnetDateTime(m.GetStartTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStartTime factory function for BACnetConstructedDataStartTime
func NewBACnetConstructedDataStartTime(startTime *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataStartTime {
	_result := &BACnetConstructedDataStartTime{
		StartTime:             startTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataStartTime(structType interface{}) *BACnetConstructedDataStartTime {
	if casted, ok := structType.(BACnetConstructedDataStartTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStartTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataStartTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataStartTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataStartTime) GetTypeName() string {
	return "BACnetConstructedDataStartTime"
}

func (m *BACnetConstructedDataStartTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataStartTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataStartTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStartTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataStartTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStartTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStartTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startTime)
	if pullErr := readBuffer.PullContext("startTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for startTime")
	}
	_startTime, _startTimeErr := BACnetDateTimeParse(readBuffer)
	if _startTimeErr != nil {
		return nil, errors.Wrap(_startTimeErr, "Error parsing 'startTime' field")
	}
	startTime := CastBACnetDateTime(_startTime)
	if closeErr := readBuffer.CloseContext("startTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for startTime")
	}

	// Virtual field
	_actualValue := startTime
	actualValue := CastBACnetDateTime(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStartTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStartTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataStartTime{
		StartTime:             CastBACnetDateTime(startTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataStartTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStartTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStartTime")
		}

		// Simple Field (startTime)
		if pushErr := writeBuffer.PushContext("startTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for startTime")
		}
		_startTimeErr := writeBuffer.WriteSerializable(m.StartTime)
		if popErr := writeBuffer.PopContext("startTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for startTime")
		}
		if _startTimeErr != nil {
			return errors.Wrap(_startTimeErr, "Error serializing 'startTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStartTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStartTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataStartTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}