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
type SearchResponse struct {
	HpaiControlEndpoint *HPAIControlEndpoint
	DibDeviceInfo       *DIBDeviceInfo
	DibSuppSvcFamilies  *DIBSuppSvcFamilies
	Parent              *KnxNetIpMessage
}

// The corresponding interface
type ISearchResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SearchResponse) MsgType() uint16 {
	return 0x0202
}

func (m *SearchResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewSearchResponse(hpaiControlEndpoint *HPAIControlEndpoint, dibDeviceInfo *DIBDeviceInfo, dibSuppSvcFamilies *DIBSuppSvcFamilies) *KnxNetIpMessage {
	child := &SearchResponse{
		HpaiControlEndpoint: hpaiControlEndpoint,
		DibDeviceInfo:       dibDeviceInfo,
		DibSuppSvcFamilies:  dibSuppSvcFamilies,
		Parent:              NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSearchResponse(structType interface{}) *SearchResponse {
	castFunc := func(typ interface{}) *SearchResponse {
		if casted, ok := typ.(SearchResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SearchResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastSearchResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastSearchResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SearchResponse) GetTypeName() string {
	return "SearchResponse"
}

func (m *SearchResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SearchResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.LengthInBits()

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.LengthInBits()

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.LengthInBits()

	return lengthInBits
}

func (m *SearchResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SearchResponseParse(io utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := io.PullContext("SearchResponse"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := io.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (hpaiControlEndpoint)
	hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(io)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field")
	}
	if closeErr := io.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("dibDeviceInfo"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dibDeviceInfo)
	dibDeviceInfo, _dibDeviceInfoErr := DIBDeviceInfoParse(io)
	if _dibDeviceInfoErr != nil {
		return nil, errors.Wrap(_dibDeviceInfoErr, "Error parsing 'dibDeviceInfo' field")
	}
	if closeErr := io.CloseContext("dibDeviceInfo"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("dibSuppSvcFamilies"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dibSuppSvcFamilies)
	dibSuppSvcFamilies, _dibSuppSvcFamiliesErr := DIBSuppSvcFamiliesParse(io)
	if _dibSuppSvcFamiliesErr != nil {
		return nil, errors.Wrap(_dibSuppSvcFamiliesErr, "Error parsing 'dibSuppSvcFamilies' field")
	}
	if closeErr := io.CloseContext("dibSuppSvcFamilies"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("SearchResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SearchResponse{
		HpaiControlEndpoint: hpaiControlEndpoint,
		DibDeviceInfo:       dibDeviceInfo,
		DibSuppSvcFamilies:  dibSuppSvcFamilies,
		Parent:              &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SearchResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("SearchResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := io.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return pushErr
		}
		_hpaiControlEndpointErr := m.HpaiControlEndpoint.Serialize(io)
		if popErr := io.PopContext("hpaiControlEndpoint"); popErr != nil {
			return popErr
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		// Simple Field (dibDeviceInfo)
		if pushErr := io.PushContext("dibDeviceInfo"); pushErr != nil {
			return pushErr
		}
		_dibDeviceInfoErr := m.DibDeviceInfo.Serialize(io)
		if popErr := io.PopContext("dibDeviceInfo"); popErr != nil {
			return popErr
		}
		if _dibDeviceInfoErr != nil {
			return errors.Wrap(_dibDeviceInfoErr, "Error serializing 'dibDeviceInfo' field")
		}

		// Simple Field (dibSuppSvcFamilies)
		if pushErr := io.PushContext("dibSuppSvcFamilies"); pushErr != nil {
			return pushErr
		}
		_dibSuppSvcFamiliesErr := m.DibSuppSvcFamilies.Serialize(io)
		if popErr := io.PopContext("dibSuppSvcFamilies"); popErr != nil {
			return popErr
		}
		if _dibSuppSvcFamiliesErr != nil {
			return errors.Wrap(_dibSuppSvcFamiliesErr, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := io.PopContext("SearchResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *SearchResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "hpaiControlEndpoint":
				var data HPAIControlEndpoint
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.HpaiControlEndpoint = &data
			case "dibDeviceInfo":
				var data DIBDeviceInfo
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DibDeviceInfo = &data
			case "dibSuppSvcFamilies":
				var data DIBSuppSvcFamilies
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DibSuppSvcFamilies = &data
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

func (m *SearchResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.HpaiControlEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiControlEndpoint"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DibDeviceInfo, xml.StartElement{Name: xml.Name{Local: "dibDeviceInfo"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DibSuppSvcFamilies, xml.StartElement{Name: xml.Name{Local: "dibSuppSvcFamilies"}}); err != nil {
		return err
	}
	return nil
}

func (m SearchResponse) String() string {
	return string(m.Box("", 120))
}

func (m SearchResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "SearchResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.HpaiControlEndpoint.Box("hpaiControlEndpoint", width-2))
		// Complex field (case complex)
		boxes = append(boxes, m.DibDeviceInfo.Box("dibDeviceInfo", width-2))
		// Complex field (case complex)
		boxes = append(boxes, m.DibSuppSvcFamilies.Box("dibSuppSvcFamilies", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
