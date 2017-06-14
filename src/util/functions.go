package util

import (

    "strings"
)

func cutToOne(input *string) string {
    return strings.Fields(*input)[0]
}