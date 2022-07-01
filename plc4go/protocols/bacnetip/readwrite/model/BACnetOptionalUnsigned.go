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

// BACnetOptionalUnsigned is the corresponding interface of BACnetOptionalUnsigned
type BACnetOptionalUnsigned interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetOptionalUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalUnsigned.
// This is useful for switch cases.
type BACnetOptionalUnsignedExactly interface {
	BACnetOptionalUnsigned
	isBACnetOptionalUnsigned() bool
}

// _BACnetOptionalUnsigned is the data-structure of this message
type _BACnetOptionalUnsigned struct {
	_BACnetOptionalUnsignedChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetOptionalUnsignedChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetOptionalUnsignedParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetOptionalUnsigned, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetOptionalUnsignedChild interface {
	utils.Serializable
	InitializeParent(parent BACnetOptionalUnsigned, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetOptionalUnsigned

	GetTypeName() string
	BACnetOptionalUnsigned
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsigned) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetOptionalUnsigned) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalUnsigned factory function for _BACnetOptionalUnsigned
func NewBACnetOptionalUnsigned(peekedTagHeader BACnetTagHeader) *_BACnetOptionalUnsigned {
	return &_BACnetOptionalUnsigned{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsigned(structType interface{}) BACnetOptionalUnsigned {
	if casted, ok := structType.(BACnetOptionalUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsigned) GetTypeName() string {
	return "BACnetOptionalUnsigned"
}

func (m *_BACnetOptionalUnsigned) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalUnsigned) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalUnsignedParse(readBuffer utils.ReadBuffer) (BACnetOptionalUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetOptionalUnsignedChildSerializeRequirement interface {
		BACnetOptionalUnsigned
		InitializeParent(BACnetOptionalUnsigned, BACnetTagHeader)
		GetParent() BACnetOptionalUnsigned
	}
	var _childTemp interface{}
	var _child BACnetOptionalUnsignedChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalUnsignedNull
		_childTemp, typeSwitchError = BACnetOptionalUnsignedNullParse(readBuffer)
	case true: // BACnetOptionalUnsignedValue
		_childTemp, typeSwitchError = BACnetOptionalUnsignedValueParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetOptionalUnsigned.")
	}
	_child = _childTemp.(BACnetOptionalUnsignedChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsigned")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetOptionalUnsigned) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetOptionalUnsigned, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetOptionalUnsigned"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsigned")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetOptionalUnsigned"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsigned")
	}
	return nil
}

func (m *_BACnetOptionalUnsigned) isBACnetOptionalUnsigned() bool {
	return true
}

func (m *_BACnetOptionalUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
