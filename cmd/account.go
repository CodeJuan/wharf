package cmd

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/codegangsta/cli"
	"github.com/dockerclouds/docker-hub/models"
	"github.com/dockerclouds/docker-hub/utils"
)

var CmdAccount = cli.Command{
	Name:        "account",
	Usage:       "通过命令行管理系统的账户",
	Description: "通过命令行添加、激活、停用 Hub 中的用户账户，账户停用后该账户下公开的 Repository 依旧可以下载。",
	Action:      runAccount,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "action",
			Value: "",
			Usage: "Action 参数: add/active/unactive，添加激活的账户/添加待激活的账户/停用账户",
		},
		cli.StringFlag{
			Name:  "email",
			Value: "",
			Usage: "账户邮件地址",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "账户名",
		},
		cli.StringFlag{
			Name:  "passwd",
			Value: "",
			Usage: "账户初始密码",
		},
	},
}

func runAccount(c *cli.Context) {
	var action, email, username, passwd string

	if len(c.String("action")) > 0 {
		models.InitDb()
		action = c.String("action")
		switch action {
		case "add":
			if len(c.String("username")) > 0 && len(c.String("email")) > 0 && len(c.String("passwd")) > 0 {
				username = c.String("username")
				email = c.String("email")
				passwd = c.String("passwd")

				user := new(models.User)
				if err := user.Add(username, passwd, email, true); err != nil {
					fmt.Println(fmt.Sprintf("%s: %s", "添加用户失败", err.Error()))
				} else {
					port, _ := beego.AppConfig.Int("email::Port")

					if err := utils.SendAddEmail(username, passwd, email,
						beego.AppConfig.String("email::Host"),
						port,
						beego.AppConfig.String("email::User"),
						beego.AppConfig.String("email::Password")); err != nil {
						fmt.Println("为用户发送新建账户邮件失败")
					} else {
						fmt.Println("添加用户成功并发送通知邮件")
					}
				}

			} else {
				fmt.Println("account add 命令必须指定 username/email/passwd 参数")
			}

			break
		case "active":
			break
		case "unactive":
			break
		default:
			fmt.Println("目前只支持 add/active/unactive 三个指令。")
		}
	} else {
		fmt.Println("需要指定操作用户的指令，仅支持 add/active/unactive 三个指令")
	}
}
