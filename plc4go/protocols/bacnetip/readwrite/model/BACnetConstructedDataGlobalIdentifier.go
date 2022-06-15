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

// BACnetConstructedDataGlobalIdentifier is the data-structure of this message
type BACnetConstructedDataGlobalIdentifier struct {
	*BACnetConstructedData
	GlobalIdentifier *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataGlobalIdentifier is the corresponding interface of BACnetConstructedDataGlobalIdentifier
type IBACnetConstructedDataGlobalIdentifier interface {
	IBACnetConstructedData
	// GetGlobalIdentifier returns GlobalIdentifier (property field)
	GetGlobalIdentifier() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataGlobalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataGlobalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GLOBAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataGlobalIdentifier) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataGlobalIdentifier) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataGlobalIdentifier) GetGlobalIdentifier() *BACnetApplicationTagUnsignedInteger {
	return m.GlobalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataGlobalIdentifier) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetGlobalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGlobalIdentifier factory function for BACnetConstructedDataGlobalIdentifier
func NewBACnetConstructedDataGlobalIdentifier(globalIdentifier *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataGlobalIdentifier {
	_result := &BACnetConstructedDataGlobalIdentifier{
		GlobalIdentifier:      globalIdentifier,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataGlobalIdentifier(structType interface{}) *BACnetConstructedDataGlobalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataGlobalIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataGlobalIdentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataGlobalIdentifier(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataGlobalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataGlobalIdentifier"
}

func (m *BACnetConstructedDataGlobalIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataGlobalIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (globalIdentifier)
	lengthInBits += m.GlobalIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataGlobalIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGlobalIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataGlobalIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGlobalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (globalIdentifier)
	if pullErr := readBuffer.PullContext("globalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for globalIdentifier")
	}
	_globalIdentifier, _globalIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _globalIdentifierErr != nil {
		return nil, errors.Wrap(_globalIdentifierErr, "Error parsing 'globalIdentifier' field")
	}
	globalIdentifier := CastBACnetApplicationTagUnsignedInteger(_globalIdentifier)
	if closeErr := readBuffer.CloseContext("globalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for globalIdentifier")
	}

	// Virtual field
	_actualValue := globalIdentifier
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGlobalIdentifier")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataGlobalIdentifier{
		GlobalIdentifier:      CastBACnetApplicationTagUnsignedInteger(globalIdentifier),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataGlobalIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGlobalIdentifier")
		}

		// Simple Field (globalIdentifier)
		if pushErr := writeBuffer.PushContext("globalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for globalIdentifier")
		}
		_globalIdentifierErr := writeBuffer.WriteSerializable(m.GlobalIdentifier)
		if popErr := writeBuffer.PopContext("globalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for globalIdentifier")
		}
		if _globalIdentifierErr != nil {
			return errors.Wrap(_globalIdentifierErr, "Error serializing 'globalIdentifier' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGlobalIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataGlobalIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}