package models

/** 返回前端的DataTable数据源 */
type DataTableResult struct {
	Draw            int64       `json:"draw"`
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int64       `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
	Error           string      `json:"error"`
}

/**
 * Ajax请求返回的结果
 */
type AjaxResult struct {
	Succ  bool        `json:"succ"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
