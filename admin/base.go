package admin

type Result struct {
	Status  bool     `json:"status"`
	Code    int32    `json:"code"`
	message string   `json:"message"`
	data    struct{} `json:"data"`
}

// 返回json数据
func (res Result) Success(code int, message string, data struct{}) {

	//result["status"] = true

}

/*func Success()  {
	map["status"] = true
	map["code"] = 200
	map["message"] = ""
	return map
}
*/
