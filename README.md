# Errors
Convenience functions for one-liners for building stacks of errors, within built-in error interface. Also handle appending to error with dual-output functions.

e.g.
```go
if v, err := someFunc(); err != nil {
	err = fmt.Errorf("Error doing <OPERATION> : %v\n", err)
	return interface{}, err
} else ...
```
becomes
```go
import "github.com/noahhai/errors"

return errors.GetFromTupleAdd("Error during <OPERATION>")(someFunc())
```

### Example
```go
import sysErrors "errors"
import "github.com/noahhai/errors"

mockApiFunc := func() error {
    return sysErrors.New("unexpected column found")
}
modelFunc := func() *errors.Error {
    return errors.From(mockApiFunc()).Add("failed to upsert value")
}
businessFunc := func() *errors.Error {
    return modelFunc().AddF("failed to perform daily sync of customer '%d'", 35)
}

// consume
handlerFunc := func() {
    if err := businessFunc(); err != nil {
        fmt.Println("Error handling request")
        fmt.Println("Error: ")
        fmt.Println(err)
        fmt.Println()
        fmt.Println("Cause:")
        fmt.Println(err.Cause())
        fmt.Println()
        fmt.Println("Symptom:")
        fmt.Println(err.Symptom())
    }
}
handlerFunc()

// return as built-in Error, e.g. to AWS API Gateway
extHandlerFunc := func() error {
    return businessFunc()
}
_ = extHandlerFunc()
```