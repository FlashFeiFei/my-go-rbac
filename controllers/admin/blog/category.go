package blog

import (
	"github.com/astaxie/beego/orm"
	. "github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

type CategoryController struct {
	baseController
}

func (this *CategoryController) Index() {
	category := new(blog.Category)
	categorys := []orm.Params{}
	isajax, _ := this.GetInt("isajax", 0)
	selector := category.Query().Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime")

	if isajax == 1 || this.IsAjax() {
		status, _ := this.GetInt64("status", 0)
		mulu, _ := this.GetInt64("mulu", 0)

		if mulu != 0 {
			selector = selector.Filter("Pid", mulu)
		}
		if status == 0 {
			selector.Values(&categorys)
		} else {
			selector.Filter("Status", status).Values(&categorys)
		}
		count := len(categorys)
		// beego.Trace("%v", categorys)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &categorys}
		this.ServeJSON()
		return
	} else {
		selector.Values(&categorys)
		this.Data["category"] = &categorys
		this.Layout = this.GetTemplate() + "/blog/layout.html"
		this.TplName = this.GetTemplate() + "/blog/listcate.html"
	}
}

func (this *CategoryController) AddCategory() {
	isajax, _ := this.GetInt("isajax", 0)
	if isajax == 1 || this.IsAjax() {
		status := false
		message := ""
		category := new(blog.Category)
		category.Createtime = GetTime()
		category.Title = this.GetString("title", "")
		category.Pid, _ = this.GetInt64("mulu", 0)
		category.Sort, _ = this.GetInt64("order", 0)
		category.Status, _ = this.GetInt64("status", 2)
		category.Content = this.GetString("content", "")
		category.Image = this.GetString("photo", "")
		category.Alias = this.GetString("alias", "")
		category.Siteid = beautyid
		category.Type = blogtype
		err := category.Insert()
		if err != nil {
			message = err.Error()
		} else {
			status = true
			message = "增加成功"
		}
		this.Rsp(status, message)
	} else {
		category := new(blog.Category)
		categorys := []orm.Params{}
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys)
		this.Data["category"] = &categorys
		this.TplName = this.GetTemplate() + "/blog/addcate.html"
	}
}


func (this *CategoryController) UpdateCategory() {
	small, _ := this.GetInt64("small", 0)
	id, _ := this.GetInt64("id", 0)
	//小更改
	if small == 1 {
		status, _ := this.GetInt64("status", 0)
		if id == 0 || status == 0 {
			this.Rsp(false, "有问题")
		} else {
			category := new(blog.Category)
			category.Id = id
			category.Status = status
			category.Updatetime = GetTime()
			err := category.Update("Status", "Updatetime")
			if err != nil {
				this.Rsp(false, "更新失败")
			} else {
				this.Rsp(true, "更新成功")
			}
		}
	} else {
		isajax, _ := this.GetInt("isajax", 0)
		if isajax == 1 || this.IsAjax() {
			//大更改
			thiscategory := new(blog.Category)
			thiscategory.Id = id
			thiscategory.Title = this.GetString("title", "")
			thiscategory.Pid, _ = this.GetInt64("mulu", 0)
			thiscategory.Sort, _ = this.GetInt64("order", 0)
			thiscategory.Status, _ = this.GetInt64("status", 2)
			thiscategory.Content = this.GetString("content", "")
			thiscategory.Alias = this.GetString("alias", "")
			thiscategory.Updatetime = GetTime()
			//不存在则不改图片
			photo := this.GetString("photo", "")
			//beego.Trace("图片：" + photo)
			var err error
			if photo != "" {
				thiscategory.Image = photo
				err = thiscategory.Update("Title", "Pid", "Sort", "Status", "Content", "Updatetime", "Image", "Alias")
			} else {
				err = thiscategory.Update("Title", "Pid", "Sort", "Status", "Content", "Updatetime", "Alias")
				//beego.Trace("空图片：" + photo)
			}
			if err != nil {
				this.Rsp(false, err.Error())
			} else {
				this.Rsp(true, "更改成功")
			}
		} else {
			if id == 0 {
				this.Rsp(false, "没有id参数")
			}
			//显示更改页面
			thiscategory := new(blog.Category)
			thiscategory.Id = id
			err := thiscategory.Read()
			if err != nil {
				this.Rsp(false, "不存在该目录或者数据库出错")
			}
			this.Data["thiscategory"] = thiscategory

			category := new(blog.Category)
			categorys := []orm.Params{}
			category.Query().Exclude("Id", id).Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys)
			this.Data["category"] = &categorys

			this.TplName = this.GetTemplate() + "/blog/updatecate.html"
		}

	}
}

func (this *CategoryController) DeleteCategory() {
	category := new(blog.Category)
	id, err := this.GetInt64("id", 0)
	if err != nil || id == 0 {
		this.Rsp(false, "出现错误")
	}
	num, err := category.Query().Filter("Id", id).Filter("Siteid", beautyid).Filter("Type", blogtype).Count()
	if err != nil {
		this.Rsp(false, err.Error())
	} else if num == 0 {
		this.Rsp(false, "找不到该目录")
	} else {
		paper := new(blog.Paper)
		num1, err1 := paper.Query().Filter("Cid", id).Count()
		if num1 != 0 {
			this.Rsp(false, "目录下有小东西")
		} else if err1 != nil {
			this.Rsp(false, err1.Error())
		} else {
			num2, err2 := category.Query().Filter("Pid", id).Count()
			if err2 != nil {
				this.Rsp(false, err2.Error())
			} else if num2 != 0 {
				this.Rsp(false, "目录下有目录")
			} else {
				category.Id = id
				err3 := category.Delete()
				if err3 != nil {
					this.Rsp(false, err2.Error())
				} else {
					this.Rsp(true, "删除成功")
				}
			}
		}
	}
}
