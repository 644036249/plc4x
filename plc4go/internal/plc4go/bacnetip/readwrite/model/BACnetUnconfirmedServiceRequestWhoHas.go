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
const BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER uint8 = 0x0B
const BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER uint8 = 0x1B
const BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER uint8 = 0x3D

// The data-structure of this message
type BACnetUnconfirmedServiceRequestWhoHas struct {
	DeviceInstanceLowLimit  uint32
	DeviceInstanceHighLimit uint32
	ObjectNameCharacterSet  uint8
	ObjectName              []int8
	Parent                  *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWhoHas interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestWhoHas) ServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestWhoHas(deviceInstanceLowLimit uint32, deviceInstanceHighLimit uint32, objectNameCharacterSet uint8, objectName []int8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceLowLimit:  deviceInstanceLowLimit,
		DeviceInstanceHighLimit: deviceInstanceHighLimit,
		ObjectNameCharacterSet:  objectNameCharacterSet,
		ObjectName:              objectName,
		Parent:                  NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestWhoHas(structType interface{}) *BACnetUnconfirmedServiceRequestWhoHas {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestWhoHas {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestWhoHas); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestWhoHas); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHas"
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (deviceInstanceLowLimitHeader)
	lengthInBits += 8

	// Simple field (deviceInstanceLowLimit)
	lengthInBits += 24

	// Const Field (deviceInstanceHighLimitHeader)
	lengthInBits += 8

	// Simple field (deviceInstanceHighLimit)
	lengthInBits += 24

	// Const Field (objectNameHeader)
	lengthInBits += 8

	// Implicit Field (objectNameLength)
	lengthInBits += 8

	// Simple field (objectNameCharacterSet)
	lengthInBits += 8

	// Array field
	if len(m.ObjectName) > 0 {
		lengthInBits += 8 * uint16(len(m.ObjectName))
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoHasParse(io utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := io.PullContext("BACnetUnconfirmedServiceRequestWhoHas"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (deviceInstanceLowLimitHeader)
	deviceInstanceLowLimitHeader, _deviceInstanceLowLimitHeaderErr := io.ReadUint8("deviceInstanceLowLimitHeader", 8)
	if _deviceInstanceLowLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceLowLimitHeaderErr, "Error parsing 'deviceInstanceLowLimitHeader' field")
	}
	if deviceInstanceLowLimitHeader != BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceLowLimitHeader))
	}

	// Simple Field (deviceInstanceLowLimit)
	deviceInstanceLowLimit, _deviceInstanceLowLimitErr := io.ReadUint32("deviceInstanceLowLimit", 24)
	if _deviceInstanceLowLimitErr != nil {
		return nil, errors.Wrap(_deviceInstanceLowLimitErr, "Error parsing 'deviceInstanceLowLimit' field")
	}

	// Const Field (deviceInstanceHighLimitHeader)
	deviceInstanceHighLimitHeader, _deviceInstanceHighLimitHeaderErr := io.ReadUint8("deviceInstanceHighLimitHeader", 8)
	if _deviceInstanceHighLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceHighLimitHeaderErr, "Error parsing 'deviceInstanceHighLimitHeader' field")
	}
	if deviceInstanceHighLimitHeader != BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceHighLimitHeader))
	}

	// Simple Field (deviceInstanceHighLimit)
	deviceInstanceHighLimit, _deviceInstanceHighLimitErr := io.ReadUint32("deviceInstanceHighLimit", 24)
	if _deviceInstanceHighLimitErr != nil {
		return nil, errors.Wrap(_deviceInstanceHighLimitErr, "Error parsing 'deviceInstanceHighLimit' field")
	}

	// Const Field (objectNameHeader)
	objectNameHeader, _objectNameHeaderErr := io.ReadUint8("objectNameHeader", 8)
	if _objectNameHeaderErr != nil {
		return nil, errors.Wrap(_objectNameHeaderErr, "Error parsing 'objectNameHeader' field")
	}
	if objectNameHeader != BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER) + " but got " + fmt.Sprintf("%d", objectNameHeader))
	}

	// Implicit Field (objectNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	objectNameLength, _objectNameLengthErr := io.ReadUint8("objectNameLength", 8)
	_ = objectNameLength
	if _objectNameLengthErr != nil {
		return nil, errors.Wrap(_objectNameLengthErr, "Error parsing 'objectNameLength' field")
	}

	// Simple Field (objectNameCharacterSet)
	objectNameCharacterSet, _objectNameCharacterSetErr := io.ReadUint8("objectNameCharacterSet", 8)
	if _objectNameCharacterSetErr != nil {
		return nil, errors.Wrap(_objectNameCharacterSetErr, "Error parsing 'objectNameCharacterSet' field")
	}

	// Array field (objectName)
	if pullErr := io.PullContext("objectName", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	objectName := make([]int8, 0)
	_objectNameLength := uint16(objectNameLength) - uint16(uint16(1))
	_objectNameEndPos := io.GetPos() + uint16(_objectNameLength)
	for io.GetPos() < _objectNameEndPos {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'objectName' field")
		}
		objectName = append(objectName, _item)
	}
	if closeErr := io.CloseContext("objectName", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("BACnetUnconfirmedServiceRequestWhoHas"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceLowLimit:  deviceInstanceLowLimit,
		DeviceInstanceHighLimit: deviceInstanceHighLimit,
		ObjectNameCharacterSet:  objectNameCharacterSet,
		ObjectName:              objectName,
		Parent:                  &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("BACnetUnconfirmedServiceRequestWhoHas"); pushErr != nil {
			return pushErr
		}

		// Const Field (deviceInstanceLowLimitHeader)
		_deviceInstanceLowLimitHeaderErr := io.WriteUint8("deviceInstanceLowLimitHeader", 8, 0x0B)
		if _deviceInstanceLowLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceLowLimitHeaderErr, "Error serializing 'deviceInstanceLowLimitHeader' field")
		}

		// Simple Field (deviceInstanceLowLimit)
		deviceInstanceLowLimit := uint32(m.DeviceInstanceLowLimit)
		_deviceInstanceLowLimitErr := io.WriteUint32("deviceInstanceLowLimit", 24, (deviceInstanceLowLimit))
		if _deviceInstanceLowLimitErr != nil {
			return errors.Wrap(_deviceInstanceLowLimitErr, "Error serializing 'deviceInstanceLowLimit' field")
		}

		// Const Field (deviceInstanceHighLimitHeader)
		_deviceInstanceHighLimitHeaderErr := io.WriteUint8("deviceInstanceHighLimitHeader", 8, 0x1B)
		if _deviceInstanceHighLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceHighLimitHeaderErr, "Error serializing 'deviceInstanceHighLimitHeader' field")
		}

		// Simple Field (deviceInstanceHighLimit)
		deviceInstanceHighLimit := uint32(m.DeviceInstanceHighLimit)
		_deviceInstanceHighLimitErr := io.WriteUint32("deviceInstanceHighLimit", 24, (deviceInstanceHighLimit))
		if _deviceInstanceHighLimitErr != nil {
			return errors.Wrap(_deviceInstanceHighLimitErr, "Error serializing 'deviceInstanceHighLimit' field")
		}

		// Const Field (objectNameHeader)
		_objectNameHeaderErr := io.WriteUint8("objectNameHeader", 8, 0x3D)
		if _objectNameHeaderErr != nil {
			return errors.Wrap(_objectNameHeaderErr, "Error serializing 'objectNameHeader' field")
		}

		// Implicit Field (objectNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		objectNameLength := uint8(uint8(uint8(len(m.ObjectName))) + uint8(uint8(1)))
		_objectNameLengthErr := io.WriteUint8("objectNameLength", 8, (objectNameLength))
		if _objectNameLengthErr != nil {
			return errors.Wrap(_objectNameLengthErr, "Error serializing 'objectNameLength' field")
		}

		// Simple Field (objectNameCharacterSet)
		objectNameCharacterSet := uint8(m.ObjectNameCharacterSet)
		_objectNameCharacterSetErr := io.WriteUint8("objectNameCharacterSet", 8, (objectNameCharacterSet))
		if _objectNameCharacterSetErr != nil {
			return errors.Wrap(_objectNameCharacterSetErr, "Error serializing 'objectNameCharacterSet' field")
		}

		// Array Field (objectName)
		if m.ObjectName != nil {
			if pushErr := io.PushContext("objectName", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ObjectName {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'objectName' field")
				}
			}
			if popErr := io.PopContext("objectName", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("BACnetUnconfirmedServiceRequestWhoHas"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "deviceInstanceLowLimit":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceInstanceLowLimit = data
			case "deviceInstanceHighLimit":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceInstanceHighLimit = data
			case "objectNameCharacterSet":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ObjectNameCharacterSet = data
			case "objectName":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.ObjectName = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetUnconfirmedServiceRequestWhoHas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DeviceInstanceLowLimit, xml.StartElement{Name: xml.Name{Local: "deviceInstanceLowLimit"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DeviceInstanceHighLimit, xml.StartElement{Name: xml.Name{Local: "deviceInstanceHighLimit"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ObjectNameCharacterSet, xml.StartElement{Name: xml.Name{Local: "objectNameCharacterSet"}}); err != nil {
		return err
	}
	_encodedObjectName := hex.EncodeToString(utils.Int8ArrayToByteArray(m.ObjectName))
	_encodedObjectName = strings.ToUpper(_encodedObjectName)
	if err := e.EncodeElement(_encodedObjectName, xml.StartElement{Name: xml.Name{Local: "objectName"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetUnconfirmedServiceRequestWhoHas) String() string {
	return string(m.Box("", 120))
}

func (m BACnetUnconfirmedServiceRequestWhoHas) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetUnconfirmedServiceRequestWhoHas"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Const Field (deviceInstanceLowLimitHeader)
		boxes = append(boxes, utils.BoxAnything("DeviceInstanceLowLimitHeader", uint8(0x0B), -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DeviceInstanceLowLimit", m.DeviceInstanceLowLimit, -1))
		// Const Field (deviceInstanceHighLimitHeader)
		boxes = append(boxes, utils.BoxAnything("DeviceInstanceHighLimitHeader", uint8(0x1B), -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DeviceInstanceHighLimit", m.DeviceInstanceHighLimit, -1))
		// Const Field (objectNameHeader)
		boxes = append(boxes, utils.BoxAnything("ObjectNameHeader", uint8(0x3D), -1))
		// Implicit Field (objectNameLength)
		objectNameLength := uint8(uint8(uint8(len(m.ObjectName))) + uint8(uint8(1)))
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ObjectNameLength", objectNameLength, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ObjectNameCharacterSet", m.ObjectNameCharacterSet, -1))
		// Array Field (objectName)
		if m.ObjectName != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.ObjectName {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("ObjectName", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
