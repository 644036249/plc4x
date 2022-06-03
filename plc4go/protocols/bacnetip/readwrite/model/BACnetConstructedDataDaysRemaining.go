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

// BACnetConstructedDataDaysRemaining is the data-structure of this message
type BACnetConstructedDataDaysRemaining struct {
	*BACnetConstructedData
	DaysRemaining *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataDaysRemaining is the corresponding interface of BACnetConstructedDataDaysRemaining
type IBACnetConstructedDataDaysRemaining interface {
	IBACnetConstructedData
	// GetDaysRemaining returns DaysRemaining (property field)
	GetDaysRemaining() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConstructedDataDaysRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDaysRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYS_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDaysRemaining) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDaysRemaining) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDaysRemaining) GetDaysRemaining() *BACnetApplicationTagSignedInteger {
	return m.DaysRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaysRemaining factory function for BACnetConstructedDataDaysRemaining
func NewBACnetConstructedDataDaysRemaining(daysRemaining *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataDaysRemaining {
	_result := &BACnetConstructedDataDaysRemaining{
		DaysRemaining:         daysRemaining,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDaysRemaining(structType interface{}) *BACnetConstructedDataDaysRemaining {
	if casted, ok := structType.(BACnetConstructedDataDaysRemaining); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaysRemaining); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDaysRemaining(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDaysRemaining(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDaysRemaining) GetTypeName() string {
	return "BACnetConstructedDataDaysRemaining"
}

func (m *BACnetConstructedDataDaysRemaining) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDaysRemaining) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (daysRemaining)
	lengthInBits += m.DaysRemaining.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDaysRemaining) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDaysRemainingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataDaysRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaysRemaining"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (daysRemaining)
	if pullErr := readBuffer.PullContext("daysRemaining"); pullErr != nil {
		return nil, pullErr
	}
	_daysRemaining, _daysRemainingErr := BACnetApplicationTagParse(readBuffer)
	if _daysRemainingErr != nil {
		return nil, errors.Wrap(_daysRemainingErr, "Error parsing 'daysRemaining' field")
	}
	daysRemaining := CastBACnetApplicationTagSignedInteger(_daysRemaining)
	if closeErr := readBuffer.CloseContext("daysRemaining"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaysRemaining"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDaysRemaining{
		DaysRemaining:         CastBACnetApplicationTagSignedInteger(daysRemaining),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDaysRemaining) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaysRemaining"); pushErr != nil {
			return pushErr
		}

		// Simple Field (daysRemaining)
		if pushErr := writeBuffer.PushContext("daysRemaining"); pushErr != nil {
			return pushErr
		}
		_daysRemainingErr := m.DaysRemaining.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("daysRemaining"); popErr != nil {
			return popErr
		}
		if _daysRemainingErr != nil {
			return errors.Wrap(_daysRemainingErr, "Error serializing 'daysRemaining' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaysRemaining"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDaysRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
