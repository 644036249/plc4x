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
type ServiceId struct {
	Child IServiceIdChild
}

// The corresponding interface
type IServiceId interface {
	ServiceType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IServiceIdParent interface {
	SerializeParent(io utils.WriteBuffer, child IServiceId, serializeChildFunction func() error) error
	GetTypeName() string
}

type IServiceIdChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *ServiceId)
	GetTypeName() string
	IServiceId
	utils.AsciiBoxer
}

func NewServiceId() *ServiceId {
	return &ServiceId{}
}

func CastServiceId(structType interface{}) *ServiceId {
	castFunc := func(typ interface{}) *ServiceId {
		if casted, ok := typ.(ServiceId); ok {
			return &casted
		}
		if casted, ok := typ.(*ServiceId); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ServiceId) GetTypeName() string {
	return "ServiceId"
}

func (m *ServiceId) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ServiceId) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ServiceId) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceType)
	lengthInBits += 8

	return lengthInBits
}

func (m *ServiceId) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ServiceIdParse(io utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := io.PullContext("ServiceId"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceType) (Used as input to a switch field)
	serviceType, _serviceTypeErr := io.ReadUint8("serviceType", 8)
	if _serviceTypeErr != nil {
		return nil, errors.Wrap(_serviceTypeErr, "Error parsing 'serviceType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ServiceId
	var typeSwitchError error
	switch {
	case serviceType == 0x02: // KnxNetIpCore
		_parent, typeSwitchError = KnxNetIpCoreParse(io)
	case serviceType == 0x03: // KnxNetIpDeviceManagement
		_parent, typeSwitchError = KnxNetIpDeviceManagementParse(io)
	case serviceType == 0x04: // KnxNetIpTunneling
		_parent, typeSwitchError = KnxNetIpTunnelingParse(io)
	case serviceType == 0x05: // KnxNetIpRouting
		_parent, typeSwitchError = KnxNetIpRoutingParse(io)
	case serviceType == 0x06: // KnxNetRemoteLogging
		_parent, typeSwitchError = KnxNetRemoteLoggingParse(io)
	case serviceType == 0x07: // KnxNetRemoteConfigurationAndDiagnosis
		_parent, typeSwitchError = KnxNetRemoteConfigurationAndDiagnosisParse(io)
	case serviceType == 0x08: // KnxNetObjectServer
		_parent, typeSwitchError = KnxNetObjectServerParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("ServiceId"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ServiceId) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *ServiceId) SerializeParent(io utils.WriteBuffer, child IServiceId, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("ServiceId"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceType) (Used as input to a switch field)
	serviceType := uint8(child.ServiceType())
	_serviceTypeErr := io.WriteUint8("serviceType", 8, (serviceType))

	if _serviceTypeErr != nil {
		return errors.Wrap(_serviceTypeErr, "Error serializing 'serviceType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("ServiceId"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ServiceId) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
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
					panic("Couldn't determine class type for childs of ServiceId")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetIpCore":
					var dt *KnxNetIpCore
					if m.Child != nil {
						dt = m.Child.(*KnxNetIpCore)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetIpDeviceManagement":
					var dt *KnxNetIpDeviceManagement
					if m.Child != nil {
						dt = m.Child.(*KnxNetIpDeviceManagement)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetIpTunneling":
					var dt *KnxNetIpTunneling
					if m.Child != nil {
						dt = m.Child.(*KnxNetIpTunneling)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetIpRouting":
					var dt *KnxNetIpRouting
					if m.Child != nil {
						dt = m.Child.(*KnxNetIpRouting)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetRemoteLogging":
					var dt *KnxNetRemoteLogging
					if m.Child != nil {
						dt = m.Child.(*KnxNetRemoteLogging)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetRemoteConfigurationAndDiagnosis":
					var dt *KnxNetRemoteConfigurationAndDiagnosis
					if m.Child != nil {
						dt = m.Child.(*KnxNetRemoteConfigurationAndDiagnosis)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.KnxNetObjectServer":
					var dt *KnxNetObjectServer
					if m.Child != nil {
						dt = m.Child.(*KnxNetObjectServer)
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

func (m *ServiceId) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (m ServiceId) String() string {
	return string(m.Box("", 120))
}

func (m *ServiceId) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *ServiceId) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "ServiceId"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Discriminator Field (serviceType) (Used as input to a switch field)
	serviceType := uint8(m.Child.ServiceType())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ServiceType", serviceType, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
