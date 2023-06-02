package nfcnorm

import (
	"strings"
	"testing"
)

const (
	nfdString = "야이놈시키야â"
	nfcString = "야이놈시키야â"
)

func TestNFCNorm(t *testing.T) {
	if strings.Compare(nfdString, nfcString) == 0 {
		t.Errorf("NFD string '%s' and NFC string '%s' should not be equal.", nfdString, nfcString)
	}

	if Length(nfdString) == Length(nfcString) {
		t.Errorf("The lengths of NFD string '%s' and NFC string '%s' should not be equal.", nfdString, nfcString)
	}

	if !Normalizable(nfdString) {
		t.Errorf("NFD string '%s' should be normalizable.", nfdString)
	}

	normalized := Normalize(nfdString)
	if strings.Compare(normalized, nfcString) != 0 {
		t.Errorf("Normalized string '%s' and NFC string '%s' should be equal.", normalized, nfcString)
	}

	if Normalizable(normalized) {
		t.Errorf("Normalized string '%s' should not be normalizable.", normalized)
	}

	if Length(normalized) != Length(nfcString) {
		t.Errorf("The lengths of normalized string '%s' and NFC string '%s' should be equal.", normalized, nfcString)
	}
}
