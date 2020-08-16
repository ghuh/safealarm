package main

import "testing"

func TestGetIfEnabledTrue(t *testing.T) {
    didRun := false
    test := getIfEnabled(true, func() { didRun = true })
    test();

    if didRun == false {
        t.Error("getIfEnabled failed to return the provided function")
    }
}

func TestGetIfEnabledFalse(t *testing.T) {
    didRun := false
    test := getIfEnabled(false, func() { didRun = true })
    test();

    if didRun == true {
        t.Error("getIfEnabled failed to return the empty function")
    }
}
