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
type COTPPacketConnectionResponse struct {
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass
	Parent               *COTPPacket
}

// The corresponding interface
type ICOTPPacketConnectionResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketConnectionResponse) TpduCode() uint8 {
	return 0xD0
}

func (m *COTPPacketConnectionResponse) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.Parent.Parameters = parameters
	m.Parent.Payload = payload
}

func NewCOTPPacketConnectionResponse(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, parameters []*COTPParameter, payload *S7Message) *COTPPacket {
	child := &COTPPacketConnectionResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
		Parent:               NewCOTPPacket(parameters, payload),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPPacketConnectionResponse(structType interface{}) *COTPPacketConnectionResponse {
	castFunc := func(typ interface{}) *COTPPacketConnectionResponse {
		if casted, ok := typ.(COTPPacketConnectionResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacketConnectionResponse); ok {
			return casted
		}
		if casted, ok := typ.(COTPPacket); ok {
			return CastCOTPPacketConnectionResponse(casted.Child)
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return CastCOTPPacketConnectionResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacketConnectionResponse) GetTypeName() string {
	return "COTPPacketConnectionResponse"
}

func (m *COTPPacketConnectionResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPPacketConnectionResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Enum Field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPPacketConnectionResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketConnectionResponseParse(io utils.ReadBuffer) (*COTPPacket, error) {
	if pullErr := io.PullContext("COTPPacketConnectionResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (destinationReference)
	destinationReference, _destinationReferenceErr := io.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}

	// Simple Field (sourceReference)
	sourceReference, _sourceReferenceErr := io.ReadUint16("sourceReference", 16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field")
	}

	if pullErr := io.PullContext("protocolClass"); pullErr != nil {
		return nil, pullErr
	}
	// Enum field (protocolClass)
	protocolClass, _protocolClassErr := COTPProtocolClassParse(io)
	if _protocolClassErr != nil {
		return nil, errors.Wrap(_protocolClassErr, "Error parsing 'protocolClass' field")
	}
	if closeErr := io.CloseContext("protocolClass"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("COTPPacketConnectionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketConnectionResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
		Parent:               &COTPPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPPacketConnectionResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("COTPPacketConnectionResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationReference)
		destinationReference := uint16(m.DestinationReference)
		_destinationReferenceErr := io.WriteUint16("destinationReference", 16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (sourceReference)
		sourceReference := uint16(m.SourceReference)
		_sourceReferenceErr := io.WriteUint16("sourceReference", 16, (sourceReference))
		if _sourceReferenceErr != nil {
			return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
		}

		if pushErr := io.PushContext("protocolClass"); pushErr != nil {
			return pushErr
		}
		// Enum field (protocolClass)
		protocolClass := CastCOTPProtocolClass(m.ProtocolClass)
		_protocolClassErr := protocolClass.Serialize(io)
		if _protocolClassErr != nil {
			return errors.Wrap(_protocolClassErr, "Error serializing 'protocolClass' field")
		}
		if popErr := io.PopContext("protocolClass"); popErr != nil {
			return popErr
		}

		if popErr := io.PopContext("COTPPacketConnectionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPPacketConnectionResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "destinationReference":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DestinationReference = data
			case "sourceReference":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceReference = data
			case "protocolClass":
				var data COTPProtocolClass
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ProtocolClass = data
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

func (m *COTPPacketConnectionResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DestinationReference, xml.StartElement{Name: xml.Name{Local: "destinationReference"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceReference, xml.StartElement{Name: xml.Name{Local: "sourceReference"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ProtocolClass, xml.StartElement{Name: xml.Name{Local: "protocolClass"}}); err != nil {
		return err
	}
	return nil
}

func (m COTPPacketConnectionResponse) String() string {
	return string(m.Box("", 120))
}

func (m COTPPacketConnectionResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "COTPPacketConnectionResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DestinationReference", m.DestinationReference, -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("SourceReference", m.SourceReference, -1))
		// Enum field (protocolClass)
		protocolClass := CastCOTPProtocolClass(m.ProtocolClass)
		boxes = append(boxes, protocolClass.Box("protocolClass", -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
