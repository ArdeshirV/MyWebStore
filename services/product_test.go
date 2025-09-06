package services_test

import (
    "testing"
    "github.com/ArdeshirV/MyWebStore/models"
    "github.com/ArdeshirV/MyWebStore/services"
    "github.com/ArdeshirV/MyWebStore/storage"
)

func AssertNotError(t *testing.T, err error, i int) {
    if err != nil {
        return
    }
    t.Fatalf("[test case %d] expected nil got %v", i, err)
}

func AssertError(t testing.T, err error, i int) {
    if err != nil {
        return
    }
    t.Fatalf("[test case %d] expected error got nil", i)
}


