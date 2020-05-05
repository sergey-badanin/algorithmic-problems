package groupanagrams

import (
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	given := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	want := [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}
	if got := groupAnagrams(given); !equals(got, want) {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = []string{"eat", "tea"}
	want = [][]string{{"eat", "tea"}}
	if got := groupAnagrams(given); !equals(got, want) {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = []string{"ate", "nat", "bat"}
	want = [][]string{{"ate"}, {"nat"}, {"bat"}}
	if got := groupAnagrams(given); !equals(got, want) {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = []string{"ate"}
	want = [][]string{{"ate"}}
	if got := groupAnagrams(given); !equals(got, want) {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}
}

func TestSortBytes(t *testing.T) {
	given := "cba"
	want := "abc"
	if got := sortBytes(given); want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}
}

func equals(a, b [][]string) bool {
	aStrings := convertToMap(a)
	bStrings := convertToMap(b)

	if len(aStrings) != len(bStrings) {
		return false
	}

	for k := range aStrings {
		if !bStrings[k] {
			return false
		}
	}
	return true
}

func convertToMap(a [][]string) map[string]bool {
	aStrings := make(map[string]bool)
	for _, v := range a {
		key := joinElements(v)
		aStrings[key] = true
	}
	return aStrings
}

func joinElements(a []string) string {
	sort.Strings(a)
	return strings.Join(a, "")
}
