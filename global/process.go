package global

import (
	"github.com/yaoapp/gou"
)

func init() {
	// 注册处理器
	gou.RegisterProcessHandler("xiang.global.ping", processPing)
}

// processCreate 运行模型 MustCreate
func processPing(process *gou.Process) interface{} {
	res := map[string]interface{}{
		"code":    200,
		"server":  "象传应用引擎",
		"version": VERSION,
		"domain":  DOMAIN,
		"allows":  Conf.Service.Allow,
	}
	return res
}

// processFile 返回文件内容
func processFile(process *gou.Process) interface{} {
	return nil
}
