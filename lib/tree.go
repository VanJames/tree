// Copyright (c) 2018 VanJames(fanxu990516@gmail.com)

package lib

//树
type Tree struct {
	List map[int]*Node
	Children map[int]Node
	Parents map[int]Node
}

//节点
type Node struct {
	Id int `json:"id"`
	Pid int `json:"pid"`
	Name string `json:"name"`
	Child []*Node `json:"child"`
}


//将原始数据创建树结构
func (this *Tree) BuildTree(nodes []Node) {
	this.List = make(map[int]*Node,0)
	for index,_ := range nodes {
		id := nodes[index].Id
		nodes[index].Child = make([]*Node,0)
		this.List[id] = &nodes[index]
	}
	for k,_ := range this.List {
		pid := this.List[k].Pid
		if _,ok := this.List[pid];ok {
			this.List[pid].Child = append(this.List[pid].Child,this.List[k])
		}
	}
}

//获取所有的parents
func (this *Tree) GetAllParents(pid int) []*Node{
	parents := make([]*Node,0)
	for k,_ := range this.List {
		if pid == this.List[k].Id {
			parents = append(parents,this.List[k])
			parents = append(parents,this.GetAllParents(this.List[k].Pid)...)
		}
	}
	this.Parents = UniqueNodes(parents)
	return parents
}

//获取所有的children
func (this *Tree) GetAllChildren(id int) ([]*Node){
	children := make([]*Node,0)
	for k,_ := range this.List {
		if id == this.List[k].Pid{
			if _,ok := this.Children[this.List[k].Id];!ok{
				children = append(children,this.List[k])
				children = append(children,this.GetAllChildren(this.List[k].Id)...)
			}
		}
	}
	this.Children = UniqueNodes(children)
	return children
}

//获取当前节点
func (this *Tree) GetNode(id int) (*Node){
	if _,ok := this.List[id];ok {
		return this.List[id]
	}
	return nil
}

//nodes 节点排重
func UniqueNodes( nodes []*Node ) map[int]Node {
	newArr := make(map[int]Node,0)
	for k,_ := range nodes {
		node := nodes[k]
		if _,ok := newArr[node.Id];!ok{
			newArr[node.Id] = Node{node.Id,node.Pid,node.Name,[]*Node{}}
		}
	}
	return newArr
}

//只拿出包含节点的ID
func (this *Tree)GetIds( nodes map[int]Node ) []int {
	newArr := make([]int,0)
	for id,_ := range nodes {
		newArr = append(newArr,id)
	}
	return newArr
}