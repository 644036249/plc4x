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

// BACnetPropertyStatesNodeType is the data-structure of this message
type BACnetPropertyStatesNodeType struct {
	*BACnetPropertyStates
	NodeType *BACnetNodeTypeTagged
}

// IBACnetPropertyStatesNodeType is the corresponding interface of BACnetPropertyStatesNodeType
type IBACnetPropertyStatesNodeType interface {
	IBACnetPropertyStates
	// GetNodeType returns NodeType (property field)
	GetNodeType() *BACnetNodeTypeTagged
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

func (m *BACnetPropertyStatesNodeType) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesNodeType) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesNodeType) GetNodeType() *BACnetNodeTypeTagged {
	return m.NodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNodeType factory function for BACnetPropertyStatesNodeType
func NewBACnetPropertyStatesNodeType(nodeType *BACnetNodeTypeTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesNodeType {
	_result := &BACnetPropertyStatesNodeType{
		NodeType:             nodeType,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesNodeType(structType interface{}) *BACnetPropertyStatesNodeType {
	if casted, ok := structType.(BACnetPropertyStatesNodeType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNodeType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNodeType(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNodeType(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesNodeType) GetTypeName() string {
	return "BACnetPropertyStatesNodeType"
}

func (m *BACnetPropertyStatesNodeType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesNodeType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nodeType)
	lengthInBits += m.NodeType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesNodeType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNodeTypeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesNodeType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNodeType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeType)
	if pullErr := readBuffer.PullContext("nodeType"); pullErr != nil {
		return nil, pullErr
	}
	_nodeType, _nodeTypeErr := BACnetNodeTypeTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _nodeTypeErr != nil {
		return nil, errors.Wrap(_nodeTypeErr, "Error parsing 'nodeType' field")
	}
	nodeType := CastBACnetNodeTypeTagged(_nodeType)
	if closeErr := readBuffer.CloseContext("nodeType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNodeType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesNodeType{
		NodeType:             CastBACnetNodeTypeTagged(nodeType),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesNodeType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNodeType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (nodeType)
		if pushErr := writeBuffer.PushContext("nodeType"); pushErr != nil {
			return pushErr
		}
		_nodeTypeErr := m.NodeType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("nodeType"); popErr != nil {
			return popErr
		}
		if _nodeTypeErr != nil {
			return errors.Wrap(_nodeTypeErr, "Error serializing 'nodeType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNodeType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesNodeType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}