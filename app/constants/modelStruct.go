package constants

type Test struct {
    Id        int64  `json:"id" orm:"column(id)"`
    Name      string `json:"name" orm:"column(Name)"`
    Age       int    `json:"age" orm:"column(age)"`
    Ctime     int64  `json:"ctime" orm:"column(ctime)"`
    Utime     int64  `json:"utime" orm:"column(utime)"`
}
