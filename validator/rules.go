package validator

import (
    "errors"
    "fmt"
    "strings"
    "unicode"
    "golang.org/x/exp/constraints"
)

func Phone(value string) Rule {
    return func() error {
        if err := String("phone", value)(); err != nil {
            return err
        }
        value = strings.Map(func(r rune) rune {
            if unicode.IsDigit(r) {
                return r
            }
            return -1
        }, value)

        if strings.HasPrefix(value, "09") && len(value) == 11 {
            return nil
        }

        return errors.New("phone number is not valid")
    }
} 


