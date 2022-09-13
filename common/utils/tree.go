package utils

type Node[T any] struct {
	id       int64
	Pid      int64
	Name     string
	Children []*Node[T]
}

//func buildTree(menus []SysMenu) []SysMenu {
//	var trees []SysMenu
//	for _, menu := range menus {
//		if menu.Pid == 0 {
//			trees = append(trees, menu)
//		}
//	}
//
//	for k, tree := range trees {
//		var child []SysMenu
//		for _, it := range menus {
//			if it.Pid == tree.Id {
//				child = append(child, it)
//			}
//		}
//		trees[k].Children = child
//	}
//
//	return trees
//
//}
