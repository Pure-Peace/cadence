/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stdlib

import (
	"fmt"

	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
)

// This file defines functions built in to the Flow runtime.

const rlpDecodeStringFunctionDocString = `
 accepts an RLP encoded byte array and decodes it into an string.
 `

var rlpDecodeStringFunctionType = &sema.FunctionType{
	Parameters: []*sema.Parameter{
		{
			Identifier: "input",
			TypeAnnotation: sema.NewTypeAnnotation(
				sema.ByteArrayType,
			),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		sema.ByteArrayType,
	),
}

const rlpDecodeListFunctionDocString = `
 accepts an RLP encoded byte array and decodes it into an array of encoded elements.
 `

var rlpDecodeListFunctionType = &sema.FunctionType{
	Parameters: []*sema.Parameter{
		{
			Identifier: "input",
			TypeAnnotation: sema.NewTypeAnnotation(
				sema.ByteArrayType,
			),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.VariableSizedType{
			Type: sema.ByteArrayType,
		},
	),
}

// RLPBuiltinImpls defines the set of functions needed to implement the RLP
// built-in functions.
type RLPBuiltinImpls struct {
	RLPDecodeString interpreter.HostFunction
	RLPDecodeList   interpreter.HostFunction
}

// RLPBuiltInFunctions returns a list of standard library functions, bound to
// the provided implementation.
func RLPBuiltInFunctions(impls RLPBuiltinImpls) StandardLibraryFunctions {
	return StandardLibraryFunctions{
		NewStandardLibraryFunction(
			"RLPDecodeString",
			rlpDecodeStringFunctionType,
			rlpDecodeStringFunctionDocString,
			impls.RLPDecodeString,
		),
		NewStandardLibraryFunction(
			"RLPDecodeList",
			rlpDecodeListFunctionType,
			rlpDecodeListFunctionDocString,
			impls.RLPDecodeList,
		),
	}
}

func DefaultRLPBuiltinImpls() RLPBuiltinImpls {
	return RLPBuiltinImpls{
		RLPDecodeString: func(invocation interpreter.Invocation) interpreter.Value {
			// input := invocation.Arguments[0].(*interpreter.ArrayValue)
			// output, err := rlp.DecodeString(input, 0)
			// if err != nil {
			// 	panic(err)
			// }
			// return output
			panic(fmt.Errorf("cannot rlp decode string"))
		},
		RLPDecodeList: func(invocation interpreter.Invocation) interpreter.Value {
			panic(fmt.Errorf("cannot rlp decode list"))
		},
	}
}