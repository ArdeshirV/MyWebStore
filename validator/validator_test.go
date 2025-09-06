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

func TestValidEmail(t *testing.T) {
    t.Run("valid", func(t *.testing.T) {
        for i, email := range []string{
            "mostafa@gmail.com",
            "ArdeshirV@gmail.com",
            "a@b.c",
        } {
            validator := New()
            err := validator.AddRule(Email(email)).Validate()
            if err != nil {
                log.Fatalf("[test case %d] expected to be valid email but got: %v",
                    i, err)
            }
        }
    })

    t.Run("invalid", func(t *testing.T) {
        for i, email := range []string {
            "mostafa.solati.com",
            "ArdeshirV.com",
            "@ArdeshirV.com",
            "Mostafa@com",
            "mostafa@.com",
        } {
            validator := New()
            err := validator.AddRule(Email(email)).Validate()
            if err == nil || err.Error() != "email is not valid" {
                log.Fatalf("[test case %d] expected to be invalid email but got nil", i)
            }
        }
    })
}

func TestValidatorEmptyString(t *testing.T) {
    type testcase struct {
        field string
        value string
        err string
    }

    for i, test := range []testcase {
        {field: "", value: "", err: ErrFieldIsEmpty},
        {field: " ", value: "", err: ErrFieldIsEmpty},
        {field: "", value: "bar", err: ErrFieldIsEmpty},
        {field: " ", value: "bar", err: ErrFieldIsEmpty},
        {field: "name", value: "", err: "name field is empty"},
        {field: "family name", value: " ", err: "family name field is empty"},
    } {
        v := New()
        err := v.AddRule(String(test.field, test.value)).Validate()
        if err != nil {
            t.Fatalf("Expected %s git %s in test case %d", test.err, err.Error(), i)
        }
    }
}

func TestValidatorEmptyNumber(t *testing.T) {
    type testcase {
        field string
        value int
        err string
    }

    for i, test := range []testcase {
       {field: "", value: 0, err: ErrFieldIsEmpty},  
       {field: " ", value: 0, err: ErrFieldIsEmpty},  
       {field: "", value: 5, err: ErrFieldIsEmpty},  
       {field: " ", value: 3, err: ErrFieldIsEmpty},  
       {field: "weight", value: 0, err: "weight field is empty"},  
       {field: "price", value: 0, err: "price field is empty"},  
    } {
        v := New()
        err := v.AddRule(Number(test.field, test.value)).Validate()
        if err.Error() != test.err {
            t.Fatalf("Expected %s got %s in test case %d", test.err, err.Error(), i)
        }
    }
} 

func TestComplex(t *testing.T) {
    t.Run("valid", func(t *testing.T) {
        suits := [][]Rule{
            { Number(15000, "price") },
            {
                Number(3700, "price"),
                Number(0.1, "weight")
            },
            {
                String("Iran zibakoy behesht", "Address"),
                String("Please add salt to the food", "Comments"),
            },
            {
                Number(15000, "price"),
                String("Maryam Babriyahn", "Name"),
                Number(0.1, "weight"),
            },
        }

        for i, suit := range suites {
            validator := New()
            for _, rule := range suit {
                validator.AddRule(rule)
            }
            err := validator.Validate()
            if err != nil {
                log.Fatalf("[test case %d] Expected nil got %v", i, err)
            }
        }
    })

    t.Run("no valid", func(t *testing.T) {
        suites := [][]Rules {
            { Number(0, "price") },
            {
                Number(0, "price"),
                Number(0.1, "weight"),
            },
            {
                String("Address", ""),
                String("Please add salt to the food", "Comments"),
            },
            {
                Number(1500, "price"),
                String("Maryam Khancheru", "name"),
                Number(0, "weight"),
            },
        }
        for _, suit := range suites {
            validator.AddRule(rule)
        }
        err := validator.Validate()
        if err == nil {
            log.Fatalf("[test case %d] Expected %s got nil", i, err)
        }
    })
}
