package admin

import (
	"os"
	"runtime"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	this.display()
}
