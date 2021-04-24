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
type BACnetConfirmedServiceACK struct {
	Child IBACnetConfirmedServiceACKChild
}

// The corresponding interface
type IBACnetConfirmedServiceACK interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IBACnetConfirmedServiceACKParent interface {
	SerializeParent(io utils.WriteBuffer, child IBACnetConfirmedServiceACK, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConfirmedServiceACKChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *BACnetConfirmedServiceACK)
	GetTypeName() string
	IBACnetConfirmedServiceACK
	utils.AsciiBoxer
}

func NewBACnetConfirmedServiceACK() *BACnetConfirmedServiceACK {
	return &BACnetConfirmedServiceACK{}
}

func CastBACnetConfirmedServiceACK(structType interface{}) *BACnetConfirmedServiceACK {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceACK {
		if casted, ok := typ.(BACnetConfirmedServiceACK); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACK); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceACK) GetTypeName() string {
	return "BACnetConfirmedServiceACK"
}

func (m *BACnetConfirmedServiceACK) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACK) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetConfirmedServiceACK) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetConfirmedServiceACK) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKParse(io utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := io.PullContext("BACnetConfirmedServiceACK"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := io.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConfirmedServiceACK
	var typeSwitchError error
	switch {
	case serviceChoice == 0x03: // BACnetConfirmedServiceACKGetAlarmSummary
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetAlarmSummaryParse(io)
	case serviceChoice == 0x04: // BACnetConfirmedServiceACKGetEnrollmentSummary
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetEnrollmentSummaryParse(io)
	case serviceChoice == 0x1D: // BACnetConfirmedServiceACKGetEventInformation
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetEventInformationParse(io)
	case serviceChoice == 0x06: // BACnetConfirmedServiceACKAtomicReadFile
		_parent, typeSwitchError = BACnetConfirmedServiceACKAtomicReadFileParse(io)
	case serviceChoice == 0x07: // BACnetConfirmedServiceACKAtomicWriteFile
		_parent, typeSwitchError = BACnetConfirmedServiceACKAtomicWriteFileParse(io)
	case serviceChoice == 0x0A: // BACnetConfirmedServiceACKCreateObject
		_parent, typeSwitchError = BACnetConfirmedServiceACKCreateObjectParse(io)
	case serviceChoice == 0x0C: // BACnetConfirmedServiceACKReadProperty
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadPropertyParse(io)
	case serviceChoice == 0x0E: // BACnetConfirmedServiceACKReadPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadPropertyMultipleParse(io)
	case serviceChoice == 0x1A: // BACnetConfirmedServiceACKReadRange
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadRangeParse(io)
	case serviceChoice == 0x12: // BACnetConfirmedServiceACKConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetConfirmedServiceACKConfirmedPrivateTransferParse(io)
	case serviceChoice == 0x15: // BACnetConfirmedServiceACKVTOpen
		_parent, typeSwitchError = BACnetConfirmedServiceACKVTOpenParse(io)
	case serviceChoice == 0x17: // BACnetConfirmedServiceACKVTData
		_parent, typeSwitchError = BACnetConfirmedServiceACKVTDataParse(io)
	case serviceChoice == 0x18: // BACnetConfirmedServiceACKRemovedAuthenticate
		_parent, typeSwitchError = BACnetConfirmedServiceACKRemovedAuthenticateParse(io)
	case serviceChoice == 0x0D: // BACnetConfirmedServiceACKRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetConfirmedServiceACKRemovedReadPropertyConditionalParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("BACnetConfirmedServiceACK"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetConfirmedServiceACK) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *BACnetConfirmedServiceACK) SerializeParent(io utils.WriteBuffer, child IBACnetConfirmedServiceACK, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("BACnetConfirmedServiceACK"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.ServiceChoice())
	_serviceChoiceErr := io.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("BACnetConfirmedServiceACK"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceACK) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// BACnetConfirmedServiceACKGetAlarmSummary needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetAlarmSummary":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKGetAlarmSummary{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKGetEnrollmentSummary needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetEnrollmentSummary":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKGetEnrollmentSummary{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKGetEventInformation needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetEventInformation":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKGetEventInformation{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKAtomicReadFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKAtomicReadFile":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKAtomicReadFile{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKAtomicWriteFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKAtomicWriteFile":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKAtomicWriteFile{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKCreateObject needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKCreateObject":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKCreateObject{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKReadProperty needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadProperty":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKReadProperty{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKReadPropertyMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadPropertyMultiple":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKReadPropertyMultiple{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKReadRange needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadRange":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKReadRange{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKConfirmedPrivateTransfer needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKConfirmedPrivateTransfer":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKConfirmedPrivateTransfer{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKVTOpen needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKVTOpen":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKVTOpen{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKVTData needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKVTData":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKVTData{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKRemovedAuthenticate needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKRemovedAuthenticate":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKRemovedAuthenticate{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceACKRemovedReadPropertyConditional needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKRemovedReadPropertyConditional":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceACKRemovedReadPropertyConditional{
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
					panic("Couldn't determine class type for childs of BACnetConfirmedServiceACK")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetAlarmSummary":
					var dt *BACnetConfirmedServiceACKGetAlarmSummary
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKGetAlarmSummary)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetEnrollmentSummary":
					var dt *BACnetConfirmedServiceACKGetEnrollmentSummary
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKGetEnrollmentSummary)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetEventInformation":
					var dt *BACnetConfirmedServiceACKGetEventInformation
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKGetEventInformation)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKAtomicReadFile":
					var dt *BACnetConfirmedServiceACKAtomicReadFile
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKAtomicReadFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKAtomicWriteFile":
					var dt *BACnetConfirmedServiceACKAtomicWriteFile
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKAtomicWriteFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKCreateObject":
					var dt *BACnetConfirmedServiceACKCreateObject
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKCreateObject)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadProperty":
					var dt *BACnetConfirmedServiceACKReadProperty
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKReadProperty)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadPropertyMultiple":
					var dt *BACnetConfirmedServiceACKReadPropertyMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKReadPropertyMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadRange":
					var dt *BACnetConfirmedServiceACKReadRange
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKReadRange)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKConfirmedPrivateTransfer":
					var dt *BACnetConfirmedServiceACKConfirmedPrivateTransfer
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKConfirmedPrivateTransfer)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKVTOpen":
					var dt *BACnetConfirmedServiceACKVTOpen
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKVTOpen)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKVTData":
					var dt *BACnetConfirmedServiceACKVTData
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKVTData)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKRemovedAuthenticate":
					var dt *BACnetConfirmedServiceACKRemovedAuthenticate
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKRemovedAuthenticate)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKRemovedReadPropertyConditional":
					var dt *BACnetConfirmedServiceACKRemovedReadPropertyConditional
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceACKRemovedReadPropertyConditional)
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

func (m *BACnetConfirmedServiceACK) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
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

func (m BACnetConfirmedServiceACK) String() string {
	return string(m.Box("", 120))
}

func (m *BACnetConfirmedServiceACK) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *BACnetConfirmedServiceACK) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "BACnetConfirmedServiceACK"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(m.Child.ServiceChoice())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ServiceChoice", serviceChoice, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
