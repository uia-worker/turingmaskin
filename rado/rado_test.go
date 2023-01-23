package rado

import (
  "reflect"
  "testing"
)

func TestRadoTM(t *testing.T) {
    type test struct {
         input float64
         want float64
    }
    tests := []test{
        {input: 2, want: 20736},
        {input: 3, want: 16777216},
        {input: 4, want: 3.56e+10},
    }

    for _, tc := range tests {
        got := RadoTM(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    }
}
