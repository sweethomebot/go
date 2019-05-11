package core

import (
	"strings"
	"strconv"
)

// System
func IsSystemModule(ch Channel, pluginName string) bool {
	return GetBool(SendRequestAndWait(ch, "system", "IsSystemModule", CreText(pluginName)))
}
func GetDataPath(ch Channel) string {
	return GetText(SendRequestAndWait(ch, "global", "GetDataPath", CreNull()))
}
func GetHardwarePlatform(ch Channel) string {
	return GetText(SendRequestAndWait(ch, "system", "GetHardwarePlatform", CreNull()))
}

// SQL shortcut

func QueryCheckTable(ch Channel, tableName, columName string) string {
	return GetText(SendRequestAndWait(ch, "sql", "checkTable", CreArrString([]string{tableName, columName})))
}

func Query(ch Channel, sql string) bool {
	return GetBool(SendRequestAndWait(ch, "sql", "exec", CreText(sql)))
}

func QueryRow(ch Channel, sql string) map[string]string {
	return GetMap(SendRequestAndWait(ch, "sql", "queryRow", CreText(sql)))
}

func QueryRows(ch Channel, sql string) []map[string]string {
	return GetArrMap(SendRequestAndWait(ch, "sql", "queryRows", CreText(sql)))
}

// Log shortcut

func LogD(ch Channel, text string) {
	SendRequest(ch, "log", "d", CreText(text))
}
func LogI(ch Channel, text string) {
	SendRequest(ch, "log", "i", CreText(text))
}
func LogW(ch Channel, text string) {
	SendRequest(ch, "log", "w", CreText(text))
}
func LogE(ch Channel, text string) {
	SendRequest(ch, "log", "e", CreText(text))
}

// Translate shortcut

func Trans(ch Channel, key string) string {
	return GetText(SendRequestAndWait(ch, "language", "Trans", CreText(key)))
}
func TransPlugin(ch Channel, plugin string, key string) string {
	return GetText(SendRequestAndWait(ch, "language", "TransPlugin", CreMap(map[string]string{"plugin":plugin, "key":key})))
}
func TransVar(ch Channel, key string, vars []string) string {
	text := Trans(ch, key)
	//i := 0
	for i, str := range vars{
		text = strings.Replace(text, "$"+strconv.Itoa(i), str, -1)
	}
	return text
}
func TransPluginVar(ch Channel, plugin string, key string, vars []string) string {
	text := TransPlugin(ch, key, plugin)
	//i := 0
	for i, str := range vars{
		text = strings.Replace(text, "$"+strconv.Itoa(i), str, -1)
	}
	return text
}

// Setting

func SettingSet(ch Channel, key string, value string) bool {
	return GetBool(SendRequestAndWait(ch, "setting", "set", CreMap(map[string]string{"key":key, "value":value})))
}
func SettingGet(ch Channel, key string) string {
	return GetText(SendRequestAndWait(ch, "setting", "get", CreMap(map[string]string{"key":key})))
}
func SettingGetBool(ch Channel, key string) bool {
	return SettingGet(ch, key) == "1"
}
func SettingGetDefault(ch Channel, key string, defaultValue string) string {
	data := SettingGet(ch, key)
	if data == "" {
		return defaultValue
	}
	return data
}
func SettingSetSystem(ch Channel, key string, value string) bool {
	return GetBool(SendRequestAndWait(ch, "setting", "setSystem", CreMap(map[string]string{"key":key, "value":value})))
}
func SettingGetSystem(ch Channel, key string) string {
	return GetText(SendRequestAndWait(ch, "setting", "getSystem", CreMap(map[string]string{"key":key})))
}
func SettingGetSystemBool(ch Channel, key string) bool {
	return SettingGetSystem(ch, key) == "1"
}
func SettingGetSystemDefault(ch Channel, key string, defaultValue string) string {
	data := SettingGetSystem(ch, key)
	if data == "" {
		return defaultValue
	}
	return data
}


// Permission

func PermissionSet(ch Channel, usrId string, key string, value string, usrIdSelf string) bool {
	return GetBool(SendRequestAndWaitUsrId(ch, "permission", "set", CreMap(map[string]string{"usr_id": usrId, "key":key, "value":value}), usrIdSelf))
}
func PermissionGet(ch Channel, usrId string, key string) string {
	return GetText(SendRequestAndWait(ch, "permission", "get", CreMap(map[string]string{"usr_id": usrId, "key":key})))
}
func PermissionGetBool(ch Channel, usrId string, key string) bool {
	return PermissionGet(ch, usrId, key) == "1"
}
func PermissionSetSystem(ch Channel, usrId string, key string, value string, usrIdSelf string) bool {
	return GetBool(SendRequestAndWaitUsrId(ch, "permission", "setSystem", CreMap(map[string]string{"usr_id": usrId, "key":key, "value":value}), usrIdSelf))
}
func PermissionGetSystem(ch Channel, usrId string, key string) string {
	return GetText(SendRequestAndWait(ch, "permission", "getSystem", CreMap(map[string]string{"usr_id": usrId, "key":key})))
}
func PermissionGetSystemBool(ch Channel, usrId string, key string) bool {
	return PermissionGetSystem(ch, usrId, key) == "1"
}

// Register

func RegisterAsDriver(ch Channel, icon string, drv_url string, title string, form string) bool {
	// INFO: needs: DvcCreateRequest, DrvExec
	data := map[string]string{}
	data["icon"] = icon
	data["url"] = drv_url
	data["title"] = title
	data["form"] = form

	return GetBool(SendRequestAndWait(ch, "hhc", "registerAsDriver", CreMap(data)))
}

func RegisterAsWebPlugin(ch Channel, icon string, title string) bool {
	// INFO: needs: ApiFunctionList
	data := map[string]string{}
	data["icon"] = icon
	data["title"] = title

	return GetBool(SendRequestAndWait(ch, "httpServer", "registerAsWebPlugin", CreMap(data)))
}

func RegisterASettingForm(ch Channel, form string) bool {
	return GetBool(SendRequestAndWait(ch, "setting", "registerASettingForm", CreText(form)))
}

func RegisterAPermissionForm(ch Channel, form string) bool {
	return GetBool(SendRequestAndWait(ch, "permission", "registerAPermissionForm", CreText(form)))
}