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
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReadCoilsRequest struct {
	StartingAddress uint16
	Quantity        uint16
	Parent          *ModbusPDU
}

// The corresponding interface
type IModbusPDUReadCoilsRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadCoilsRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReadCoilsRequest) FunctionFlag() uint8 {
	return 0x01
}

func (m *ModbusPDUReadCoilsRequest) Response() bool {
	return false
}

func (m *ModbusPDUReadCoilsRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadCoilsRequest(startingAddress uint16, quantity uint16) *ModbusPDU {
	child := &ModbusPDUReadCoilsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReadCoilsRequest(structType interface{}) *ModbusPDUReadCoilsRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadCoilsRequest {
		if casted, ok := typ.(ModbusPDUReadCoilsRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadCoilsRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadCoilsRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadCoilsRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadCoilsRequest) GetTypeName() string {
	return "ModbusPDUReadCoilsRequest"
}

func (m *ModbusPDUReadCoilsRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadCoilsRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadCoilsRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadCoilsRequestParse(io utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := io.PullContext("ModbusPDUReadCoilsRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (startingAddress)
	startingAddress, _startingAddressErr := io.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}

	// Simple Field (quantity)
	quantity, _quantityErr := io.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}

	if closeErr := io.CloseContext("ModbusPDUReadCoilsRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadCoilsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReadCoilsRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("ModbusPDUReadCoilsRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := io.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := io.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := io.PopContext("ModbusPDUReadCoilsRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReadCoilsRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "startingAddress":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.StartingAddress = data
			case "quantity":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Quantity = data
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

func (m *ModbusPDUReadCoilsRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.StartingAddress, xml.StartElement{Name: xml.Name{Local: "startingAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUReadCoilsRequest) String() string {
	return string(m.Box("", 120))
}

func (m ModbusPDUReadCoilsRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "ModbusPDUReadCoilsRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("StartingAddress", m.StartingAddress, -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Quantity", m.Quantity, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
