// Code generated by ragel DO NOT EDIT.

//.... lightning/mydump/parser.rl:1
// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Please edit `parser.rl` if you want to modify this file. To generate
// `parser_generated.go`, please execute
//
// ```sh
// make data_parsers
// ```

package mydump

import (
	"io"

	"github.com/pingcap/errors"
	"github.com/pingcap/tidb-lightning/lightning/common"
)

//.... lightning/mydump/parser.rl:92

//.... tmp_parser.go:38
const chunk_parser_start int = 27
const chunk_parser_first_final int = 27
const chunk_parser_error int = 0

const chunk_parser_en_main int = 27

//.... lightning/mydump/parser.rl:95

func (parser *ChunkParser) lex() (token, []byte, error) {
	var cs, ts, te, act, p int

	//.... tmp_parser.go:51
	{
		cs = chunk_parser_start
		ts = 0
		te = 0
		act = 0
	}

	//.... lightning/mydump/parser.rl:99

	for {
		data := parser.buf
		consumedToken := tokNil
		pe := len(data)
		eof := -1
		if parser.isLastChunk {
			eof = pe
		}

		//.... tmp_parser.go:71
		{
			if p == pe {
				goto _test_eof
			}
			switch cs {
			case 27:
				goto st_case_27
			case 28:
				goto st_case_28
			case 1:
				goto st_case_1
			case 2:
				goto st_case_2
			case 3:
				goto st_case_3
			case 0:
				goto st_case_0
			case 4:
				goto st_case_4
			case 5:
				goto st_case_5
			case 6:
				goto st_case_6
			case 7:
				goto st_case_7
			case 8:
				goto st_case_8
			case 9:
				goto st_case_9
			case 10:
				goto st_case_10
			case 11:
				goto st_case_11
			case 12:
				goto st_case_12
			case 13:
				goto st_case_13
			case 14:
				goto st_case_14
			case 15:
				goto st_case_15
			case 29:
				goto st_case_29
			case 30:
				goto st_case_30
			case 16:
				goto st_case_16
			case 17:
				goto st_case_17
			case 31:
				goto st_case_31
			case 18:
				goto st_case_18
			case 19:
				goto st_case_19
			case 32:
				goto st_case_32
			case 33:
				goto st_case_33
			case 34:
				goto st_case_34
			case 20:
				goto st_case_20
			case 21:
				goto st_case_21
			case 22:
				goto st_case_22
			case 23:
				goto st_case_23
			case 24:
				goto st_case_24
			case 35:
				goto st_case_35
			case 25:
				goto st_case_25
			case 26:
				goto st_case_26
			case 36:
				goto st_case_36
			case 37:
				goto st_case_37
			case 38:
				goto st_case_38
			case 39:
				goto st_case_39
			case 40:
				goto st_case_40
			case 41:
				goto st_case_41
			case 42:
				goto st_case_42
			case 43:
				goto st_case_43
			case 44:
				goto st_case_44
			case 45:
				goto st_case_45
			case 46:
				goto st_case_46
			}
			goto st_out
		tr0:
			//.... NONE:1
			switch act {
			case 0:
				{
					{
						goto st0
					}
				}
			case 2:
				{
					p = (te) - 1

					consumedToken = tokValues
					{
						p++
						cs = 27
						goto _out
					}
				}
			case 4:
				{
					p = (te) - 1

					consumedToken = tokName
					{
						p++
						cs = 27
						goto _out
					}
				}
			default:
				{
					p = (te) - 1
				}
			}

			goto st27
		tr9:
			//.... lightning/mydump/parser.rl:80
			te = p + 1
			{
				consumedToken = tokRow
				{
					p++
					cs = 27
					goto _out
				}
			}
			goto st27
		tr19:
			//.... lightning/mydump/parser.rl:85
			p = (te) - 1
			{
				consumedToken = tokName
				{
					p++
					cs = 27
					goto _out
				}
			}
			goto st27
		tr21:
			//.... lightning/mydump/parser.rl:73
			te = p + 1

			goto st27
		tr40:
			//.... lightning/mydump/parser.rl:85
			te = p
			p--
			{
				consumedToken = tokName
				{
					p++
					cs = 27
					goto _out
				}
			}
			goto st27
		tr41:
			//.... lightning/mydump/parser.rl:73
			te = p
			p--

			goto st27
		st27:
			//.... NONE:1
			ts = 0

			//.... NONE:1
			act = 0

			if p++; p == pe {
				goto _test_eof27
			}
		st_case_27:
			//.... NONE:1
			ts = p

			//.... tmp_parser.go:246
			switch data[p] {
			case 32:
				goto tr21
			case 34:
				goto st1
			case 40:
				goto st4
			case 44:
				goto tr21
			case 45:
				goto tr36
			case 47:
				goto tr37
			case 59:
				goto tr21
			case 73:
				goto tr38
			case 86:
				goto tr39
			case 96:
				goto st3
			case 105:
				goto tr38
			case 118:
				goto tr39
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto st0
				}
			case data[p] >= 9:
				goto tr21
			}
			goto tr2
		tr2:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st28
		tr43:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:73
			act = 1
			goto st28
		tr53:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:75
			act = 2
			goto st28
		st28:
			if p++; p == pe {
				goto _test_eof28
			}
		st_case_28:
			//.... tmp_parser.go:308
			switch data[p] {
			case 32:
				goto tr0
			case 34:
				goto st1
			case 44:
				goto tr0
			case 59:
				goto tr0
			case 96:
				goto st3
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr0
				}
			case data[p] >= 9:
				goto tr0
			}
			goto tr2
		st1:
			if p++; p == pe {
				goto _test_eof1
			}
		st_case_1:
			switch data[p] {
			case 34:
				goto tr2
			case 92:
				goto st2
			}
			goto st1
		st2:
			if p++; p == pe {
				goto _test_eof2
			}
		st_case_2:
			goto st1
		st3:
			if p++; p == pe {
				goto _test_eof3
			}
		st_case_3:
			if data[p] == 96 {
				goto tr2
			}
			goto st3
		st_case_0:
		st0:
			cs = 0
			goto _out
		st4:
			if p++; p == pe {
				goto _test_eof4
			}
		st_case_4:
			switch data[p] {
			case 34:
				goto st5
			case 39:
				goto st7
			case 40:
				goto st9
			case 41:
				goto tr9
			case 96:
				goto st15
			}
			goto st4
		st5:
			if p++; p == pe {
				goto _test_eof5
			}
		st_case_5:
			switch data[p] {
			case 34:
				goto st4
			case 92:
				goto st6
			}
			goto st5
		st6:
			if p++; p == pe {
				goto _test_eof6
			}
		st_case_6:
			goto st5
		st7:
			if p++; p == pe {
				goto _test_eof7
			}
		st_case_7:
			switch data[p] {
			case 39:
				goto st4
			case 92:
				goto st8
			}
			goto st7
		st8:
			if p++; p == pe {
				goto _test_eof8
			}
		st_case_8:
			goto st7
		st9:
			if p++; p == pe {
				goto _test_eof9
			}
		st_case_9:
			switch data[p] {
			case 34:
				goto st10
			case 39:
				goto st12
			case 40:
				goto st0
			case 41:
				goto st4
			case 96:
				goto st14
			}
			goto st9
		st10:
			if p++; p == pe {
				goto _test_eof10
			}
		st_case_10:
			switch data[p] {
			case 34:
				goto st9
			case 92:
				goto st11
			}
			goto st10
		st11:
			if p++; p == pe {
				goto _test_eof11
			}
		st_case_11:
			goto st10
		st12:
			if p++; p == pe {
				goto _test_eof12
			}
		st_case_12:
			switch data[p] {
			case 39:
				goto st9
			case 92:
				goto st13
			}
			goto st12
		st13:
			if p++; p == pe {
				goto _test_eof13
			}
		st_case_13:
			goto st12
		st14:
			if p++; p == pe {
				goto _test_eof14
			}
		st_case_14:
			if data[p] == 96 {
				goto st9
			}
			goto st14
		st15:
			if p++; p == pe {
				goto _test_eof15
			}
		st_case_15:
			if data[p] == 96 {
				goto st4
			}
			goto st15
		tr36:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st29
		st29:
			if p++; p == pe {
				goto _test_eof29
			}
		st_case_29:
			//.... tmp_parser.go:499
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 45:
				goto tr24
			case 59:
				goto tr40
			case 96:
				goto st3
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr24:
			//.... NONE:1
			te = p + 1

			goto st30
		st30:
			if p++; p == pe {
				goto _test_eof30
			}
		st_case_30:
			//.... tmp_parser.go:533
			switch data[p] {
			case 10:
				goto tr21
			case 32:
				goto st16
			case 34:
				goto st17
			case 44:
				goto st16
			case 59:
				goto st16
			case 96:
				goto st19
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto st16
				}
			case data[p] >= 9:
				goto st16
			}
			goto tr24
		st16:
			if p++; p == pe {
				goto _test_eof16
			}
		st_case_16:
			if data[p] == 10 {
				goto tr21
			}
			goto st16
		st17:
			if p++; p == pe {
				goto _test_eof17
			}
		st_case_17:
			switch data[p] {
			case 10:
				goto tr23
			case 34:
				goto tr24
			case 92:
				goto st18
			}
			goto st17
		tr23:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:73
			act = 1
			goto st31
		st31:
			if p++; p == pe {
				goto _test_eof31
			}
		st_case_31:
			//.... tmp_parser.go:592
			switch data[p] {
			case 34:
				goto tr2
			case 92:
				goto st2
			}
			goto st1
		st18:
			if p++; p == pe {
				goto _test_eof18
			}
		st_case_18:
			if data[p] == 10 {
				goto tr23
			}
			goto st17
		st19:
			if p++; p == pe {
				goto _test_eof19
			}
		st_case_19:
			switch data[p] {
			case 10:
				goto tr27
			case 96:
				goto tr24
			}
			goto st19
		tr27:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:73
			act = 1
			goto st32
		st32:
			if p++; p == pe {
				goto _test_eof32
			}
		st_case_32:
			//.... tmp_parser.go:633
			if data[p] == 96 {
				goto tr2
			}
			goto st3
		tr37:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st33
		st33:
			if p++; p == pe {
				goto _test_eof33
			}
		st_case_33:
			//.... tmp_parser.go:650
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 42:
				goto tr31
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 96:
				goto st3
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr31:
			//.... NONE:1
			te = p + 1

			goto st34
		st34:
			if p++; p == pe {
				goto _test_eof34
			}
		st_case_34:
			//.... tmp_parser.go:684
			switch data[p] {
			case 32:
				goto st20
			case 34:
				goto st22
			case 42:
				goto tr42
			case 44:
				goto st20
			case 59:
				goto st20
			case 96:
				goto st25
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto st20
				}
			case data[p] >= 9:
				goto st20
			}
			goto tr31
		st20:
			if p++; p == pe {
				goto _test_eof20
			}
		st_case_20:
			if data[p] == 42 {
				goto st21
			}
			goto st20
		st21:
			if p++; p == pe {
				goto _test_eof21
			}
		st_case_21:
			switch data[p] {
			case 42:
				goto st21
			case 47:
				goto tr21
			}
			goto st20
		st22:
			if p++; p == pe {
				goto _test_eof22
			}
		st_case_22:
			switch data[p] {
			case 34:
				goto tr31
			case 42:
				goto st23
			case 92:
				goto st24
			}
			goto st22
		st23:
			if p++; p == pe {
				goto _test_eof23
			}
		st_case_23:
			switch data[p] {
			case 34:
				goto tr31
			case 42:
				goto st23
			case 47:
				goto tr23
			case 92:
				goto st24
			}
			goto st22
		st24:
			if p++; p == pe {
				goto _test_eof24
			}
		st_case_24:
			if data[p] == 42 {
				goto st23
			}
			goto st22
		tr42:
			//.... NONE:1
			te = p + 1

			goto st35
		st35:
			if p++; p == pe {
				goto _test_eof35
			}
		st_case_35:
			//.... tmp_parser.go:778
			switch data[p] {
			case 32:
				goto st20
			case 34:
				goto st22
			case 42:
				goto tr42
			case 44:
				goto st20
			case 47:
				goto tr43
			case 59:
				goto st20
			case 96:
				goto st25
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto st20
				}
			case data[p] >= 9:
				goto st20
			}
			goto tr31
		st25:
			if p++; p == pe {
				goto _test_eof25
			}
		st_case_25:
			switch data[p] {
			case 42:
				goto st26
			case 96:
				goto tr31
			}
			goto st25
		st26:
			if p++; p == pe {
				goto _test_eof26
			}
		st_case_26:
			switch data[p] {
			case 42:
				goto st26
			case 47:
				goto tr27
			case 96:
				goto tr31
			}
			goto st25
		tr38:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st36
		st36:
			if p++; p == pe {
				goto _test_eof36
			}
		st_case_36:
			//.... tmp_parser.go:842
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 78:
				goto tr44
			case 96:
				goto st3
			case 110:
				goto tr44
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr44:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st37
		st37:
			if p++; p == pe {
				goto _test_eof37
			}
		st_case_37:
			//.... tmp_parser.go:880
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 83:
				goto tr45
			case 84:
				goto tr46
			case 96:
				goto st3
			case 115:
				goto tr45
			case 116:
				goto tr46
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr45:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st38
		st38:
			if p++; p == pe {
				goto _test_eof38
			}
		st_case_38:
			//.... tmp_parser.go:922
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 69:
				goto tr47
			case 96:
				goto st3
			case 101:
				goto tr47
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr47:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st39
		st39:
			if p++; p == pe {
				goto _test_eof39
			}
		st_case_39:
			//.... tmp_parser.go:960
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 82:
				goto tr48
			case 96:
				goto st3
			case 114:
				goto tr48
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr48:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st40
		st40:
			if p++; p == pe {
				goto _test_eof40
			}
		st_case_40:
			//.... tmp_parser.go:998
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 84:
				goto tr43
			case 96:
				goto st3
			case 116:
				goto tr43
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr46:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st41
		st41:
			if p++; p == pe {
				goto _test_eof41
			}
		st_case_41:
			//.... tmp_parser.go:1036
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 79:
				goto tr43
			case 96:
				goto st3
			case 111:
				goto tr43
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr39:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st42
		st42:
			if p++; p == pe {
				goto _test_eof42
			}
		st_case_42:
			//.... tmp_parser.go:1074
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 65:
				goto tr49
			case 96:
				goto st3
			case 97:
				goto tr49
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr49:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st43
		st43:
			if p++; p == pe {
				goto _test_eof43
			}
		st_case_43:
			//.... tmp_parser.go:1112
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 76:
				goto tr50
			case 96:
				goto st3
			case 108:
				goto tr50
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr50:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st44
		st44:
			if p++; p == pe {
				goto _test_eof44
			}
		st_case_44:
			//.... tmp_parser.go:1150
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 85:
				goto tr51
			case 96:
				goto st3
			case 117:
				goto tr51
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr51:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st45
		st45:
			if p++; p == pe {
				goto _test_eof45
			}
		st_case_45:
			//.... tmp_parser.go:1188
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 69:
				goto tr52
			case 96:
				goto st3
			case 101:
				goto tr52
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		tr52:
			//.... NONE:1
			te = p + 1

			//.... lightning/mydump/parser.rl:85
			act = 4
			goto st46
		st46:
			if p++; p == pe {
				goto _test_eof46
			}
		st_case_46:
			//.... tmp_parser.go:1226
			switch data[p] {
			case 32:
				goto tr40
			case 34:
				goto st1
			case 44:
				goto tr40
			case 59:
				goto tr40
			case 83:
				goto tr53
			case 96:
				goto st3
			case 115:
				goto tr53
			}
			switch {
			case data[p] > 13:
				if 39 <= data[p] && data[p] <= 41 {
					goto tr40
				}
			case data[p] >= 9:
				goto tr40
			}
			goto tr2
		st_out:
		_test_eof27:
			cs = 27
			goto _test_eof
		_test_eof28:
			cs = 28
			goto _test_eof
		_test_eof1:
			cs = 1
			goto _test_eof
		_test_eof2:
			cs = 2
			goto _test_eof
		_test_eof3:
			cs = 3
			goto _test_eof
		_test_eof4:
			cs = 4
			goto _test_eof
		_test_eof5:
			cs = 5
			goto _test_eof
		_test_eof6:
			cs = 6
			goto _test_eof
		_test_eof7:
			cs = 7
			goto _test_eof
		_test_eof8:
			cs = 8
			goto _test_eof
		_test_eof9:
			cs = 9
			goto _test_eof
		_test_eof10:
			cs = 10
			goto _test_eof
		_test_eof11:
			cs = 11
			goto _test_eof
		_test_eof12:
			cs = 12
			goto _test_eof
		_test_eof13:
			cs = 13
			goto _test_eof
		_test_eof14:
			cs = 14
			goto _test_eof
		_test_eof15:
			cs = 15
			goto _test_eof
		_test_eof29:
			cs = 29
			goto _test_eof
		_test_eof30:
			cs = 30
			goto _test_eof
		_test_eof16:
			cs = 16
			goto _test_eof
		_test_eof17:
			cs = 17
			goto _test_eof
		_test_eof31:
			cs = 31
			goto _test_eof
		_test_eof18:
			cs = 18
			goto _test_eof
		_test_eof19:
			cs = 19
			goto _test_eof
		_test_eof32:
			cs = 32
			goto _test_eof
		_test_eof33:
			cs = 33
			goto _test_eof
		_test_eof34:
			cs = 34
			goto _test_eof
		_test_eof20:
			cs = 20
			goto _test_eof
		_test_eof21:
			cs = 21
			goto _test_eof
		_test_eof22:
			cs = 22
			goto _test_eof
		_test_eof23:
			cs = 23
			goto _test_eof
		_test_eof24:
			cs = 24
			goto _test_eof
		_test_eof35:
			cs = 35
			goto _test_eof
		_test_eof25:
			cs = 25
			goto _test_eof
		_test_eof26:
			cs = 26
			goto _test_eof
		_test_eof36:
			cs = 36
			goto _test_eof
		_test_eof37:
			cs = 37
			goto _test_eof
		_test_eof38:
			cs = 38
			goto _test_eof
		_test_eof39:
			cs = 39
			goto _test_eof
		_test_eof40:
			cs = 40
			goto _test_eof
		_test_eof41:
			cs = 41
			goto _test_eof
		_test_eof42:
			cs = 42
			goto _test_eof
		_test_eof43:
			cs = 43
			goto _test_eof
		_test_eof44:
			cs = 44
			goto _test_eof
		_test_eof45:
			cs = 45
			goto _test_eof
		_test_eof46:
			cs = 46
			goto _test_eof

		_test_eof:
			{
			}
			if p == eof {
				switch cs {
				case 28:
					goto tr0
				case 1:
					goto tr0
				case 2:
					goto tr0
				case 3:
					goto tr0
				case 29:
					goto tr40
				case 30:
					goto tr40
				case 16:
					goto tr19
				case 17:
					goto tr19
				case 31:
					goto tr41
				case 18:
					goto tr19
				case 19:
					goto tr19
				case 32:
					goto tr41
				case 33:
					goto tr40
				case 34:
					goto tr40
				case 20:
					goto tr19
				case 21:
					goto tr19
				case 22:
					goto tr19
				case 23:
					goto tr19
				case 24:
					goto tr19
				case 35:
					goto tr40
				case 25:
					goto tr19
				case 26:
					goto tr19
				case 36:
					goto tr40
				case 37:
					goto tr40
				case 38:
					goto tr40
				case 39:
					goto tr40
				case 40:
					goto tr40
				case 41:
					goto tr40
				case 42:
					goto tr40
				case 43:
					goto tr40
				case 44:
					goto tr40
				case 45:
					goto tr40
				case 46:
					goto tr40
				}
			}

		_out:
			{
			}
		}

		//.... lightning/mydump/parser.rl:110

		if cs == 0 {
			common.AppLogger.Errorf("Syntax error near byte %d, content is «%s»", parser.pos, string(data))
			return tokNil, nil, errors.New("Syntax error")
		}

		if consumedToken != tokNil {
			result := data[ts:te]
			parser.buf = data[te:]
			parser.pos += int64(te)
			return consumedToken, result, nil
		}

		if parser.isLastChunk {
			return tokNil, nil, io.EOF
		}

		parser.buf = parser.buf[ts:]
		parser.pos += int64(ts)
		p -= ts
		te -= ts
		ts = 0
		if err := parser.readBlock(); err != nil {
			return tokNil, nil, errors.Trace(err)
		}
	}

	return tokNil, nil, nil
}
