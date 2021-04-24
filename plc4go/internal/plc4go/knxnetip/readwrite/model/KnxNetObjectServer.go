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
type KnxNetObjectServer struct {
	Version uint8
	Parent  *ServiceId
}

// The corresponding interface
type IKnxNetObjectServer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxNetObjectServer) ServiceType() uint8 {
	return 0x08
}

func (m *KnxNetObjectServer) InitializeParent(parent *ServiceId) {
}

func NewKnxNetObjectServer(version uint8) *ServiceId {
	child := &KnxNetObjectServer{
		Version: version,
		Parent:  NewServiceId(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastKnxNetObjectServer(structType interface{}) *KnxNetObjectServer {
	castFunc := func(typ interface{}) *KnxNetObjectServer {
		if casted, ok := typ.(KnxNetObjectServer); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetObjectServer); ok {
			return casted
		}
		if casted, ok := typ.(ServiceId); ok {
			return CastKnxNetObjectServer(casted.Child)
		}
		if casted, ok := typ.(*ServiceId); ok {
			return CastKnxNetObjectServer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetObjectServer) GetTypeName() string {
	return "KnxNetObjectServer"
}

func (m *KnxNetObjectServer) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetObjectServer) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetObjectServer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetObjectServerParse(io utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := io.PullContext("KnxNetObjectServer"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (version)
	version, _versionErr := io.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	if closeErr := io.CloseContext("KnxNetObjectServer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetObjectServer{
		Version: version,
		Parent:  &ServiceId{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *KnxNetObjectServer) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("KnxNetObjectServer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := io.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := io.PopContext("KnxNetObjectServer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *KnxNetObjectServer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "version":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Version = data
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

func (m *KnxNetObjectServer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Version, xml.StartElement{Name: xml.Name{Local: "version"}}); err != nil {
		return err
	}
	return nil
}

func (m KnxNetObjectServer) String() string {
	return string(m.Box("", 120))
}

func (m KnxNetObjectServer) Box(name string, width int) utils.AsciiBox {
	boxName := "KnxNetObjectServer"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Version", m.Version, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
