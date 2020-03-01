//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

[discriminatedType 'KNXGroupAddress' [uint 2 'numLevels']
    [typeSwitch 'numLevels'
        ['1' KNXGroupAddressFreeLevel
            [simple uint 16 'subGroup']
        ]
        ['2' KNXGroupAddress2Level
            [simple uint 5  'mainGroup']
            [simple uint 11 'subGroup']
        ]
        ['3' KNXGroupAddress3Level
            [simple uint 5 'mainGroup']
            [simple uint 3 'middleGroup']
            [simple uint 8 'subGroup']
        ]
    ]
]

[discriminatedType 'KnxDatapoint' [uint 10 'mainNumber', uint 10 'subNumber']
    [typeSwitch 'mainNumber','subNumber'
        ['1' KnxDatapointB1
            [reserved uint 7 '0x0']
            [simple   bit    'value']
        ]
        ['2' KnxDatapointB2
            [reserved uint 6 '0x0']
            [simple   bit    'control']
            [simple   bit    'value']
        ]
        ['21' KnxDatapointB8
            [simple   bit    'b7']
            [simple   bit    'b6']
            [simple   bit    'b5']
            [simple   bit    'b4']
            [simple   bit    'b3']
            [simple   bit    'b2']
            [simple   bit    'b1']
            [simple   bit    'b0']
        ]
        ['3' KnxDatapointB1U3
            [reserved uint 4 '0x0']
            [simple   bit    'control']
            [simple   uint 3 'value']
        ]
        ['18' KnxDatapointB1U6
            [simple   bit    'control']
            [reserved uint 1 '0x0']
            [simple   uint 6 'value']
        ]
        ['17' KnxDatapointU6
            [reserved uint 2 '0x0']
            [simple   uint 6 'value']
        ]
        ['5' KnxDatapointU8
            [reserved uint 8 '0x0']
            [simple   uint 8 'value']
        ]
        ['7' KnxDatapointU16
            [reserved uint 8 '0x0']
            [simple uint 16 'value']
        ]
        ['12' KnxDatapointU32
            [reserved uint 8 '0x0']
            [simple uint 32 'value']
        ]
        ['6','20' KnxDatapointB5I3
            [simple   bit   'a']
            [simple   bit   'b']
            [simple   bit   'c']
            [simple   bit   'd']
            [simple   bit   'e']
            [simple   int 8 'value']
        ]
        ['6' KnxDatapointI8
            [reserved uint 8 '0x0']
            [simple   int  8 'value']
        ]
        ['8' KnxDatapointI16
            [reserved uint 8  '0x0']
            [simple   int  16 'value']
        ]
        ['13' KnxDatapointI32
            [reserved uint 8  '0x0']
            [simple   int  32 'value']
        ]
        ['9' KnxDatapointF16
            [reserved uint  8    '0x0']
            [simple   float 4.11 'value']
        ]
        ['14' KnxDatapointF32
            [reserved uint  8    '0x0']
            [simple   float 8.23 'value']
        ]
        ['4' KnxDatapointA8
            [reserved uint 8 '0x0']
            [simple   int  8 'value']
        ]
        ['16' KnxDatapointA112
            [reserved uint 8 '0x0']
            [array    int  8 'value' count '14']
        ]
        ['10' KnxDatapointTime24
            [simple   uint 3 'day']
            [simple   uint 5 'hour']
            [reserved uint 2 '0x0']
            [simple   uint 6 'minutes']
            [reserved uint 2 '0x0']
            [simple   uint 6 'seconds']
        ]
        ['11' KnxDatapointDate24
            [reserved uint 3 '0x0']
            [simple   uint 5 'day']
            [reserved uint 4 '0x0']
            [simple   uint 4 'month']
            [reserved uint 1 '0x0']
            [simple   uint 6 'year']
        ]
        ['19' KnxDatapointDateTime64
            [simple   uint 8 'year']
            [reserved uint 4 '0x0']
            [simple   uint 4 'month']
            [reserved uint 3 '0x0']
            [simple   uint 5 'dayOfMonth']
            [simple   uint 3 'dayOfWeek']
            [simple   uint 5 'hourOfDay']
            [reserved uint 2 '0x0']
            [simple   uint 6 'minutes']
            [reserved uint 2 '0x0']
            [simple   uint 6 'seconds']
            [simple   bit    'fault']
            [simple   bit    'workingDay']
            [simple   bit    'workingDayValid']
            [simple   bit    'yearValid']
            [simple   bit    'dayAndMonthValid']
            [simple   bit    'dayOfWeekValid']
            [simple   bit    'timeValid']
            [simple   bit    'standardSummerTime']
            [simple   bit    'clockQuality']
        ]
        ['15' KnxDatapointDPT_Access_Data
            [simple   uint 4 'D6']
            [simple   uint 4 'D5']
            [simple   uint 4 'D4']
            [simple   uint 4 'D3']
            [simple   uint 4 'D2']
            [simple   uint 4 'D1']
            [simple   bit    'BE']
            [simple   bit    'BP']
            [simple   bit    'BD']
            [simple   bit    'BC']
            [simple   uint 4 'index']
        ]
    ]
]
