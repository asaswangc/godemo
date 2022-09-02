package select_service

import (
	"goframework/src/app/cmdb/handler/enter"
	"goframework/src/framework/result"
)

// GetList 通用查询接口
func GetList(Base *enter.BaseModel, Page int, Size int, like bool) (scans []map[string]interface{}, total int64) {
	GDB := enter.LikeScopes(Base, like)
	result.Result(GDB.Count(&total)).Process()(result.SelectErr)
	result.Result(GDB.Scopes(enter.Paginate(Page, Size)).Omit("tenant_id").Find(&scans)).Process()(result.SelectErr)
	return
}
