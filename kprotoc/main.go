package main // 测试

import (
	"flag"
	"fmt"
	"os"

	"github.com/taodev/kproto"
	"github.com/taodev/kproto/kprotoc/plugin"
)

var (
	argLang = flag.String("lang", "go", "convert language:go ts")
	skipRPC = flag.String("skip-rpc", "false", "skip remote rpc")
	fname   string
)

func parseArgs() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	fname = flag.Arg(0)
}

func main() {
	parseArgs()

	var err error

	generater := plugin.GetPlugin(*argLang)
	if generater == nil {
		fmt.Println(fname, "Unknown langage:", *argLang)
		os.Exit(1)
	}

	file, err := kproto.LoadProtoFile1(fname)
	if err != nil {
		fmt.Println("#", fname)
		fmt.Println(err)
		os.Exit(1)
	}

	if err = generater.Generate(file); err != nil {
		fmt.Println(fname, err)
		os.Exit(2)
	}
}
