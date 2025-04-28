package scriptsdk

import (
	"github.com/lz520520/mentalout_sdk/modules/mentalout/cs_common_models"
	"github.com/lz520520/mentalout_sdk/modules/mentalout/mental_models"
	"github.com/lz520520/mentalout_sdk/pkg/utils/serial"
)

type ScriptApi interface {
	Eval(params *serial.ControlParams) (respParams *serial.ControlParams, err error)
	PluginEval(name string, taskId string, params map[string]any) mental_models.CommonResp

	GetPluginInfo() *cs_common_models.MentalPlugin
	GetC2Profile() mental_models.C2Profile
	LogPrint(format string, v ...interface{})
	UploadFile(filename string, fileBytes []byte) (err error)
}
