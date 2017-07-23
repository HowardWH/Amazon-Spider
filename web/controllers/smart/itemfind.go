package smart

import (
	"encoding/csv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
)

type ItemFindController struct {
	baseController
}

func (this *ItemFindController) Index() {
	this.Layout = this.GetTemplate() + "/base/layout.html"
	this.TplName = this.GetTemplate() + "/back/listitemfind.html"

}

func (this *ItemFindController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	page, _ := this.GetInt("page", 1)
	rows, _ := this.GetInt("rows", 30)
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	start := (page - 1) * rows
	num := 0
	var maps []orm.Params
	if date == "" {
		date = "20161101"
	}

	dudu := "SELECT * FROM goods_info where Createtime='" + date + "' limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

	DB.Raw(dudu).Values(&maps)

	dudu1 := "SELECT count(*) as num FROM goods_info where Createtime='" + date + "';"

	DB.Raw(dudu1).QueryRow(&num)

	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}

func (this *ItemFindController) Export() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	dudu := "SELECT * FROM goods_info where Createtime='" + date + "'"
	var maps []orm.Params
	num, err := DB.Raw(dudu).Values(&maps)
	if err != nil {
		this.Rsp(false, "内部错误："+err.Error())
	}
	if num == 0 {
		this.Rsp(false, "日期找不到或查询结果为空:"+dudu)
	}
	f, err := os.Create("file/back/find" + date + ".csv")
	if err != nil {
		this.Rsp(false, "Excel创建出错")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	//	SELECT `goods_info`.`id`,
	//    `goods_info`.`Rank`,
	//    `goods_info`.`ASIN`,
	//    `goods_info`.`DetailUrl`,
	//    `goods_info`.`Condition`,
	//    `goods_info`.`Star`,
	//    `goods_info`.`Reviews`,
	//    `goods_info`.`Createtime`,
	//    `goods_info`.`col1`,
	//    `goods_info`.`col2`,
	//    `goods_info`.`col3`
	//FROM `smart_backstage`.`goods_info`;
	w.Write([]string{"Rank", "ASIN", "DetailUrl", "Condition", "Star", "Reviews", "Createtime"})
	for _, i := range maps {
		temp := map[string]string{}
		for index, j := range i {
			innertemp := ""
			if j == nil {
				innertemp = " "
			} else {
				switch j.(type) { //多选语句switch
				case string:
					//是字符时做的事情
					innertemp = j.(string)
				case int:
					innertemp = strconv.Itoa(j.(int))
				}
			}
			temp[index] = innertemp
		}
		w.Write([]string{temp["Rank"], temp["ASIN"], temp["DetailUrl"], temp["Condition"], temp["Star"], temp["Reviews"], temp["Createtime"]})

		// iscatch:1
		// asin:B000BVXDZM
		// dbname:19
	}
	w.Flush()

	this.Rsp(true, "/file/back/find"+date+".csv")
}
