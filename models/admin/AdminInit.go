/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/
package admin

import (
	"fmt"
	"math/rand"

	"github.com/astaxie/beego"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
	InsertConfig()
	InsertCategory()
	InsertRoll()
	InsertPaper()
}

//插入网站配置
func InsertConfig() {
	fmt.Println("insert config start")
	c := new(blog.Config)
	c.Photo = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
	c.Id = 1
	c.Title = "Rabbit(Tuzi) Enterprise Web"
	c.Webinfo = `
	{
		"1":{"name":"About","limit":6},
        "2":{"name":"News","limit":6},
        "3":{"name":"Lifes","limit":6},
        "4":{"name":"Production","limit":6},
        "5":{"name":"Flower","limit":6},
        "6":{"name":"TeaCup","limit":6}
	}
	`
	c.Phone = "0750-12345678"
	c.Content = `
<div align="center">
	<p>
		<img src="/file/image/53/68756e7465726875671e5573ac53bb5813b6b51d47d2db806b.gif" alt="" />
	</p>
</div>
	`
	c.Slogan = "A Enterprise Web, You can have a try"
	c.Address = `<meta description="rabbit" >
<!-- some other script put in here -->`
	c.Code3 = `Power by hunterhug at 2017 此处页脚版权`
	c.Code2 = `Stats Code  此处放统计代码`
	c.Code1 = "Comment Code 此处放跟帖代码"
	err := c.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert config end")
}

