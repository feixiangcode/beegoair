package errors

import error2 "beegoair/core/error"

const ERR_APP = 2000
const ERR_PARAM = 2001
func init() {
    error2.ErrorList.Add(ERR_APP,"app error", 200)
    error2.ErrorList.Add(ERR_PARAM,"param error", 200)
}
