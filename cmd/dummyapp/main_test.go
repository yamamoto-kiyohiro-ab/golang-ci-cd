
package main

import "testing"

func TestGetHoge(t *testing.T) {
    if getHoge() != "hoge" {
        t.Errorf("failed...")
    }
}

func TestGetFuga(t *testing.T) {
    if getFuga() != "fuga" {
        t.Errorf("failed...")
    }
}

func TestGetPuyo(t *testing.T) {
    if getPuyo() != "puyo" {
        t.Errorf("failed...")
    }
}
