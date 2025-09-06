package validator

import (
    "log"
    "testing"
)

func TestPhoneNumber(t *testing.T) {
    t.Run("valid", func(t *testing.T) {
        for i, phone := range []string {
                "09121234567",
                "09129876543",
                "09365432109",
                "09217998990",
         } {
            validator := New()
            err := validator.AddRule(Phone(phone)).Validate()
            if err != nil {
                t.Fatalf("[test case %d] expected nil got %v", i, err)
            }
         }
    })

    t.Run("invalid", func(t *testing.T) {
        for i, phone := range []string {
            "",
            "0912",
            "0936",
            "1234567",
            "0214567895",
            "sdf349f9sdf",
        }{
            validator := New()
            err := validator.AddRule(Phone(phone)).Validate()
            if err != nil {
                t.Fatalf("[test case %d] expected nil got %v", i, err)
            }
        }
    })
}


