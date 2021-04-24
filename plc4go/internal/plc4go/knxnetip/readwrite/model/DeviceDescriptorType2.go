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
type DeviceDescriptorType2 struct {
	ManufacturerId uint16
	DeviceType     uint16
	Version        uint8
	ReadSupported  bool
	WriteSupported bool
	LogicalTagBase uint8
	ChannelInfo1   *ChannelInformation
	ChannelInfo2   *ChannelInformation
	ChannelInfo3   *ChannelInformation
	ChannelInfo4   *ChannelInformation
}

// The corresponding interface
type IDeviceDescriptorType2 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewDeviceDescriptorType2(manufacturerId uint16, deviceType uint16, version uint8, readSupported bool, writeSupported bool, logicalTagBase uint8, channelInfo1 *ChannelInformation, channelInfo2 *ChannelInformation, channelInfo3 *ChannelInformation, channelInfo4 *ChannelInformation) *DeviceDescriptorType2 {
	return &DeviceDescriptorType2{ManufacturerId: manufacturerId, DeviceType: deviceType, Version: version, ReadSupported: readSupported, WriteSupported: writeSupported, LogicalTagBase: logicalTagBase, ChannelInfo1: channelInfo1, ChannelInfo2: channelInfo2, ChannelInfo3: channelInfo3, ChannelInfo4: channelInfo4}
}

func CastDeviceDescriptorType2(structType interface{}) *DeviceDescriptorType2 {
	castFunc := func(typ interface{}) *DeviceDescriptorType2 {
		if casted, ok := typ.(DeviceDescriptorType2); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceDescriptorType2); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceDescriptorType2) GetTypeName() string {
	return "DeviceDescriptorType2"
}

func (m *DeviceDescriptorType2) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DeviceDescriptorType2) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (manufacturerId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (version)
	lengthInBits += 8

	// Simple field (readSupported)
	lengthInBits += 1

	// Simple field (writeSupported)
	lengthInBits += 1

	// Simple field (logicalTagBase)
	lengthInBits += 6

	// Simple field (channelInfo1)
	lengthInBits += m.ChannelInfo1.LengthInBits()

	// Simple field (channelInfo2)
	lengthInBits += m.ChannelInfo2.LengthInBits()

	// Simple field (channelInfo3)
	lengthInBits += m.ChannelInfo3.LengthInBits()

	// Simple field (channelInfo4)
	lengthInBits += m.ChannelInfo4.LengthInBits()

	return lengthInBits
}

