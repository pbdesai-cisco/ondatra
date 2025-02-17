// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixweb

import (
	"encoding/json"
	"testing"
)

func TestOpArgs(t *testing.T) {
	const want1, want2 = "a1", "a2"
	input := OpArgs{want1, want2}
	bytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("OpArgs got error marshaling: %v", err)
	}
	output := struct {
		Arg1 string `json:"arg1"`
		Arg2 string `json:"arg2"`
	}{}
	if err := json.Unmarshal(bytes, &output); err != nil {
		t.Fatalf("OpArgs got error unmarshaling: %v", err)
	}
	if got1, got2 := output.Arg1, output.Arg2; got1 != want1 || got2 != want2 {
		t.Errorf("OpArgs unmarshal got [%s, %s], want [%s, %s]", got1, got2, want1, want2)
	}
}
