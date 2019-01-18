package runtime

import (
    "fmt"
    "runtime"
    "testing"
)

const (
    PACKAGE_PREFIX = "github.com/emreisikligil/golang-utils/runtime."
)

func TestGetCallerFuncDetails(t *testing.T) {
    // Gets details of the calling function
    frame := callForFuncDetails(0)
    if frame.Function != PACKAGE_PREFIX + "callForFuncDetails" {
        t.Error("Failed to refer calling function")
    }

    // Gets details of the caller functions in the stack
    for i := 0; i < 5; i++ {
        frame = wrapCallForFuncDetails(i + 1, i)
        if frame.Function != PACKAGE_PREFIX + "wrapCallForFuncDetails" {
            t.Error(fmt.Sprintf("Failed to refer wrapping function %d", i))
        }
    }


}

func TestGetCallerFuncQualifier(t *testing.T) {
    // Gets qualified name of the calling function
    qualifier := callForFuncQualifier(0)
    if qualifier != PACKAGE_PREFIX + "callForFuncQualifier" {
        t.Error("Failed to refer calling function")
    }

    // Gets qualified name of the caller functions in the stack
    for i := 0; i < 5; i++ {
        qualifier = wrapCallForFuncQualifier(i + 1, i)
        if qualifier != PACKAGE_PREFIX + "wrapCallForFuncQualifier" {
            t.Error(fmt.Sprintf("Failed to refer wrapping function %d", i))
        }
    }
}

func TestGetCallerFuncName(t *testing.T) {
    // Gets name of the calling function
    name := callForFuncName(0)
    if name != "runtime.callForFuncName" {
        t.Error("Failed to refer calling function")
    }

    // Gets name of the caller functions in the stack
    for i := 0; i < 5; i++ {
        name = wrapCallForFuncName(i + 1, i)
        if name != "runtime.wrapCallForFuncName" {
            t.Error(fmt.Sprintf("Failed to refer wrapping function %d", i))
        }
    }
}

func callForFuncDetails(skip int) runtime.Frame {
    return GetCallerFuncDetails(skip)
}

func callForFuncQualifier(skip int) string {
    return GetCallerFuncQualifier(skip)
}

func callForFuncName(skip int) string {
    return GetCallerFuncName(skip)
}

func wrapCallForFuncDetails(skip int, additionalWraps int) runtime.Frame {
    if additionalWraps <= 0 {
        return callForFuncDetails(skip)
    } else {
        return wrapCallForFuncDetails(skip, additionalWraps - 1)
    }
}

func wrapCallForFuncQualifier(skip int, additionalWraps int) string {
    if additionalWraps <= 0 {
        return callForFuncQualifier(skip)
    } else {
        return wrapCallForFuncQualifier(skip, additionalWraps - 1)
    }
}

func wrapCallForFuncName(skip int, additionalWraps int) string {
    if additionalWraps <= 0 {
        return callForFuncName(skip)
    } else {
        return wrapCallForFuncName(skip, additionalWraps - 1)
    }
}