func InsertCategory() {
	fmt.Println("insert category start")
	cs := map[int64]string{1: "About", 2: "News", 3: "Lifes", 4: "Production", 5: "Flower", 6: "TeaCup", 7: "Books", 8: "Musics"}
	for k, v := range cs {
		c := new(blog.Category)
		c.Id = k
		c.Title = v + "-T"
		c.Alias = v
		c.Createtime = lib.GetTime()
		c.Status = 1
		c.Image = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
		c.Content = v
		if k == 4 {
			c.Type = 1
		}
		if k > 4 {
			c.Pid = 4
			c.Type = 1
		}
		err := c.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("insert category end")
}

func InsertPaper() {
	aaa := []string{
		"/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"/file/image/12/68756e74657268756795ad72d42c7ef1b56c04c66297db1c27.jpeg",
		"/file/image/64/68756e74657268756759fc5ee18fa5210bb76003976900fae9.jpeg",
	}
	k := 140
	for k > 0 {
		k = k - 1
		paper := new(blog.Paper)
		paper.Title = "Test test test test data"
		paper.Status = 1
		paper.Photo = aaa[rand.Intn(3)]
		paper.Descontent = "淭一厘晢臹隒丌蒛霒冘，庳乜砐淈琲葙丌漍厹刌。仈"
		paper.Content = `
So her / Ce are Xi Zan, but what Xiao this Wu may not put up the Yun Zhang ping. These Jie Tong Kang Schrodinger
Jiu Da Chi Yu a Feng Bing Wu Dao tea, Ke Bei Fu Tu cou not drag type go to squall process gas Bureau ze. A Qu% Xiu
Yan Que Yin find not weighted, Bi E Qu Pei cases Guo are not Rou cut. He Ge howl Xi Cang Hu Bo Kuang Nan Xin loop Yi
Fu Ru Chu Le Xie row full of a surname. He got Shan creating Zhi w Gun and Dun Cu Qiong Jian Hu You Yi Wei Ding Hai Li.
After a Xuan Bei Tu yao type seal Zeng ang, Guo Wu You Liu Wen Torr + Quan market dissatisfied. A Xian how prepared and
not Zhu Che Han he Xia, Chu Zhong Chen Yi Yan not falsification of Ji E. A Gong Wu Fei Guo zhe Ju Er through, please.
Guan melancholy that he was Jiang urn donburi. A host Gui Yuncheng Qin TA type. Zhi Ji, Ji type He Xun Zan Mei not Jiao
Dao le. Kun a sincere Sun what taste as a surname Ju Qi, sow not You Si Xiao type Hou tonnes through the. A Yan Zhao
`
		paper.Author = "hunterhug"
		paper.Createtime = lib.GetTime()
		paper.Cid = int64(rand.Intn(8) + 1)
		if paper.Cid >= 4 {
			paper.Type = 1
		}
		paper.Istop = int64(rand.Intn(2))
		paper.Insert()
	}
}

func InsertRoll() {
	rolls := map[string]string{
		"tuzi":   "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"me":     "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
		"tuzizi": "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"me1":    "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
	}
	for k, v := range rolls {
		t := new(blog.Roll)
		t.Photo = v
		t.Status = 1
		t.Title = k
		t.Createtime = lib.GetTime()
		t.Insert()
	}
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = beego.AppConfig.String("rbac_admin_user")
	u.Nickname = "TuziAdmin"
	u.Password = lib.Pwdhash(beego.AppConfig.String("rbac_admin_user"))
	u.Email = "569929309@qq.com"
	u.Remark = "God in Rabbit Country"
	// 2 stand for close, but it has very high authority
	u.Status = 2
	u.Createtime = lib.GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "test"
	u1.Nickname = "TuziTest"
	u1.Password = lib.Pwdhash("test")
	u1.Email = "569929309@qq.com"
	u1.Remark = "Just a Test User"
	u1.Status = 1
	u1.Createtime = lib.GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "兔子后台"
	g.Title = "后台管理"
	g.Sort = 1
	g.Id = 1
	g.Status = 1
	e := g.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g1 := new(Group)
	g1.Name = "兔子后台"
	g1.Title = "文章管理"
	g1.Sort = 2
	g1.Id = 2
	g1.Status = 1
	e = g1.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g2 := new(Group)
	g2.Name = "兔子后台"
	g2.Title = "图片管理"
	g2.Sort = 3
	g2.Id = 3
	g2.Status = 1
	e = g2.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println("insert group end")
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	r.Insert()
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	g1 := new(Group)
	g1.Id = 2
	g2 := new(Group)
	g2.Id = 3
	nodes := []Node{
		/*

			RBAC管理中心

		*/
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},

		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},

		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},

		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},

		/*

			配置中心

		*/
		{Id: 25, Name: "config", Title: "配置中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		//-------
		//网站配置
		{Id: 26, Name: "option/index", Title: "网站配置", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 27, Name: "Index", Title: "网站配置首页", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 28, Name: "UpdateOption", Title: "更新网站配置", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		//网站配置
		//个人信息
		{Id: 29, Name: "user/index", Title: "个人信息", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 30, Name: "Index", Title: "个人信息首页", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 31, Name: "UpdateUser", Title: "更新个人信息", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		//个人信息

		/*

			文章中心

		*/
		{Id: 32, Name: "blog", Title: "文章中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g1},
		//------
		//文章目录
		{Id: 33, Name: "category/index", Title: "目录列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 34, Name: "Index", Title: "目录列表首页", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 35, Name: "AddCategory", Title: "增加目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 36, Name: "UpdateCategory", Title: "修改目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		//文章目录
		//文章
		{Id: 37, Name: "paper/index", Title: "文章列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 38, Name: "Index", Title: "文章列表首页", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 39, Name: "AddPaper", Title: "增加文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 40, Name: "UpdatePaper", Title: "修改文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 41, Name: "DeletePaper", Title: "回收文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 42, Name: "RealDelPaper", Title: "删除文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		//文章

		/*

			图片管理

		*/
		{Id: 43, Name: "picture", Title: "图片中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g2},
		//---------
		//相册
		{Id: 44, Name: "album/index", Title: "相册管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 45, Name: "Index", Title: "相册首页", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 46, Name: "AddAlbum", Title: "增加相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 47, Name: "DeleteAlbum", Title: "删除相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 48, Name: "UpdateAlbum", Title: "修改相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		//相册
		//图片
		{Id: 49, Name: "photo/index", Title: "图片管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 50, Name: "Index", Title: "图片首页", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 51, Name: "AddPhoto", Title: "增加图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 52, Name: "DeletePhoto", Title: "回收图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 53, Name: "UpdatePhoto", Title: "修改图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 54, Name: "RealDelPhoto", Title: "删除图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		//图片

		//补充的
		{Id: 55, Name: "DeleteCategory", Title: "删除目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},

		{Id: 56, Name: "paper/rubbish", Title: "文章回收站", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 57, Name: "Rubbish", Title: "文章回收站", Remark: "", Level: 3, Pid: 56, Status: 1, Group: g1},

		{Id: 58, Name: "photo/rubbish", Title: "图片回收站", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 59, Name: "Rubbish", Title: "图片回收站", Remark: "", Level: 3, Pid: 58, Status: 1, Group: g2},

		//首页图片轮转
		{Id: 60, Name: "roll/index", Title: "首页轮转", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 61, Name: "Index", Title: "轮转列表", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 62, Name: "AddRoll", Title: "增加轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 63, Name: "UpdateRoll", Title: "更新轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 64, Name: "DeleteRoll", Title: "删除轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Id = v.Id // 这句是无效的,后来 bug 被 beego 官方改好了
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		e := n.Insert()
		if e != nil {
			fmt.Printf("%#v:%#v\n", n, e.Error())
		}
	}
	fmt.Println("insert node end")
}
