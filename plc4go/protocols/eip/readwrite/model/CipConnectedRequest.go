/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CipConnectedRequest is the corresponding interface of CipConnectedRequest
type CipConnectedRequest interface {
	utils.LengthAware
	utils.Serializable
	CipService
	// GetPathSegments returns PathSegments (property field)
	GetPathSegments() []byte
}

// CipConnectedRequestExactly can be used when we want exactly this type and not a type which fulfills CipConnectedRequest.
// This is useful for switch cases.
type CipConnectedRequestExactly interface {
	CipConnectedRequest
	isCipConnectedRequest() bool
}

// _CipConnectedRequest is the data-structure of this message
type _CipConnectedRequest struct {
	*_CipService
	PathSegments []byte
	// Reserved Fields
	reservedField0 *uint16
	reservedField1 *uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectedRequest) GetService() uint8 {
	return 0x52
}

func (m *_CipConnectedRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectedRequest) GetConnected() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectedRequest) InitializeParent(parent CipService) {}

func (m *_CipConnectedRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectedRequest) GetPathSegments() []byte {
	return m.PathSegments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectedRequest factory function for _CipConnectedRequest
func NewCipConnectedRequest(pathSegments []byte, serviceLen uint16) *_CipConnectedRequest {
	_result := &_CipConnectedRequest{
		PathSegments: pathSegments,
		_CipService:  NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectedRequest(structType interface{}) CipConnectedRequest {
	if casted, ok := structType.(CipConnectedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectedRequest) GetTypeName() string {
	return "CipConnectedRequest"
}

func (m *_CipConnectedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.PathSegments) > 0 {
		lengthInBits += 8 * uint16(len(m.PathSegments))
	}

	// Reserved Field (reserved)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CipConnectedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectedRequestParse(theBytes []byte, connected bool, serviceLen uint16) (CipConnectedRequest, error) {
	return CipConnectedRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectedRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	requestPathSize, _requestPathSizeErr := readBuffer.ReadUint8("requestPathSize", 8)
	_ = requestPathSize
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field of CipConnectedRequest")
	}
	// Byte Array field (pathSegments)
	numberOfBytespathSegments := int(uint16(requestPathSize) * uint16(uint16(2)))
	pathSegments, _readArrayErr := readBuffer.ReadByteArray("pathSegments", numberOfBytespathSegments)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'pathSegments' field of CipConnectedRequest")
	}

	var reservedField0 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectedRequest")
		}
		if reserved != uint16(0x0001) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0001),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *uint32
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint32("reserved", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectedRequest")
		}
		if reserved != uint32(0x00000000) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint32(0x00000000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("CipConnectedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectedRequest")
	}

	// Create a partially initialized instance
	_child := &_CipConnectedRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		PathSegments:   pathSegments,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectedRequest")
		}

		// Implicit Field (requestPathSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		requestPathSize := uint8(uint8(uint8(len(m.GetPathSegments()))) / uint8(uint8(2)))
		_requestPathSizeErr := writeBuffer.WriteUint8("requestPathSize", 8, (requestPathSize))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Array Field (pathSegments)
		// Byte Array field (pathSegments)
		if err := writeBuffer.WriteByteArray("pathSegments", m.GetPathSegments()); err != nil {
			return errors.Wrap(err, "Error serializing 'pathSegments' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint16 = uint16(0x0001)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint16(0x0001),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint16("reserved", 16, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint32 = uint32(0x00000000)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint32(0x00000000),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint32("reserved", 32, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("CipConnectedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectedRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectedRequest) isCipConnectedRequest() bool {
	return true
}

func (m *_CipConnectedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}