//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequestWriteGroup struct {
	Parent *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWriteGroup interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestWriteGroup) ServiceChoice() uint8 {
	return 0x0A
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestWriteGroup() *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWriteGroup{
		Parent: NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestWriteGroup(structType interface{}) *BACnetUnconfirmedServiceRequestWriteGroup {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestWriteGroup {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestWriteGroup); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestWriteGroup); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWriteGroup(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWriteGroup(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWriteGroup"
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWriteGroupParse(io utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := io.PullContext("BACnetUnconfirmedServiceRequestWriteGroup"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := io.CloseContext("BACnetUnconfirmedServiceRequestWriteGroup"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWriteGroup{
		Parent: &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("BACnetUnconfirmedServiceRequestWriteGroup"); pushErr != nil {
			return pushErr
		}

		if popErr := io.PopContext("BACnetUnconfirmedServiceRequestWriteGroup"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetUnconfirmedServiceRequestWriteGroup) String() string {
	return string(m.Box("", 120))
}

func (m BACnetUnconfirmedServiceRequestWriteGroup) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetUnconfirmedServiceRequestWriteGroup"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
