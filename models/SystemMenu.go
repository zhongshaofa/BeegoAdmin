package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type SystemMenu struct {
	Id       int       `json:"id"`
	Pid      int       `json:"pid"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Href     string    `json:"href"`
	Target   string    `json:"target"`
	Remark   string    `json:"remark"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"create_at";orm:"auto_now;type(datetime)"`
}

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

func (m *SystemMenu) TableName() string {
	return TableName("system_menu")
}

/**
菜单列表
*/
func (m *SystemMenu) MenuList() []*MenuTreeList {
	return m.getMenu(0)
}

/**
递归获取树形菜单
*/
func (m *SystemMenu) getMenu(pid int) []*MenuTreeList {
	o := orm.NewOrm()
	var menu []SystemMenu
	_, _ = o.QueryTable(m.TableName()).Filter("pid", pid).OrderBy("sort").All(&menu)
	fmt.Println("======遍历========")
	fmt.Println(menu[0])

	treeList := []*MenuTreeList{}
	for _, v := range menu {
		child := v.getMenu(v.Id)
		node := &MenuTreeList{
			Id:     v.Id,
			Title:  v.Title,
			Icon:   v.Icon,
			Target: v.Target,
			Pid:    v.Pid,
		}
		node.Child = child
		treeList = append(treeList, node)
	}
	return treeList
}
