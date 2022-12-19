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
package org.apache.plc4x.java.df1.readwrite;

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

public abstract class DF1Command implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getCommandCode();

  // Properties.
  protected final short status;
  protected final int transactionCounter;

  public DF1Command(short status, int transactionCounter) {
    super();
    this.status = status;
    this.transactionCounter = transactionCounter;
  }

  public short getStatus() {
    return status;
  }

  public int getTransactionCounter() {
    return transactionCounter;
  }

  protected abstract void serializeDF1CommandChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DF1Command");

    // Discriminator Field (commandCode) (Used as input to a switch field)
    writeDiscriminatorField("commandCode", getCommandCode(), writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (transactionCounter)
    writeSimpleField("transactionCounter", transactionCounter, writeUnsignedInt(writeBuffer, 16));

    // Switch field (Serialize the sub-type)
    serializeDF1CommandChild(writeBuffer);

    writeBuffer.popContext("DF1Command");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DF1Command _value = this;

    // Discriminator Field (commandCode)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Simple field (transactionCounter)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static DF1Command staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DF1Command staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DF1Command");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short commandCode = readDiscriminatorField("commandCode", readUnsignedShort(readBuffer, 8));

    short status = readSimpleField("status", readUnsignedShort(readBuffer, 8));

    int transactionCounter = readSimpleField("transactionCounter", readUnsignedInt(readBuffer, 16));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    DF1CommandBuilder builder = null;
    if (EvaluationHelper.equals(commandCode, (short) 0x01)) {
      builder = DF1UnprotectedReadRequest.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandCode, (short) 0x41)) {
      builder = DF1UnprotectedReadResponse.staticParseBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandCode="
              + commandCode
              + "]");
    }

    readBuffer.closeContext("DF1Command");
    // Create the instance
    DF1Command _dF1Command = builder.build(status, transactionCounter);
    return _dF1Command;
  }

  public static interface DF1CommandBuilder {
    DF1Command build(short status, int transactionCounter);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DF1Command)) {
      return false;
    }
    DF1Command that = (DF1Command) o;
    return (getStatus() == that.getStatus())
        && (getTransactionCounter() == that.getTransactionCounter())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getStatus(), getTransactionCounter());
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