func (m *DeviceDescriptorType2) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceDescriptorType2Parse(io utils.ReadBuffer) (*DeviceDescriptorType2, error) {
	if pullErr := io.PullContext("DeviceDescriptorType2"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (manufacturerId)
	manufacturerId, _manufacturerIdErr := io.ReadUint16("manufacturerId", 16)
	if _manufacturerIdErr != nil {
		return nil, errors.Wrap(_manufacturerIdErr, "Error parsing 'manufacturerId' field")
	}

	// Simple Field (deviceType)
	deviceType, _deviceTypeErr := io.ReadUint16("deviceType", 16)
	if _deviceTypeErr != nil {
		return nil, errors.Wrap(_deviceTypeErr, "Error parsing 'deviceType' field")
	}

	// Simple Field (version)
	version, _versionErr := io.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	// Simple Field (readSupported)
	readSupported, _readSupportedErr := io.ReadBit("readSupported")
	if _readSupportedErr != nil {
		return nil, errors.Wrap(_readSupportedErr, "Error parsing 'readSupported' field")
	}

	// Simple Field (writeSupported)
	writeSupported, _writeSupportedErr := io.ReadBit("writeSupported")
	if _writeSupportedErr != nil {
		return nil, errors.Wrap(_writeSupportedErr, "Error parsing 'writeSupported' field")
	}

	// Simple Field (logicalTagBase)
	logicalTagBase, _logicalTagBaseErr := io.ReadUint8("logicalTagBase", 6)
	if _logicalTagBaseErr != nil {
		return nil, errors.Wrap(_logicalTagBaseErr, "Error parsing 'logicalTagBase' field")
	}

	if pullErr := io.PullContext("channelInfo1"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo1)
	channelInfo1, _channelInfo1Err := ChannelInformationParse(io)
	if _channelInfo1Err != nil {
		return nil, errors.Wrap(_channelInfo1Err, "Error parsing 'channelInfo1' field")
	}
	if closeErr := io.CloseContext("channelInfo1"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("channelInfo2"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo2)
	channelInfo2, _channelInfo2Err := ChannelInformationParse(io)
	if _channelInfo2Err != nil {
		return nil, errors.Wrap(_channelInfo2Err, "Error parsing 'channelInfo2' field")
	}
	if closeErr := io.CloseContext("channelInfo2"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("channelInfo3"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo3)
	channelInfo3, _channelInfo3Err := ChannelInformationParse(io)
	if _channelInfo3Err != nil {
		return nil, errors.Wrap(_channelInfo3Err, "Error parsing 'channelInfo3' field")
	}
	if closeErr := io.CloseContext("channelInfo3"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("channelInfo4"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo4)
	channelInfo4, _channelInfo4Err := ChannelInformationParse(io)
	if _channelInfo4Err != nil {
		return nil, errors.Wrap(_channelInfo4Err, "Error parsing 'channelInfo4' field")
	}
	if closeErr := io.CloseContext("channelInfo4"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("DeviceDescriptorType2"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDeviceDescriptorType2(manufacturerId, deviceType, version, readSupported, writeSupported, logicalTagBase, channelInfo1, channelInfo2, channelInfo3, channelInfo4), nil
}

func (m *DeviceDescriptorType2) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("DeviceDescriptorType2"); pushErr != nil {
		return pushErr
	}

	// Simple Field (manufacturerId)
	manufacturerId := uint16(m.ManufacturerId)
	_manufacturerIdErr := io.WriteUint16("manufacturerId", 16, (manufacturerId))
	if _manufacturerIdErr != nil {
		return errors.Wrap(_manufacturerIdErr, "Error serializing 'manufacturerId' field")
	}

	// Simple Field (deviceType)
	deviceType := uint16(m.DeviceType)
	_deviceTypeErr := io.WriteUint16("deviceType", 16, (deviceType))
	if _deviceTypeErr != nil {
		return errors.Wrap(_deviceTypeErr, "Error serializing 'deviceType' field")
	}

	// Simple Field (version)
	version := uint8(m.Version)
	_versionErr := io.WriteUint8("version", 8, (version))
	if _versionErr != nil {
		return errors.Wrap(_versionErr, "Error serializing 'version' field")
	}

	// Simple Field (readSupported)
	readSupported := bool(m.ReadSupported)
	_readSupportedErr := io.WriteBit("readSupported", (readSupported))
	if _readSupportedErr != nil {
		return errors.Wrap(_readSupportedErr, "Error serializing 'readSupported' field")
	}

	// Simple Field (writeSupported)
	writeSupported := bool(m.WriteSupported)
	_writeSupportedErr := io.WriteBit("writeSupported", (writeSupported))
	if _writeSupportedErr != nil {
		return errors.Wrap(_writeSupportedErr, "Error serializing 'writeSupported' field")
	}

	// Simple Field (logicalTagBase)
	logicalTagBase := uint8(m.LogicalTagBase)
	_logicalTagBaseErr := io.WriteUint8("logicalTagBase", 6, (logicalTagBase))
	if _logicalTagBaseErr != nil {
		return errors.Wrap(_logicalTagBaseErr, "Error serializing 'logicalTagBase' field")
	}

	// Simple Field (channelInfo1)
	if pushErr := io.PushContext("channelInfo1"); pushErr != nil {
		return pushErr
	}
	_channelInfo1Err := m.ChannelInfo1.Serialize(io)
	if popErr := io.PopContext("channelInfo1"); popErr != nil {
		return popErr
	}
	if _channelInfo1Err != nil {
		return errors.Wrap(_channelInfo1Err, "Error serializing 'channelInfo1' field")
	}

	// Simple Field (channelInfo2)
	if pushErr := io.PushContext("channelInfo2"); pushErr != nil {
		return pushErr
	}
	_channelInfo2Err := m.ChannelInfo2.Serialize(io)
	if popErr := io.PopContext("channelInfo2"); popErr != nil {
		return popErr
	}
	if _channelInfo2Err != nil {
		return errors.Wrap(_channelInfo2Err, "Error serializing 'channelInfo2' field")
	}

	// Simple Field (channelInfo3)
	if pushErr := io.PushContext("channelInfo3"); pushErr != nil {
		return pushErr
	}
	_channelInfo3Err := m.ChannelInfo3.Serialize(io)
	if popErr := io.PopContext("channelInfo3"); popErr != nil {
		return popErr
	}
	if _channelInfo3Err != nil {
		return errors.Wrap(_channelInfo3Err, "Error serializing 'channelInfo3' field")
	}

	// Simple Field (channelInfo4)
	if pushErr := io.PushContext("channelInfo4"); pushErr != nil {
		return pushErr
	}
	_channelInfo4Err := m.ChannelInfo4.Serialize(io)
	if popErr := io.PopContext("channelInfo4"); popErr != nil {
		return popErr
	}
	if _channelInfo4Err != nil {
		return errors.Wrap(_channelInfo4Err, "Error serializing 'channelInfo4' field")
	}

	if popErr := io.PopContext("DeviceDescriptorType2"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DeviceDescriptorType2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
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
			case "manufacturerId":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ManufacturerId = data
			case "deviceType":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceType = data
			case "version":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Version = data
			case "readSupported":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReadSupported = data
			case "writeSupported":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.WriteSupported = data
			case "logicalTagBase":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.LogicalTagBase = data
			case "channelInfo1":
				var data ChannelInformation
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ChannelInfo1 = &data
			case "channelInfo2":
				var data ChannelInformation
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ChannelInfo2 = &data
			case "channelInfo3":
				var data ChannelInformation
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ChannelInfo3 = &data
			case "channelInfo4":
				var data ChannelInformation
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ChannelInfo4 = &data
			}
		}
	}
}

func (m *DeviceDescriptorType2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.DeviceDescriptorType2"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ManufacturerId, xml.StartElement{Name: xml.Name{Local: "manufacturerId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DeviceType, xml.StartElement{Name: xml.Name{Local: "deviceType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Version, xml.StartElement{Name: xml.Name{Local: "version"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReadSupported, xml.StartElement{Name: xml.Name{Local: "readSupported"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.WriteSupported, xml.StartElement{Name: xml.Name{Local: "writeSupported"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.LogicalTagBase, xml.StartElement{Name: xml.Name{Local: "logicalTagBase"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ChannelInfo1, xml.StartElement{Name: xml.Name{Local: "channelInfo1"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ChannelInfo2, xml.StartElement{Name: xml.Name{Local: "channelInfo2"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ChannelInfo3, xml.StartElement{Name: xml.Name{Local: "channelInfo3"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ChannelInfo4, xml.StartElement{Name: xml.Name{Local: "channelInfo4"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m DeviceDescriptorType2) String() string {
	return string(m.Box("", 120))
}

func (m DeviceDescriptorType2) Box(name string, width int) utils.AsciiBox {
	boxName := "DeviceDescriptorType2"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ManufacturerId", m.ManufacturerId, -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("DeviceType", m.DeviceType, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Version", m.Version, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ReadSupported", m.ReadSupported, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("WriteSupported", m.WriteSupported, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("LogicalTagBase", m.LogicalTagBase, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.ChannelInfo1.Box("channelInfo1", width-2))
	// Complex field (case complex)
	boxes = append(boxes, m.ChannelInfo2.Box("channelInfo2", width-2))
	// Complex field (case complex)
	boxes = append(boxes, m.ChannelInfo3.Box("channelInfo3", width-2))
	// Complex field (case complex)
	boxes = append(boxes, m.ChannelInfo4.Box("channelInfo4", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
