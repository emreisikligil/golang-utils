package runtime

import (
    "runtime"
    "strings"
)

func GetCallerFuncDetails(skip int) runtime.Frame {
    pCounters := make([]uintptr, 1)

    n := runtime.Callers(skip + 2, pCounters)

    if n > 0 {
        frames := runtime.CallersFrames(pCounters)
        f, _:= frames.Next()
        return f
    }

    return runtime.Frame{}
}

func GetCallerFuncQualifier(skip int) string {
    frame := GetCallerFuncDetails(skip + 1)

    return frame.Function
}

func GetCallerFuncName(skip int) string {
    name := GetCallerFuncQualifier(skip + 1)

    return name[strings.LastIndexAny(name, "/") + 1:]
}