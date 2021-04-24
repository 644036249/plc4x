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
type AdsInvalidResponse struct {
	Parent *AdsData
}

// The corresponding interface
type IAdsInvalidResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsInvalidResponse) CommandId() CommandId {
	return CommandId_INVALID
}

func (m *AdsInvalidResponse) Response() bool {
	return true
}

func (m *AdsInvalidResponse) InitializeParent(parent *AdsData) {
}

func NewAdsInvalidResponse() *AdsData {
	child := &AdsInvalidResponse{
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsInvalidResponse(structType interface{}) *AdsInvalidResponse {
	castFunc := func(typ interface{}) *AdsInvalidResponse {
		if casted, ok := typ.(AdsInvalidResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsInvalidResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsInvalidResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsInvalidResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsInvalidResponse) GetTypeName() string {
	return "AdsInvalidResponse"
}

func (m *AdsInvalidResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsInvalidResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *AdsInvalidResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsInvalidResponseParse(io utils.ReadBuffer) (*AdsData, error) {
	if pullErr := io.PullContext("AdsInvalidResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := io.CloseContext("AdsInvalidResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsInvalidResponse{
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsInvalidResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("AdsInvalidResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := io.PopContext("AdsInvalidResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsInvalidResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *AdsInvalidResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m AdsInvalidResponse) String() string {
	return string(m.Box("", 120))
}

func (m AdsInvalidResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsInvalidResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
