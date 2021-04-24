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
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetErrorReadProperty_ERRORCLASSHEADER uint8 = 0x12
const BACnetErrorReadProperty_ERRORCODEHEADER uint8 = 0x12

// The data-structure of this message
type BACnetErrorReadProperty struct {
	ErrorClassLength uint8
	ErrorClass       []int8
	ErrorCodeLength  uint8
	ErrorCode        []int8
	Parent           *BACnetError
}

// The corresponding interface
type IBACnetErrorReadProperty interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorReadProperty) ServiceChoice() uint8 {
	return 0x0C
}

func (m *BACnetErrorReadProperty) InitializeParent(parent *BACnetError) {
}

func NewBACnetErrorReadProperty(errorClassLength uint8, errorClass []int8, errorCodeLength uint8, errorCode []int8) *BACnetError {
	child := &BACnetErrorReadProperty{
		ErrorClassLength: errorClassLength,
		ErrorClass:       errorClass,
		ErrorCodeLength:  errorCodeLength,
		ErrorCode:        errorCode,
		Parent:           NewBACnetError(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetErrorReadProperty(structType interface{}) *BACnetErrorReadProperty {
	castFunc := func(typ interface{}) *BACnetErrorReadProperty {
		if casted, ok := typ.(BACnetErrorReadProperty); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorReadProperty); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorReadProperty(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorReadProperty(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorReadProperty) GetTypeName() string {
	return "BACnetErrorReadProperty"
}

func (m *BACnetErrorReadProperty) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorReadProperty) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (errorClassHeader)
	lengthInBits += 5

	// Simple field (errorClassLength)
	lengthInBits += 3

	// Array field
	if len(m.ErrorClass) > 0 {
		lengthInBits += 8 * uint16(len(m.ErrorClass))
	}

	// Const Field (errorCodeHeader)
	lengthInBits += 5

	// Simple field (errorCodeLength)
	lengthInBits += 3

	// Array field
	if len(m.ErrorCode) > 0 {
		lengthInBits += 8 * uint16(len(m.ErrorCode))
	}

	return lengthInBits
}

func (m *BACnetErrorReadProperty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorReadPropertyParse(io utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := io.PullContext("BACnetErrorReadProperty"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (errorClassHeader)
	errorClassHeader, _errorClassHeaderErr := io.ReadUint8("errorClassHeader", 5)
	if _errorClassHeaderErr != nil {
		return nil, errors.Wrap(_errorClassHeaderErr, "Error parsing 'errorClassHeader' field")
	}
	if errorClassHeader != BACnetErrorReadProperty_ERRORCLASSHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetErrorReadProperty_ERRORCLASSHEADER) + " but got " + fmt.Sprintf("%d", errorClassHeader))
	}

	// Simple Field (errorClassLength)
	errorClassLength, _errorClassLengthErr := io.ReadUint8("errorClassLength", 3)
	if _errorClassLengthErr != nil {
		return nil, errors.Wrap(_errorClassLengthErr, "Error parsing 'errorClassLength' field")
	}

	// Array field (errorClass)
	if pullErr := io.PullContext("errorClass", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	errorClass := make([]int8, errorClassLength)
	for curItem := uint16(0); curItem < uint16(errorClassLength); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'errorClass' field")
		}
		errorClass[curItem] = _item
	}
	if closeErr := io.CloseContext("errorClass", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (errorCodeHeader)
	errorCodeHeader, _errorCodeHeaderErr := io.ReadUint8("errorCodeHeader", 5)
	if _errorCodeHeaderErr != nil {
		return nil, errors.Wrap(_errorCodeHeaderErr, "Error parsing 'errorCodeHeader' field")
	}
	if errorCodeHeader != BACnetErrorReadProperty_ERRORCODEHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetErrorReadProperty_ERRORCODEHEADER) + " but got " + fmt.Sprintf("%d", errorCodeHeader))
	}

	// Simple Field (errorCodeLength)
	errorCodeLength, _errorCodeLengthErr := io.ReadUint8("errorCodeLength", 3)
	if _errorCodeLengthErr != nil {
		return nil, errors.Wrap(_errorCodeLengthErr, "Error parsing 'errorCodeLength' field")
	}

	// Array field (errorCode)
	if pullErr := io.PullContext("errorCode", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	errorCode := make([]int8, errorCodeLength)
	for curItem := uint16(0); curItem < uint16(errorCodeLength); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'errorCode' field")
		}
		errorCode[curItem] = _item
	}
	if closeErr := io.CloseContext("errorCode", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("BACnetErrorReadProperty"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorReadProperty{
		ErrorClassLength: errorClassLength,
		ErrorClass:       errorClass,
		ErrorCodeLength:  errorCodeLength,
		ErrorCode:        errorCode,
		Parent:           &BACnetError{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetErrorReadProperty) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("BACnetErrorReadProperty"); pushErr != nil {
			return pushErr
		}

		// Const Field (errorClassHeader)
		_errorClassHeaderErr := io.WriteUint8("errorClassHeader", 5, 0x12)
		if _errorClassHeaderErr != nil {
			return errors.Wrap(_errorClassHeaderErr, "Error serializing 'errorClassHeader' field")
		}

		// Simple Field (errorClassLength)
		errorClassLength := uint8(m.ErrorClassLength)
		_errorClassLengthErr := io.WriteUint8("errorClassLength", 3, (errorClassLength))
		if _errorClassLengthErr != nil {
			return errors.Wrap(_errorClassLengthErr, "Error serializing 'errorClassLength' field")
		}

		// Array Field (errorClass)
		if m.ErrorClass != nil {
			if pushErr := io.PushContext("errorClass", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ErrorClass {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'errorClass' field")
				}
			}
			if popErr := io.PopContext("errorClass", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Const Field (errorCodeHeader)
		_errorCodeHeaderErr := io.WriteUint8("errorCodeHeader", 5, 0x12)
		if _errorCodeHeaderErr != nil {
			return errors.Wrap(_errorCodeHeaderErr, "Error serializing 'errorCodeHeader' field")
		}

		// Simple Field (errorCodeLength)
		errorCodeLength := uint8(m.ErrorCodeLength)
		_errorCodeLengthErr := io.WriteUint8("errorCodeLength", 3, (errorCodeLength))
		if _errorCodeLengthErr != nil {
			return errors.Wrap(_errorCodeLengthErr, "Error serializing 'errorCodeLength' field")
		}

		// Array Field (errorCode)
		if m.ErrorCode != nil {
			if pushErr := io.PushContext("errorCode", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ErrorCode {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'errorCode' field")
				}
			}
			if popErr := io.PopContext("errorCode", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("BACnetErrorReadProperty"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetErrorReadProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "errorClassLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ErrorClassLength = data
			case "errorClass":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.ErrorClass = utils.ByteArrayToInt8Array(_decoded[0:_len])
			case "errorCodeLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ErrorCodeLength = data
			case "errorCode":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.ErrorCode = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetErrorReadProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ErrorClassLength, xml.StartElement{Name: xml.Name{Local: "errorClassLength"}}); err != nil {
		return err
	}
	_encodedErrorClass := hex.EncodeToString(utils.Int8ArrayToByteArray(m.ErrorClass))
	_encodedErrorClass = strings.ToUpper(_encodedErrorClass)
	if err := e.EncodeElement(_encodedErrorClass, xml.StartElement{Name: xml.Name{Local: "errorClass"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ErrorCodeLength, xml.StartElement{Name: xml.Name{Local: "errorCodeLength"}}); err != nil {
		return err
	}
	_encodedErrorCode := hex.EncodeToString(utils.Int8ArrayToByteArray(m.ErrorCode))
	_encodedErrorCode = strings.ToUpper(_encodedErrorCode)
	if err := e.EncodeElement(_encodedErrorCode, xml.StartElement{Name: xml.Name{Local: "errorCode"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetErrorReadProperty) String() string {
	return string(m.Box("", 120))
}

func (m BACnetErrorReadProperty) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetErrorReadProperty"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Const Field (errorClassHeader)
		boxes = append(boxes, utils.BoxAnything("ErrorClassHeader", uint8(0x12), -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ErrorClassLength", m.ErrorClassLength, -1))
		// Array Field (errorClass)
		if m.ErrorClass != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.ErrorClass {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("ErrorClass", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Const Field (errorCodeHeader)
		boxes = append(boxes, utils.BoxAnything("ErrorCodeHeader", uint8(0x12), -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ErrorCodeLength", m.ErrorCodeLength, -1))
		// Array Field (errorCode)
		if m.ErrorCode != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.ErrorCode {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("ErrorCode", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
