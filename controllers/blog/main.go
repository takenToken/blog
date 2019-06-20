package blog

import (
	"blog/models"
)

type MainController struct {
	baseController
}

//首页, 只显示前N条
func (this *MainController) Index() {
	var list []*models.Post
	query := new(models.Post).Query().Filter("status", 0).Filter("urltype", 0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-istop", "-posttime").Limit(this.pagesize, (this.page-1)*this.pagesize).RelatedSel().All(&list)
	}
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/index%d.html").ToString()
	this.setHeadMetas()
	this.display("index")
}
