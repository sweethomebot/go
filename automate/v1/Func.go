package automate

import (
	"fmt"
	"github.com/sweethomebot/go/core/v1"
	"regexp"
	"strings"
)

func GetVar(ch core.Channel, localProgramId, varName string) string {
	data := map[string]string{}
	data["localProgramId"] = localProgramId
	data["varName"] = varName
	return core.GetText(core.SendRequestAndWait(ch, "automate", "GetVar", core.CreMap(data)))
}

func GetBlockConfig(block Block, key, def string) string {
	if core.HasKey(block.Config, key){
		return block.Config[key]
	}
	return def
}

func GetBlockConfigVarReplaced(ch core.Channel, block Block, key, def string) string {
	if core.HasKey(block.Config, key){
		value := block.Config[key]
		if strings.Contains(value, "$") {
			reg, err := regexp.Compile("\\$([a-zA-Z0-9_]+)\\$")
			if err != nil {
				fmt.Println("automate GetBlockConfig Regex Error: ", err)
			}
			vars := reg.FindAllString(value, -1)
			for _, varName := range vars {
				varName = varName[1:len(varName)-1]
				value = strings.ReplaceAll(value, "$"+varName+"$", GetVar(ch, block.ProgramId, varName) )
			}
		}
		return value
	}
	return def
}

func CreBlockExecAnswer(next BlockNextConn, output map[string]string) core.MessageData {
	return core.CreJson(BlockExecAnswer{next, output})
}