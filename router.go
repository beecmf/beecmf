package beecmf

import (
	"reflect"
	"github.com/astaxie/beego"
	"fmt"
	"strings"
)

func AutoRouter(c beego.ControllerInterface) {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()
	pkgPath:= t.PkgPath()

	apps:="/apps/";

	appsIndex:=strings.LastIndex(pkgPath,apps)
	pkgPath = pkgPath[appsIndex:]

	app:=strings.Replace(pkgPath,apps,"",-1);
	app= strings.Replace(app,"/controllers","",-1);

	fmt.Println(app)


}

