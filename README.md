# go-interface-method-type-bug

Run `go run main.go`.

Go 1.10.5 outputs:

```
func (github.com/ccbrown/go-interface-method-type-bug/thepackage.A).Foo()
func (github.com/ccbrown/go-interface-method-type-bug/thepackage.B).Foo()
```

Go 1.11.2 outputs:

```
func (interface).Foo()
func (github.com/ccbrown/go-interface-method-type-bug/thepackage.B).Foo()
```
