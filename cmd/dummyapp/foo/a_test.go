package foo

import "testing"

func TestGetA(t *testing.T) {
    if GetA(true) != "a" {
        t.Errorf("failed...")
    }
}
