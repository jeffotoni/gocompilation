// Go in action
// @jeffotoni
// 2019-01-24

package main

import "testing"

func TestSum(t *testing.T) {

    t.Run("Sum=1,2", func(t *testing.T) {
        if 1 != 3 {
            t.Error("error A=1")
        }
    })
    t.Run("A=2,5", func(t *testing.T) {
        if 1 == 1 {
            println("success")
        }
    })
}
