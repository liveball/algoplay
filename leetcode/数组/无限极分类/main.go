package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cf := &Classify{}
	tgs := cf.Create(4)

	//迭代法
	// tree := getTreesByIteration(tgs)
	// spew.Dump(tree)

	//递归法
	tree := getTreesByRecursion(tgs, 0)
	spew.Dump(tree)

	chids := getChildsByPIDByRecursion(tgs, 1)
	fmt.Println("chids", chids)

	parents := getParentsByIDByRecursion(tgs, 3)
	fmt.Println("parents", parents)
}

type Tag struct {
	ID       int    `json:"id"`
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	Children []*Tag `json:"children"`
}

type Classify struct {
	Tags []*Tag
}

func (c *Classify) Create(level int) []*Tag {
	res := make([]*Tag, 0, level)
	for i := 0; i < level; i++ {
		t := &Tag{
			PID:  i,
			ID:   i + 1,
			Name: fmt.Sprintf("%d级", i+1),
		}
		res = append(res, t)
	}

	return res
}

//根据子节点id，获取所有父节点id
func getParentsByIDByRecursion(tgs []*Tag, id int) (res []int) {
	for _, p := range tgs {
		if p.ID == id {
			if p.PID == 0 { //已经是一级
				return
			}
			res = append(res, p.PID) //追加父级节点id
			parents := getParentsByIDByRecursion(tgs, p.PID)
			res = append(res, parents...)
			// fmt.Println(111, p.PID, res)
		}
	}
	return
}

//根据父节点id，获取所有子节点id
func getChildsByPIDByRecursion(tgs []*Tag, pid int) (res []int) {
	for _, p := range tgs {
		if p.PID == pid {
			res = append(res, p.ID)
			chids := getChildsByPIDByRecursion(tgs, p.ID)
			res = append(res, chids...)
			fmt.Println(111, p.ID, res)
		}
	}
	return
}

func getTreesByIteration(tgs []*Tag) (res []*Tag) {
	res = []*Tag{}

	tgMap := make(map[int]*Tag)
	for _, v := range tgs {
		if v == nil {
			return
		}
		tgMap[v.ID] = v

		t, ok := tgMap[v.ID]
		if !ok || t == nil {
			continue
		}

		if p, ok := tgMap[v.PID]; ok && p != nil {
			p.Children = append(p.Children, t)
		} else {
			res = append(res, t)
		}
	}

	return
}

func getTreesByRecursion(tgs []*Tag, pid int) (res []*Tag) {
	res = []*Tag{}
	for _, v := range tgs {
		if v.PID == pid {
			node := new(Tag)
			*node = *v
			node.Children = getTreesByRecursion(tgs, v.ID)
			res = append(res, node)

			// v.Children = getTreesByRecursion(tgs, v.ID)
			// res = append(res, v)
		}
	}
	return res
}
