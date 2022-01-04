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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetTag struct {
	TagNumber                uint8
	LengthValueType          uint8
	ExtTagNumber             *uint8
	ExtLength                *uint8
	ExtExtLength             *uint16
	ExtExtExtLength          *uint32
	ActualTagNumber          uint8
	IsPrimitiveAndNotBoolean bool
	ActualLength             uint32
	Child                    IBACnetTagChild
}

// The corresponding interface
type IBACnetTag interface {
	TagClass() TagClass
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32)
	GetTypeName() string
	IBACnetTag
}

func NewBACnetTag(tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTag {
	return &BACnetTag{TagNumber: tagNumber, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength}
}

func CastBACnetTag(structType interface{}) *BACnetTag {
	castFunc := func(typ interface{}) *BACnetTag {
		if casted, ok := typ.(BACnetTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTag) GetTypeName() string {
	return "BACnetTag"
}

func (m *BACnetTag) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTag) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetTag) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (tagNumber)
	lengthInBits += 4
	// Discriminator Field (tagClass)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTag) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagParse(readBuffer utils.ReadBuffer) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetTag"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (tagNumber)
	_tagNumber, _tagNumberErr := readBuffer.ReadUint8("tagNumber", 4)
	if _tagNumberErr != nil {
		return nil, errors.Wrap(_tagNumberErr, "Error parsing 'tagNumber' field")
	}
	tagNumber := _tagNumber

	// Discriminator Field (tagClass) (Used as input to a switch field)
	tagClass_temp, _tagClassErr := TagClassParse(readBuffer)
	var tagClass TagClass = tagClass_temp
	if _tagClassErr != nil {
		return nil, errors.Wrap(_tagClassErr, "Error parsing 'tagClass' field")
	}

	// Simple Field (lengthValueType)
	_lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}
	lengthValueType := _lengthValueType

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Virtual field
	_actualTagNumber := utils.InlineIf(bool((tagNumber) < (15)), func() interface{} { return uint8(tagNumber) }, func() interface{} { return uint8((*extTagNumber)) }).(uint8)
	actualTagNumber := uint8(_actualTagNumber)

	// Virtual field
	_isPrimitiveAndNotBoolean := bool(!(bool(bool(bool((tagClass) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((lengthValueType) == (6)))))) && bool(bool((tagNumber) != (1)))
	isPrimitiveAndNotBoolean := bool(_isPrimitiveAndNotBoolean)

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5))) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Optional Field (extExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtLength *uint16 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (254))) {
		_val, _err := readBuffer.ReadUint16("extExtLength", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtLength' field")
		}
		extExtLength = &_val
	}

	// Optional Field (extExtExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtExtLength *uint32 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (255))) {
		_val, _err := readBuffer.ReadUint32("extExtExtLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtExtLength' field")
		}
		extExtExtLength = &_val
	}

	// Virtual field
	_actualLength := utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() interface{} { return uint32((*extExtExtLength)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() interface{} { return uint32((*extExtLength)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(bool((lengthValueType) == (5)), func() interface{} { return uint32((*extLength)) }, func() interface{} {
				return uint32(uint32(utils.InlineIf(isPrimitiveAndNotBoolean, func() interface{} { return uint32(lengthValueType) }, func() interface{} { return uint32(uint32(0)) }).(uint32)))
			}).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualLength := uint32(_actualLength)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetTag
	var typeSwitchError error
	switch {
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x0: // BACnetApplicationTagNull
		_parent, typeSwitchError = BACnetApplicationTagNullParse(readBuffer)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x1: // BACnetApplicationTagBoolean
		_parent, typeSwitchError = BACnetApplicationTagBooleanParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x2: // BACnetApplicationTagUnsignedInteger
		_parent, typeSwitchError = BACnetApplicationTagUnsignedIntegerParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x3: // BACnetApplicationTagSignedInteger
		_parent, typeSwitchError = BACnetApplicationTagSignedIntegerParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x4: // BACnetApplicationTagReal
		_parent, typeSwitchError = BACnetApplicationTagRealParse(readBuffer)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x5: // BACnetApplicationTagDouble
		_parent, typeSwitchError = BACnetApplicationTagDoubleParse(readBuffer)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x6: // BACnetApplicationTagOctetString
		_parent, typeSwitchError = BACnetApplicationTagOctetStringParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x7: // BACnetApplicationTagCharacterString
		_parent, typeSwitchError = BACnetApplicationTagCharacterStringParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x8: // BACnetApplicationTagBitString
		_parent, typeSwitchError = BACnetApplicationTagBitStringParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0x9: // BACnetApplicationTagEnumerated
		_parent, typeSwitchError = BACnetApplicationTagEnumeratedParse(readBuffer, actualLength)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0xA: // BACnetApplicationTagDate
		_parent, typeSwitchError = BACnetApplicationTagDateParse(readBuffer)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0xB: // BACnetApplicationTagTime
		_parent, typeSwitchError = BACnetApplicationTagTimeParse(readBuffer)
	case tagClass == TagClass_APPLICATION_TAGS && tagNumber == 0xC: // BACnetApplicationTagObjectIdentifier
		_parent, typeSwitchError = BACnetApplicationTagObjectIdentifierParse(readBuffer)
	case tagClass == TagClass_CONTEXT_SPECIFIC_TAGS: // BACnetContextTagWithoutContext
		_parent, typeSwitchError = BACnetContextTagWithoutContextParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength, actualTagNumber, isPrimitiveAndNotBoolean, actualLength)
	return _parent, nil
}

func (m *BACnetTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (tagNumber)
	tagNumber := uint8(m.TagNumber)
	_tagNumberErr := writeBuffer.WriteUint8("tagNumber", 4, (tagNumber))
	if _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}

	// Discriminator Field (tagClass) (Used as input to a switch field)
	tagClass := TagClass(child.TagClass())
	_tagClassErr := tagClass.Serialize(writeBuffer)

	if _tagClassErr != nil {
		return errors.Wrap(_tagClassErr, "Error serializing 'tagClass' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}
	// Virtual field
	if _actualTagNumberErr := writeBuffer.WriteVirtual("actualTagNumber", m.ActualTagNumber); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	if _isPrimitiveAndNotBooleanErr := writeBuffer.WriteVirtual("isPrimitiveAndNotBoolean", m.IsPrimitiveAndNotBoolean); _isPrimitiveAndNotBooleanErr != nil {
		return errors.Wrap(_isPrimitiveAndNotBooleanErr, "Error serializing 'isPrimitiveAndNotBoolean' field")
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Optional Field (extExtLength) (Can be skipped, if the value is null)
	var extExtLength *uint16 = nil
	if m.ExtExtLength != nil {
		extExtLength = m.ExtExtLength
		_extExtLengthErr := writeBuffer.WriteUint16("extExtLength", 16, *(extExtLength))
		if _extExtLengthErr != nil {
			return errors.Wrap(_extExtLengthErr, "Error serializing 'extExtLength' field")
		}
	}

	// Optional Field (extExtExtLength) (Can be skipped, if the value is null)
	var extExtExtLength *uint32 = nil
	if m.ExtExtExtLength != nil {
		extExtExtLength = m.ExtExtExtLength
		_extExtExtLengthErr := writeBuffer.WriteUint32("extExtExtLength", 32, *(extExtExtLength))
		if _extExtExtLengthErr != nil {
			return errors.Wrap(_extExtExtLengthErr, "Error serializing 'extExtExtLength' field")
		}
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.ActualLength); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
