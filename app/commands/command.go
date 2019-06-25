package commands

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/astaxie/beego"
)

type Command struct {

}

func (this *Command) Test(cli *cli.Context) {
    uid := cli.Int("uid")
    fmt.Print(uid)
    fmt.Println(beego.AppConfig.String("httpport"))
}
