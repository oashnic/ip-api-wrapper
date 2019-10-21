package ipapi

import (
	"testing"
)

func TestGetDNS(t *testing.T) {
	_, err := GetDNS()

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLocation(t *testing.T) {
	_, err := GetLocation("8.8.8.8", LOCAL_ENGLISH)

	if err != nil {
		t.Fatal(err)
	}
}
