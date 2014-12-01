package pkgbuild

import "testing"

// Test version parsing
func TestVersionParsing(t *testing.T) {
	versions := map[string]bool{
		"1.0beta":   true,
		"1.0.0.0.2": true,
		"a.3_4":     true,
		"A.2":       true,
		"_1.2":      false,
		".2":        false,
		"a.2Ø":      false,
		"1.?":       false,
		"1.-":       false,
	}

	for version, valid := range versions {
		_, err := parseVersion(version)
		if err != nil && valid {
			t.Errorf("Version string '%s' should pass", version)
		}

		if err == nil && !valid {
			t.Errorf("Version string '%s' should not pass", version)
		}
	}
}

// Test parsing array value as value
func TestValueParsing(t *testing.T) {
	input := "pkgdesc=([0]=\"value1\" [1]=\"value2\")\n"

	pkgb, err := parse(input)
	if err != nil {
		t.Error("parse should not fail")
	}

	if pkgb.Pkgdesc != "value1" {
		t.Errorf("should equal 'value1', was: %#v", pkgb.Pkgdesc)
	}
}
