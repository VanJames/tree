package main

import (
	"encoding/json"
	"log"
	"github.com/VanJames/tree/lib"
)

func main()  {
	//原始数据格式 目前支持转成该种方式 [{"id":1,"name":"test","pid":0},{"id":2,"name":"test1","pid":1},{"id":3,"name":"test2","pid":2}]
	menus := []byte(`[{"id":1,"name":"test","pid":0},{"id":2,"name":"test1","pid":1},{"id":3,"name":"test2","pid":2}]`)
	TestNode(menus)
	TestCommonNode()
}
//标准格式 树生成 需要转成标准格式字段
func TestNode(menus []byte){
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

//通用格式 树结构生成 注意 key parentkey要是 int 类型 其他数据字段自定义
func TestCommonNode(){
	menus := []byte(`[{"id":1,"name":"test","parentid":0},{"id":2,"name":"test1","parentid":1},{"id":3,"name":"test2","parentid":2}]`)
	var nodes []lib.CommonNode
	err := json.Unmarshal(menus,&nodes)
	if err != nil{
		log.Fatal("JSON decode error:",err)
		return
	}
	//构建树
	var exampleTree lib.CommonTree
	exampleTree.BuildTree(nodes,"id","parentid","children")
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
	log.Println("children:",string(bs))
	log.Println("ids:",exampleTree.GetIds(exampleTree.Children))
}
