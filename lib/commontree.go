// Copyright (c) 2018 VanJames(fanxu990516@gmail.com)

package lib

//树
type CommonTree struct {
	List map[int]*CommonNode
	Children map[int]*CommonNode
	Parents map[int]*CommonNode
	Key string
	ParentKey string
	ChildKey string
}

type CommonNode map[string]interface{}
//将原始数据创建树结构
func (this *CommonTree) BuildTree(nodes []CommonNode,key string,parentKey string,childKey string) {
	this.List = make(map[int]*CommonNode,0)
	this.Key = key
	this.ParentKey = parentKey
	this.ChildKey = childKey
	for index,_ := range nodes {
		id := this.ToInt( nodes[index][this.Key] )
		nodes[index][this.ChildKey] = make([]*CommonNode,0)
		node := nodes[index]
		this.List[id] = &node
	}
	for k,_ := range this.List {
		pid := this.ToInt( (*this.List[k])[this.ParentKey] )
		if _,ok := this.List[pid];ok {
			(*this.List[pid])[this.ChildKey] = append((*this.List[pid])[this.ChildKey].([]*CommonNode),this.List[k])
		}
	}
}

//获取所有的parents
func (this *CommonTree) GetAllParents(pid int) []*CommonNode{
	parents := make([]*CommonNode,0)
	for k,_ := range this.List {
		if pid == this.ToInt( (*this.List[k])[this.Key] ) {
			parents = append(parents,this.List[k])
			parents = append(parents,this.GetAllParents(this.ToInt( (*this.List[k])[this.ParentKey] ))...)
		}
	}
	this.Parents = this.UniqueNodes(parents)
	return parents
}

//nodes 节点排重
func (this *CommonTree)UniqueNodes( nodes []*CommonNode ) map[int]*CommonNode {
	newArr := make(map[int]*CommonNode,0)
	for k,_ := range nodes {
		node := (*nodes[k])
		if _,ok := newArr[this.ToInt( node[this.Key] )];!ok{
			newArr[this.ToInt( node[this.Key] )] = &node
		}
	}
	return newArr
}

//获取所有的children
func (this *CommonTree) GetAllChildren(id int) ([]*CommonNode){
	children := make([]*CommonNode,0)
	for k,_ := range this.List {
		if id == this.ToInt( (*this.List[k])[this.ParentKey] ){
			children = append(children,this.List[k])
			children = append(children,this.GetAllChildren(this.ToInt( (*this.List[k])[this.Key] ))...)
		}
	}
	this.Children = this.UniqueNodes(children)
	return children
}

//获取当前节点
func (this *CommonTree) GetNode(id int) (*CommonNode){
	if _,ok := this.List[id];ok {
		return this.List[id]
	}
	return nil
}

//只拿出包含节点的ID
func (this *CommonTree)GetIds( nodes map[int]*CommonNode ) []int {
	newArr := make([]int,0)
	for id,_ := range nodes {
		newArr = append(newArr,id)
	}
	return newArr
}

//interface{} 转int
func (this *CommonTree)ToInt( val interface{} ) int {
	value := val.(float64)
	return int(value)
}