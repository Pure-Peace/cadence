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

package interpreter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/runtime/interpreter"
)

func TestInterpretInterfaceDefaultImplementation(t *testing.T) {

	t.Parallel()

	t.Run("interface", func(t *testing.T) {

		t.Parallel()

		inter := parseCheckAndInterpret(t, `

          struct interface IA {
              fun test(): Int {
                  return 42
              }
          }

          struct Test: IA {

          }

          fun main(): Int {
              return Test().test()
          }
        `)

		value, err := inter.Invoke("main")
		require.NoError(t, err)

		assert.Equal(t,
			interpreter.NewIntValueFromInt64(42),
			value,
		)
	})

	t.Run("type requirement", func(t *testing.T) {

		t.Parallel()

		inter, err := parseCheckAndInterpretWithOptions(t, `

              contract interface IA {

                  struct X {
                      fun test(): Int {
                          return 42
                      }
                  }
              }

              contract Test: IA {
                  struct X {
                  }
              }

              fun main(): Int {
                  return Test.X().test()
              }
            `,
			ParseCheckAndInterpretOptions{
				Options: []interpreter.Option{
					makeContractValueHandler(nil, nil, nil),
				},
			},
		)
		require.NoError(t, err)

		value, err := inter.Invoke("main")
		require.NoError(t, err)

		assert.Equal(t,
			interpreter.NewIntValueFromInt64(42),
			value,
		)
	})
}

func TestInterpretInterfaceDefaultImplementationWhenOverriden(t *testing.T) {

	t.Parallel()

	t.Run("interface", func(t *testing.T) {

		t.Parallel()

		inter := parseCheckAndInterpret(t, `

          struct interface IA {
              fun test(): Int {
                  return 41
              }
          }

          struct Test: IA {
              fun test(): Int {
                  return 42
              }
          }

          fun main(): Int {
              return Test().test()
          }
        `)

		value, err := inter.Invoke("main")
		require.NoError(t, err)

		assert.Equal(t,
			interpreter.NewIntValueFromInt64(42),
			value,
		)
	})

	t.Run("type requirement", func(t *testing.T) {

		t.Parallel()

		inter, err := parseCheckAndInterpretWithOptions(t,
			`
              contract interface IA {

                  struct X {
                      fun test(): Int {
                          return 41
                     }
                  }
              }

              contract Test: IA {

                  struct X {
                      fun test(): Int {
                          return 42
                      }
                  }
              }

              fun main(): Int {
                  return Test.X().test()
              }
            `,
			ParseCheckAndInterpretOptions{
				Options: []interpreter.Option{
					makeContractValueHandler(nil, nil, nil),
				},
			},
		)

		require.NoError(t, err)

		value, err := inter.Invoke("main")
		require.NoError(t, err)

		assert.Equal(t,
			interpreter.NewIntValueFromInt64(42),
			value,
		)
	})

}