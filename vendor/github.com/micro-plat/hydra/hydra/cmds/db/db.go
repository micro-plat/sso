package db

import (
	"fmt"
	"regexp"

	"github.com/lib4dev/cli/cmds"
	logs "github.com/lib4dev/cli/logger"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/global/compatible"
	"github.com/micro-plat/hydra/hydra/cmds/pkgs"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

func init() {
	cmds.RegisterFunc(func() cli.Command {
		return cli.Command{
			Name:  "db",
			Usage: "数据库, 数据库初始化管理",
			Subcommands: []cli.Command{
				{
					Name:   "install",
					Usage:  "-将数据表等安装到数据库",
					Flags:  getInstallFlags(),
					Action: install,
				},
			},
		}
	})
}
func install(c *cli.Context) (err error) {
	defer func() {
		logNow(err)
		err = nil
	}()

	//1. 绑定应用程序参数
	global.Current().Log().Pause()
	if err := global.Def.Bind(c); err != nil {
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}

	//2. 获取执行参数
	sqls := global.Installer.DB.GetSQLs()
	handlers := global.Installer.DB.GetHandlers()
	if len(sqls) == 0 && len(handlers) == 0 {
		return fmt.Errorf("未指定SQL或安装程序")
	}

	//2.检查是否安装注册中心配置
	if registry.GetProto(global.Current().GetRegistryAddr()) == registry.LocalMemory {
		if err := pkgs.Pub2Registry(true); err != nil {
			return err
		}
	}

	//3. 拉取注册中心配置
	if err := app.PullAndSave(); err != nil {
		return err
	}

	//4. 执行SQL语句
	if len(sqls) > 0 {
		db, err := components.Def.DB().GetDB(types.GetString(dbName, "db"))
		if err != nil {
			return err
		}
		for _, sql := range sqls {
			if _, err := db.Execute(sql, nil); err != nil {
				err = fmt.Errorf("%32s\t%w", getMessage(sql), err)
				if !skip {
					return err
				}
				logs.Log.Error(err, compatible.FAILED)
				continue
			}
			msg := fmt.Sprintf("%32s", getMessage(sql))
			logs.Log.Info(msg, compatible.SUCCESS)
		}
	}

	//5. 执行处理函数
	for _, handle := range handlers {
		if err := handle(); err != nil {
			return err
		}
	}
	return nil
}
func logNow(err error) {
	if err != nil {
		logs.Log.Error(err, compatible.FAILED)
		return
	}
}
func getMessage(input string) string {
	raw := input[:types.GetMin(32, len(input))]
	regx := regexp.MustCompile("[^\\(\\[]+")
	nstr := regx.FindString(raw)
	return nstr
}

/*
raw := input[:types.GetMin(32, len(input))]
	return strings.TrimSpace(strings.Replace(raw, "\n" "", -1))
*/
