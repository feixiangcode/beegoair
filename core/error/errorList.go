package error

const (
	ERR_UNKNOWN = -1000
	ERR_NOAUTH = -2000
	ERR_SELECT = -2001
	ERR_UPDATE = -2002
	ERR_STORE = -2003

	ERR_HTTP_400 = -400
	ERR_HTTP_401 = -401
	ERR_HTTP_402 = -402
	ERR_HTTP_403 = -404
	ERR_HTTP_404 = -404
	ERR_HTTP_405 = -405
	ERR_HTTP_406 = -406
	ERR_HTTP_407 = -407
	ERR_HTTP_408 = -408
	ERR_HTTP_409 = -409
	ERR_HTTP_410 = -410
	ERR_HTTP_500 = -500
	ERR_HTTP_501 = -501
	ERR_HTTP_502 = -502
	ERR_HTTP_503 = -505
	ERR_HTTP_504 = -504
	ERR_HTTP_505 = -505
)

var ErrorList *List

type Item struct {
	Code   int
	Status int
	Msg    string
}

type List struct {
	maps map[int]Item
}

func (this *List) Add(code int, msg string, status int) {
	if this.maps == nil {
		this.maps = make(map[int]Item)
	}

	this.maps[code] = *&Item{
		Code:   code,
		Status: status,
		Msg:    msg,
	}
}

func (this *List) Get(code int) Item {
	if item, ok := this.maps[code]; ok {
		return item
	}
	return this.maps[ERR_UNKNOWN]
}

func init() {
	ErrorList = new(List)
	ErrorList.Add(ERR_UNKNOWN, "Unknown error", 200)
	ErrorList.Add(ERR_NOAUTH, "Noauth error", 200)
	ErrorList.Add(ERR_SELECT, "Find error", 200)
	ErrorList.Add(ERR_UPDATE, "Update error", 200)
	ErrorList.Add(ERR_STORE, "Store error", 200)

	ErrorList.Add(ERR_HTTP_404, "404 Not Found", 404)
	ErrorList.Add(ERR_HTTP_400, "400 Bad Request", 400)
	ErrorList.Add(ERR_HTTP_401, "401 Unauthorized", 401)
	ErrorList.Add(ERR_HTTP_402, "402 Payment Required", 402)
	ErrorList.Add(ERR_HTTP_403, "403 Forbidden", 403)
	ErrorList.Add(ERR_HTTP_404, "404 Not Found", 404)
	ErrorList.Add(ERR_HTTP_405, "405 Method Not Allowed", 405)
	ErrorList.Add(ERR_HTTP_406, "406 Not Acceptable", 406)
	ErrorList.Add(ERR_HTTP_407, "407 Proxy Authentication Required", 407)
	ErrorList.Add(ERR_HTTP_408, "408 Request Timeout", 408)
	ErrorList.Add(ERR_HTTP_409, "409 Conflict", 409)
	ErrorList.Add(ERR_HTTP_410, "410 Gone", 410)
	ErrorList.Add(ERR_HTTP_500, "500 Internal Server Error", 500)
	ErrorList.Add(ERR_HTTP_501, "501 Not Implemented", 501)
	ErrorList.Add(ERR_HTTP_502, "502 Bad Gateway", 502)
	ErrorList.Add(ERR_HTTP_503, "503 Service Unavailable", 503)
	ErrorList.Add(ERR_HTTP_504, "504 Gateway Timeout", 504)
	ErrorList.Add(ERR_HTTP_505, "505 HTTP Version Not Supported", 505)
}
