```go
package __commonpkg

import (
	"flag"
	"fmt"
	"testing"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var age = flag.Int("age", 1, "age")
var temperature = CelsiusFlag("temperature", 18, "temperature")

func TestFlagValue(t *testing.T) {
	flag.Parse()
	fmt.Printf("age is %d\n", *age)
	fmt.Printf("temperature is %s\n", temperature)
}
```

编译

```
go test -c -o flag-value-test.exe .\7_command_line_test.go
```

运行

```
~> .\flag-value-test.exe -temperature 7C
age is 1
temperature is 7°C
PASS
```

