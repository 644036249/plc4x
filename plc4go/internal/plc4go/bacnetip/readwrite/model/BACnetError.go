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
type BACnetError struct {
	Child IBACnetErrorChild
}

// The corresponding interface
type IBACnetError interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IBACnetErrorParent interface {
	SerializeParent(io utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetErrorChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *BACnetError)
	GetTypeName() string
	IBACnetError
	utils.AsciiBoxer
}

func NewBACnetError() *BACnetError {
	return &BACnetError{}
}

func CastBACnetError(structType interface{}) *BACnetError {
	castFunc := func(typ interface{}) *BACnetError {
		if casted, ok := typ.(BACnetError); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetError); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetError) GetTypeName() string {
	return "BACnetError"
}

func (m *BACnetError) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetError) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetError) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetError) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorParse(io utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := io.PullContext("BACnetError"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := io.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetError
	var typeSwitchError error
	switch {
	case serviceChoice == 0x03: // BACnetErrorGetAlarmSummary
		_parent, typeSwitchError = BACnetErrorGetAlarmSummaryParse(io)
	case serviceChoice == 0x04: // BACnetErrorGetEnrollmentSummary
		_parent, typeSwitchError = BACnetErrorGetEnrollmentSummaryParse(io)
	case serviceChoice == 0x1D: // BACnetErrorGetEventInformation
		_parent, typeSwitchError = BACnetErrorGetEventInformationParse(io)
	case serviceChoice == 0x06: // BACnetErrorAtomicReadFile
		_parent, typeSwitchError = BACnetErrorAtomicReadFileParse(io)
	case serviceChoice == 0x07: // BACnetErrorAtomicWriteFile
		_parent, typeSwitchError = BACnetErrorAtomicWriteFileParse(io)
	case serviceChoice == 0x0A: // BACnetErrorCreateObject
		_parent, typeSwitchError = BACnetErrorCreateObjectParse(io)
	case serviceChoice == 0x0C: // BACnetErrorReadProperty
		_parent, typeSwitchError = BACnetErrorReadPropertyParse(io)
	case serviceChoice == 0x0E: // BACnetErrorReadPropertyMultiple
		_parent, typeSwitchError = BACnetErrorReadPropertyMultipleParse(io)
	case serviceChoice == 0x1A: // BACnetErrorReadRange
		_parent, typeSwitchError = BACnetErrorReadRangeParse(io)
	case serviceChoice == 0x12: // BACnetErrorConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetErrorConfirmedPrivateTransferParse(io)
	case serviceChoice == 0x15: // BACnetErrorVTOpen
		_parent, typeSwitchError = BACnetErrorVTOpenParse(io)
	case serviceChoice == 0x17: // BACnetErrorVTData
		_parent, typeSwitchError = BACnetErrorVTDataParse(io)
	case serviceChoice == 0x18: // BACnetErrorRemovedAuthenticate
		_parent, typeSwitchError = BACnetErrorRemovedAuthenticateParse(io)
	case serviceChoice == 0x0D: // BACnetErrorRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetErrorRemovedReadPropertyConditionalParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("BACnetError"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetError) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *BACnetError) SerializeParent(io utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("BACnetError"); pushErr != nil {
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

	if popErr := io.PopContext("BACnetError"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// BACnetErrorGetAlarmSummary needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetAlarmSummary":
			if m.Child == nil {
				m.Child = &BACnetErrorGetAlarmSummary{
					Parent: m,
				}
			}
		// BACnetErrorGetEnrollmentSummary needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEnrollmentSummary":
			if m.Child == nil {
				m.Child = &BACnetErrorGetEnrollmentSummary{
					Parent: m,
				}
			}
		// BACnetErrorGetEventInformation needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEventInformation":
			if m.Child == nil {
				m.Child = &BACnetErrorGetEventInformation{
					Parent: m,
				}
			}
		// BACnetErrorAtomicReadFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicReadFile":
			if m.Child == nil {
				m.Child = &BACnetErrorAtomicReadFile{
					Parent: m,
				}
			}
		// BACnetErrorAtomicWriteFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicWriteFile":
			if m.Child == nil {
				m.Child = &BACnetErrorAtomicWriteFile{
					Parent: m,
				}
			}
		// BACnetErrorCreateObject needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorCreateObject":
			if m.Child == nil {
				m.Child = &BACnetErrorCreateObject{
					Parent: m,
				}
			}
		// BACnetErrorReadPropertyMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadPropertyMultiple":
			if m.Child == nil {
				m.Child = &BACnetErrorReadPropertyMultiple{
					Parent: m,
				}
			}
		// BACnetErrorReadRange needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadRange":
			if m.Child == nil {
				m.Child = &BACnetErrorReadRange{
					Parent: m,
				}
			}
		// BACnetErrorConfirmedPrivateTransfer needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorConfirmedPrivateTransfer":
			if m.Child == nil {
				m.Child = &BACnetErrorConfirmedPrivateTransfer{
					Parent: m,
				}
			}
		// BACnetErrorVTOpen needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTOpen":
			if m.Child == nil {
				m.Child = &BACnetErrorVTOpen{
					Parent: m,
				}
			}
		// BACnetErrorVTData needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTData":
			if m.Child == nil {
				m.Child = &BACnetErrorVTData{
					Parent: m,
				}
			}
		// BACnetErrorRemovedAuthenticate needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedAuthenticate":
			if m.Child == nil {
				m.Child = &BACnetErrorRemovedAuthenticate{
					Parent: m,
				}
			}
		// BACnetErrorRemovedReadPropertyConditional needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedReadPropertyConditional":
			if m.Child == nil {
				m.Child = &BACnetErrorRemovedReadPropertyConditional{
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
					panic("Couldn't determine class type for childs of BACnetError")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetAlarmSummary":
					var dt *BACnetErrorGetAlarmSummary
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorGetAlarmSummary)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEnrollmentSummary":
					var dt *BACnetErrorGetEnrollmentSummary
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorGetEnrollmentSummary)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEventInformation":
					var dt *BACnetErrorGetEventInformation
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorGetEventInformation)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicReadFile":
					var dt *BACnetErrorAtomicReadFile
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorAtomicReadFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicWriteFile":
					var dt *BACnetErrorAtomicWriteFile
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorAtomicWriteFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorCreateObject":
					var dt *BACnetErrorCreateObject
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorCreateObject)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadProperty":
					var dt *BACnetErrorReadProperty
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorReadProperty)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadPropertyMultiple":
					var dt *BACnetErrorReadPropertyMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorReadPropertyMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadRange":
					var dt *BACnetErrorReadRange
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorReadRange)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorConfirmedPrivateTransfer":
					var dt *BACnetErrorConfirmedPrivateTransfer
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorConfirmedPrivateTransfer)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTOpen":
					var dt *BACnetErrorVTOpen
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorVTOpen)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTData":
					var dt *BACnetErrorVTData
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorVTData)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedAuthenticate":
					var dt *BACnetErrorRemovedAuthenticate
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorRemovedAuthenticate)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedReadPropertyConditional":
					var dt *BACnetErrorRemovedReadPropertyConditional
					if m.Child != nil {
						dt = m.Child.(*BACnetErrorRemovedReadPropertyConditional)
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

func (m *BACnetError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (m BACnetError) String() string {
	return string(m.Box("", 120))
}

func (m *BACnetError) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *BACnetError) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "BACnetError"
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
