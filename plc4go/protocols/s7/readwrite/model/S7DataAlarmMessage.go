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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7DataAlarmMessage_FUNCTIONID uint8 = 0x00
const S7DataAlarmMessage_NUMBERMESSAGEOBJ uint8 = 0x01

// S7DataAlarmMessage is the corresponding interface of S7DataAlarmMessage
type S7DataAlarmMessage interface {
	utils.LengthAware
	utils.Serializable
	// GetCpuFunctionType returns CpuFunctionType (discriminator field)
	GetCpuFunctionType() uint8
}

// S7DataAlarmMessageExactly can be used when we want exactly this type and not a type which fulfills S7DataAlarmMessage.
// This is useful for switch cases.
type S7DataAlarmMessageExactly interface {
	S7DataAlarmMessage
	isS7DataAlarmMessage() bool
}

// _S7DataAlarmMessage is the data-structure of this message
type _S7DataAlarmMessage struct {
	_S7DataAlarmMessageChildRequirements
}

type _S7DataAlarmMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetCpuFunctionType() uint8
}

type S7DataAlarmMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child S7DataAlarmMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7DataAlarmMessageChild interface {
	utils.Serializable
	InitializeParent(parent S7DataAlarmMessage)
	GetParent() *S7DataAlarmMessage

	GetTypeName() string
	S7DataAlarmMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7DataAlarmMessage) GetFunctionId() uint8 {
	return S7DataAlarmMessage_FUNCTIONID
}

func (m *_S7DataAlarmMessage) GetNumberMessageObj() uint8 {
	return S7DataAlarmMessage_NUMBERMESSAGEOBJ
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7DataAlarmMessage factory function for _S7DataAlarmMessage
func NewS7DataAlarmMessage() *_S7DataAlarmMessage {
	return &_S7DataAlarmMessage{}
}

// Deprecated: use the interface for direct cast
func CastS7DataAlarmMessage(structType interface{}) S7DataAlarmMessage {
	if casted, ok := structType.(S7DataAlarmMessage); ok {
		return casted
	}
	if casted, ok := structType.(*S7DataAlarmMessage); ok {
		return *casted
	}
	return nil
}

func (m *_S7DataAlarmMessage) GetTypeName() string {
	return "S7DataAlarmMessage"
}

func (m *_S7DataAlarmMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7DataAlarmMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7DataAlarmMessageParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8) (S7DataAlarmMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7DataAlarmMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7DataAlarmMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	if functionId != S7DataAlarmMessage_FUNCTIONID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7DataAlarmMessage_FUNCTIONID) + " but got " + fmt.Sprintf("%d", functionId))
	}

	// Const Field (numberMessageObj)
	numberMessageObj, _numberMessageObjErr := readBuffer.ReadUint8("numberMessageObj", 8)
	if _numberMessageObjErr != nil {
		return nil, errors.Wrap(_numberMessageObjErr, "Error parsing 'numberMessageObj' field")
	}
	if numberMessageObj != S7DataAlarmMessage_NUMBERMESSAGEOBJ {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7DataAlarmMessage_NUMBERMESSAGEOBJ) + " but got " + fmt.Sprintf("%d", numberMessageObj))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7DataAlarmMessageChildSerializeRequirement interface {
		S7DataAlarmMessage
		InitializeParent(S7DataAlarmMessage)
		GetParent() S7DataAlarmMessage
	}
	var _childTemp interface{}
	var _child S7DataAlarmMessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case cpuFunctionType == 0x04: // S7MessageObjectRequest
		_childTemp, typeSwitchError = S7MessageObjectRequestParse(readBuffer, cpuFunctionType)
	case cpuFunctionType == 0x08: // S7MessageObjectResponse
		_childTemp, typeSwitchError = S7MessageObjectResponseParse(readBuffer, cpuFunctionType)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [cpuFunctionType=%v]", cpuFunctionType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7DataAlarmMessage.")
	}
	_child = _childTemp.(S7DataAlarmMessageChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7DataAlarmMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7DataAlarmMessage")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_S7DataAlarmMessage) SerializeParent(writeBuffer utils.WriteBuffer, child S7DataAlarmMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7DataAlarmMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7DataAlarmMessage")
	}

	// Const Field (functionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, 0x00)
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Const Field (numberMessageObj)
	_numberMessageObjErr := writeBuffer.WriteUint8("numberMessageObj", 8, 0x01)
	if _numberMessageObjErr != nil {
		return errors.Wrap(_numberMessageObjErr, "Error serializing 'numberMessageObj' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7DataAlarmMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7DataAlarmMessage")
	}
	return nil
}

func (m *_S7DataAlarmMessage) isS7DataAlarmMessage() bool {
	return true
}

func (m *_S7DataAlarmMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
