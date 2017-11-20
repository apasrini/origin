/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package slice

import (
	"reflect"
	"testing"
)

func TestCopyStrings(t *testing.T) {
	var src1 []string
	dest1 := CopyStrings(src1)

	if !reflect.DeepEqual(src1, dest1) {
		t.Errorf("%v and %v are not equal", src1, dest1)
	}

	src2 := []string{}
	dest2 := CopyStrings(src2)

	if !reflect.DeepEqual(src2, dest2) {
		t.Errorf("%v and %v are not equal", src2, dest2)
	}

	src3 := []string{"a", "c", "b"}
	dest3 := CopyStrings(src3)

	if !reflect.DeepEqual(src3, dest3) {
		t.Errorf("%v and %v are not equal", src3, dest3)
	}

	src3[0] = "A"
	if reflect.DeepEqual(src3, dest3) {
		t.Errorf("CopyStrings didn't make a copy")
	}
}

func TestSortStrings(t *testing.T) {
	src := []string{"a", "c", "b"}
	dest := SortStrings(src)
	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("SortString didn't sort the strings")
	}

	if !reflect.DeepEqual(src, expected) {
		t.Errorf("SortString didn't sort in place")
	}
}

func TestShuffleStrings(t *testing.T) {
	var src []string
	dest := ShuffleStrings(src)

	if dest != nil {
		t.Errorf("ShuffleStrings for a nil slice got a non-nil slice")
	}

	src = []string{"a", "b", "c", "d", "e", "f"}
	dest = ShuffleStrings(src)

	if len(src) != len(dest) {
		t.Errorf("Shuffled slice is wrong length, expected %v got %v", len(src), len(dest))
	}

	m := make(map[string]bool, len(dest))
	for _, s := range dest {
		m[s] = true
	}

	for _, k := range src {
		if _, exists := m[k]; !exists {
			t.Errorf("Element %v missing from shuffled slice", k)
		}
	}
}
