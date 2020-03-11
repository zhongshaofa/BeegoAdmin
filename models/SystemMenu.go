package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SystemMenu struct {
	Id       int       `json:"id"`
	Pid      int       `json:"pid"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Href     string    `json:"href"`
	Sort     string    `json:"sort"`
	Target   string    `json:"target"`
	Remark   string    `json:"remark"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"create_at";orm:"auto_now;type(datetime)"`
}

func (m *SystemMenu) TableName() string {
	return TableName("system_menu")
}

// 初始化结构体
type SystemInit struct {
	HomeInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct {
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"logoInfo"`
	MenuInfo []*MenuTreeList `json:"menuInfo"`
}

// 菜单结构体
type MenuTreeList struct {
	Id     int             `json:"id"`
	Pid    int             `json:"pid"`
	Title  string          `json:"title"`
	Icon   string          `json:"icon"`
	Href   string          `json:"href"`
	Target string          `json:"target"`
	Remark string          `json:"remark"`
	Child  []*MenuTreeList `json:"child"`
}

// 获取初始化数据
func (m *SystemMenu) GetSystemInit() SystemInit {
	var systemInit SystemInit

	// 首页
	systemInit.HomeInfo.Title = "首页"
	systemInit.HomeInfo.Href = "static/page/welcome-1.html?t=1"

	// logo
	systemInit.LogoInfo.Title = "BeeAdmin"
	systemInit.LogoInfo.Image = "static/images/logo.png"

	// 菜单
	systemInit.MenuInfo = m.GetMenuList()

	return systemInit
}

// 获取菜单列表
func (m *SystemMenu) GetMenuList() []*MenuTreeList {
	o := orm.NewOrm()
	var menuList []SystemMenu
	_, _ = o.QueryTable(m.TableName()).All(&menuList)
	return m.buildMenuChild(0, menuList)
}

//递归获取子菜单
func (m *SystemMenu) buildMenuChild(pid int, menuList []SystemMenu) []*MenuTreeList {
	var treeList []*MenuTreeList
	for _, v := range menuList {
		if pid == v.Pid {
			node := &MenuTreeList{
				Id:     v.Id,
				Title:  v.Title,
				Icon:   v.Icon,
				Href:   v.Href,
				Target: v.Target,
				Pid:    v.Pid,
			}
			child := v.buildMenuChild(v.Id, menuList)
			if len(child) != 0 {
				node.Child = child
			}
			treeList = append(treeList, node)
		}
	}
	return treeList
}
