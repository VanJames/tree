package main

import (
	"encoding/json"
	"log"
	"github.com/VanJames/tree/lib"
)

func main()  {
	//原始数据格式 目前支持转成该种方式 [{"id":1,"name":"test","pid":0},{"id":2,"name":"test1","pid":1},{"id":3,"name":"test2","pid":2}]
	menus := []byte(`[{"id":1,"name":"test","pid":0},{"id":2,"name":"test1","pid":1},{"id":3,"name":"test2","pid":2}]`)
	var nodes []lib.Node
	err := json.Unmarshal(menus,&nodes)
	if err != nil{
		log.Fatal("JSON decode error:",err)
		return
	}
	//构建树
	var exampleTree lib.Tree
	exampleTree.BuildTree(nodes)
	bs,_ := json.Marshal(exampleTree.List)
	log.Println("tree:",string(bs))
	//通过pid 查询所有父节点
	exampleTree.GetAllParents(2)
	bs,_ = json.Marshal(exampleTree.Parents)
	log.Println("parents:",string(bs))
	log.Println("ids:",exampleTree.GetIds(exampleTree.Parents))
	//通过id 查询所有子节点
	exampleTree.GetAllChildren(1)
	bs,_ = json.Marshal(exampleTree.Children)
	log.Println("childs:",string(bs))
	log.Println("ids:",exampleTree.GetIds(exampleTree.Children))
}
