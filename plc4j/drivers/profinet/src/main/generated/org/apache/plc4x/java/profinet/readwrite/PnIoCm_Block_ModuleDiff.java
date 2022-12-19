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
package org.apache.plc4x.java.profinet.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class PnIoCm_Block_ModuleDiff extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.MODULE_DIFF_BLOCK;
  }

  // Properties.
  protected final List<PnIoCm_ModuleDiffBlockApi> apis;

  public PnIoCm_Block_ModuleDiff(
      short blockVersionHigh, short blockVersionLow, List<PnIoCm_ModuleDiffBlockApi> apis) {
    super(blockVersionHigh, blockVersionLow);
    this.apis = apis;
  }

  public List<PnIoCm_ModuleDiffBlockApi> getApis() {
    return apis;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Block_ModuleDiff");

    // Implicit Field (numberOfApis) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int numberOfApis = (int) (COUNT(getApis()));
    writeImplicitField("numberOfApis", numberOfApis, writeUnsignedInt(writeBuffer, 16));

    // Array Field (apis)
    writeComplexTypeArrayField("apis", apis, writeBuffer);

    writeBuffer.popContext("PnIoCm_Block_ModuleDiff");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_ModuleDiff _value = this;

    // Implicit Field (numberOfApis)
    lengthInBits += 16;

    // Array field
    if (apis != null) {
      int i = 0;
      for (PnIoCm_ModuleDiffBlockApi element : apis) {
        boolean last = ++i >= apis.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static PnIoCm_Block_ModuleDiffBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_ModuleDiff");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int numberOfApis =
        readImplicitField(
            "numberOfApis",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<PnIoCm_ModuleDiffBlockApi> apis =
        readCountArrayField(
            "apis",
            new DataReaderComplexDefault<>(
                () -> PnIoCm_ModuleDiffBlockApi.staticParse(readBuffer), readBuffer),
            numberOfApis,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_ModuleDiff");
    // Create the instance
    return new PnIoCm_Block_ModuleDiffBuilder(apis);
  }

  public static class PnIoCm_Block_ModuleDiffBuilder implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final List<PnIoCm_ModuleDiffBlockApi> apis;

    public PnIoCm_Block_ModuleDiffBuilder(List<PnIoCm_ModuleDiffBlockApi> apis) {

      this.apis = apis;
    }

    public PnIoCm_Block_ModuleDiff build(short blockVersionHigh, short blockVersionLow) {
      PnIoCm_Block_ModuleDiff pnIoCm_Block_ModuleDiff =
          new PnIoCm_Block_ModuleDiff(blockVersionHigh, blockVersionLow, apis);
      return pnIoCm_Block_ModuleDiff;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_ModuleDiff)) {
      return false;
    }
    PnIoCm_Block_ModuleDiff that = (PnIoCm_Block_ModuleDiff) o;
    return (getApis() == that.getApis()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getApis());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}