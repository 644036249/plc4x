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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ConnectionRequestInformation struct {
	Child IConnectionRequestInformationChild
}

// The corresponding interface
type IConnectionRequestInformation interface {
	ConnectionType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IConnectionRequestInformationParent interface {
	SerializeParent(io utils.WriteBuffer, child IConnectionRequestInformation, serializeChildFunction func() error) error
	GetTypeName() string
}

type IConnectionRequestInformationChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *ConnectionRequestInformation)
	GetTypeName() string
	IConnectionRequestInformation
	utils.AsciiBoxer
}

func NewConnectionRequestInformation() *ConnectionRequestInformation {
	return &ConnectionRequestInformation{}
}

func CastConnectionRequestInformation(structType interface{}) *ConnectionRequestInformation {
	castFunc := func(typ interface{}) *ConnectionRequestInformation {
		if casted, ok := typ.(ConnectionRequestInformation); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionRequestInformation); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionRequestInformation) GetTypeName() string {
	return "ConnectionRequestInformation"
}

func (m *ConnectionRequestInformation) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ConnectionRequestInformation) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ConnectionRequestInformation) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionRequestInformation) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionRequestInformationParse(io utils.ReadBuffer) (*ConnectionRequestInformation, error) {
	if pullErr := io.PullContext("ConnectionRequestInformation"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := io.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType, _connectionTypeErr := io.ReadUint8("connectionType", 8)
	if _connectionTypeErr != nil {
		return nil, errors.Wrap(_connectionTypeErr, "Error parsing 'connectionType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ConnectionRequestInformation
	var typeSwitchError error
	switch {
	case connectionType == 0x03: // ConnectionRequestInformationDeviceManagement
		_parent, typeSwitchError = ConnectionRequestInformationDeviceManagementParse(io)
	case connectionType == 0x04: // ConnectionRequestInformationTunnelConnection
		_parent, typeSwitchError = ConnectionRequestInformationTunnelConnectionParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("ConnectionRequestInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ConnectionRequestInformation) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *ConnectionRequestInformation) SerializeParent(io utils.WriteBuffer, child IConnectionRequestInformation, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("ConnectionRequestInformation"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := io.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType := uint8(child.ConnectionType())
	_connectionTypeErr := io.WriteUint8("connectionType", 8, (connectionType))

	if _connectionTypeErr != nil {
		return errors.Wrap(_connectionTypeErr, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("ConnectionRequestInformation"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ConnectionRequestInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// ConnectionRequestInformationDeviceManagement needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionRequestInformationDeviceManagement":
			if m.Child == nil {
				m.Child = &ConnectionRequestInformationDeviceManagement{
					Parent: m,
				}
			}
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of ConnectionRequestInformation")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionRequestInformationDeviceManagement":
					var dt *ConnectionRequestInformationDeviceManagement
					if m.Child != nil {
						dt = m.Child.(*ConnectionRequestInformationDeviceManagement)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionRequestInformationTunnelConnection":
					var dt *ConnectionRequestInformationTunnelConnection
					if m.Child != nil {
						dt = m.Child.(*ConnectionRequestInformationTunnelConnection)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *ConnectionRequestInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m ConnectionRequestInformation) String() string {
	return string(m.Box("", 120))
}

func (m *ConnectionRequestInformation) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *ConnectionRequestInformation) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "ConnectionRequestInformation"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Implicit Field (structureLength)
	structureLength := uint8(uint8(m.LengthInBytes()))
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("StructureLength", structureLength, -1))
	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType := uint8(m.Child.ConnectionType())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ConnectionType", connectionType, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
