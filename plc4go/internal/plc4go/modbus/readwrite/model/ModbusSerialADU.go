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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusSerialADU struct {
	TransactionId uint16
	Length        uint16
	Address       uint8
	Pdu           *ModbusPDU
}

// The corresponding interface
type IModbusSerialADU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewModbusSerialADU(transactionId uint16, length uint16, address uint8, pdu *ModbusPDU) *ModbusSerialADU {
	return &ModbusSerialADU{TransactionId: transactionId, Length: length, Address: address, Pdu: pdu}
}

func CastModbusSerialADU(structType interface{}) *ModbusSerialADU {
	castFunc := func(typ interface{}) *ModbusSerialADU {
		if casted, ok := typ.(ModbusSerialADU); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusSerialADU); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusSerialADU) GetTypeName() string {
	return "ModbusSerialADU"
}

func (m *ModbusSerialADU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusSerialADU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (transactionId)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (length)
	lengthInBits += 16

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.LengthInBits()

	return lengthInBits
}

func (m *ModbusSerialADU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusSerialADUParse(io utils.ReadBuffer, response bool) (*ModbusSerialADU, error) {
	if pullErr := io.PullContext("ModbusSerialADU"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (transactionId)
	transactionId, _transactionIdErr := io.ReadUint16("transactionId", 16)
	if _transactionIdErr != nil {
		return nil, errors.Wrap(_transactionIdErr, "Error parsing 'transactionId' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (length)
	length, _lengthErr := io.ReadUint16("length", 16)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (address)
	address, _addressErr := io.ReadUint8("address", 8)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}

	if pullErr := io.PullContext("pdu"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pdu)
	pdu, _pduErr := ModbusPDUParse(io, response)
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field")
	}
	if closeErr := io.CloseContext("pdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("ModbusSerialADU"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusSerialADU(transactionId, length, address, pdu), nil
}

func (m *ModbusSerialADU) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("ModbusSerialADU"); pushErr != nil {
		return pushErr
	}

	// Simple Field (transactionId)
	transactionId := uint16(m.TransactionId)
	_transactionIdErr := io.WriteUint16("transactionId", 16, (transactionId))
	if _transactionIdErr != nil {
		return errors.Wrap(_transactionIdErr, "Error serializing 'transactionId' field")
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (length)
	length := uint16(m.Length)
	_lengthErr := io.WriteUint16("length", 16, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (address)
	address := uint8(m.Address)
	_addressErr := io.WriteUint8("address", 8, (address))
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}

	// Simple Field (pdu)
	if pushErr := io.PushContext("pdu"); pushErr != nil {
		return pushErr
	}
	_pduErr := m.Pdu.Serialize(io)
	if popErr := io.PopContext("pdu"); popErr != nil {
		return popErr
	}
	if _pduErr != nil {
		return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
	}

	if popErr := io.PopContext("ModbusSerialADU"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusSerialADU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "transactionId":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransactionId = data
			case "length":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Length = data
			case "address":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Address = data
			case "pdu":
				var dt *ModbusPDU
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Pdu = dt
			}
		}
	}
}

func (m *ModbusSerialADU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.modbus.readwrite.ModbusSerialADU"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransactionId, xml.StartElement{Name: xml.Name{Local: "transactionId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Length, xml.StartElement{Name: xml.Name{Local: "length"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Pdu, xml.StartElement{Name: xml.Name{Local: "pdu"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m ModbusSerialADU) String() string {
	return string(m.Box("", 120))
}

func (m ModbusSerialADU) Box(name string, width int) utils.AsciiBox {
	boxName := "ModbusSerialADU"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TransactionId", m.TransactionId, -1))
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint16(0x0000), -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Length", m.Length, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Address", m.Address, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Pdu.Box("pdu", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
