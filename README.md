## go-config

Simple config wrapper over [viper](https://github.com/spf13/viper), customised pretty much for the way I use it.

Loads config from:
   - command line flags
   - environment variables (upper-cased flag values, hyphens replaced by underscores)
   - YAML config file in `~/.config/<binary-name>`

Example:
```go
import (
	"flag"
	"fmt"

	"github.com/venkytv/go-config"
)

func main() {
	flag.String("test-val", "123", "help for test")
	cfg := config.Load(nil, "FOO")
	fmt.Println(cfg.GetString("test-val"))
}
```

Usage:
```shell
./testprog --test foo
FOO_TEST_VAL=bar ./testprog
```
