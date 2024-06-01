package bootstrap

import (
	"database/sql"
	"demo/config"
	"demo/db"
	"github.com/lingdor/gmodel"
	"github.com/lingdor/spannerlib/E"
	"github.com/sirupsen/logrus"
	"os"
)

type App struct {
	conn   *sql.DB
	logger *logrus.Logger
}

func (a *App) printerr() {
	var err = recover()
	if err != nil {
		a.logger.Errorf("panic :%v", err)
		os.Exit(1)
	}
}

func (a *App) Prepare() {
	defer a.printerr()
	//init log
	logger := logrus.New()
	gmodel.SetLogger(logger)

	//load config
	config.Parse("./gmodel.yml")

	//load db
	a.conn = E.Must1(db.Connect(config.AppConfig.Gmodel.Connection["default"].Dsn))

}

func (a *App) Run() {

	//ctx :=context.Background()
	// where := orm.Eq(schema.Tb1Schema.id, 1)
	// var search = orm.Select(orm.CountAll()).From(schema.tb1Schema).
	//var arr = E.Must1(gmodel.QueryArrRowsContext(ctx, a.conn, search))
	//json := E.Must1(array.JsonMarshal(arr))
	//fmt.Println(string(json))
}
