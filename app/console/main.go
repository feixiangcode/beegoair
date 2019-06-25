package main

import (
    _ "beegoair/app/init"
    "github.com/urfave/cli"
    "os"
    "github.com/astaxie/beego"
    "beegoair/core/application"
    "beegoair/app/commands"
    "fmt"
)

func main() {
    beego.LoadAppConfig("ini","../conf/app.conf")
    application.ConsoleRun()
    app := cli.NewApp()
    app.Commands = []cli.Command{
        {
            Name:  "test",
            Usage: "test --uid=",
            Action: (&commands.Command{}).Test,
            Flags: []cli.Flag{
                cli.IntFlag{Name: "uid",Usage:"--uid"},
            },
        },
    }
    err := app.Run(os.Args)
    if err != nil {
        fmt.Print("command error :" + err.Error())
    }
}
