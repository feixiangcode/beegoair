package init

import (
    "github.com/astaxie/beego/orm"
    "beegoair/app/constants"
)

func init() {
    orm.RegisterModel(new(constants.Test))
}
