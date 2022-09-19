package domain

import (
	"fmt"
	"testing"
)

func TestCheck4DBnbDomain(t *testing.T) {
	for i := 0; i < 9999; i++ {
		domain := fmt.Sprintf("%04d", i)
		ok, succeed, err := CheckBnbDomain([]string{}, domain)
		if err != nil {
			continue
		}

		if !succeed {
			t.Log("need change proxy")
			continue
		}

		if ok {
			t.Logf("%s can registe", domain)
		}
	}
}
