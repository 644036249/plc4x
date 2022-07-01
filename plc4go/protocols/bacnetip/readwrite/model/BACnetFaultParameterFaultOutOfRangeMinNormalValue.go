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

// BACnetFaultParameterFaultOutOfRangeMinNormalValue is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValue
type BACnetFaultParameterFaultOutOfRangeMinNormalValue interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValue.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	isBACnetFaultParameterFaultOutOfRangeMinNormalValue() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValue is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValue struct {
	_BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetFaultParameterFaultOutOfRangeMinNormalValueParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMinNormalValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetFaultParameterFaultOutOfRangeMinNormalValueChild interface {
	utils.Serializable
	InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetFaultParameterFaultOutOfRangeMinNormalValue

	GetTypeName() string
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValue factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValue
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValue(structType interface{}) BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValue"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{"only application tags allowed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement interface {
		BACnetFaultParameterFaultOutOfRangeMinNormalValue
		InitializeParent(BACnetFaultParameterFaultOutOfRangeMinNormalValue, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValue
	}
	var _childTemp interface{}
	var _child BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x4: // BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueRealParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x2: // BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x5: // BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x3: // BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParse(readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue.")
	}
	_child = _childTemp.(BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMinNormalValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) isBACnetFaultParameterFaultOutOfRangeMinNormalValue() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
