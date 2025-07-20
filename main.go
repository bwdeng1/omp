package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "omp/internal/logic"
	_ "omp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"omp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
