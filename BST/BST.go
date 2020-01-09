package main
import "fmt"
type Node struct {
    left* Node
    right* Node
    value float64
}
type BinaryTree struct {
    root *Node
}
func trimString(text string,count int) string {
    text = text[:len(text)-count]
    return text
}
func newNode(data float64) *Node {
    node := new(Node)
    node.left = nil
    node.right = nil
    node.value = data
    return node
}
func newBinaryTree() *BinaryTree {
    tree := new(BinaryTree)
    tree.root = nil
    return tree
}
func printInOrder(node* Node) {
    if node == nil {
        return
    }
    printInOrder(node.left)
    fmt.Print(node.value,", ")
    printInOrder(node.right)
}
func printPreOrder(node* Node) {
    if node == nil {
        return
    }
    fmt.Print(node.value,", ")
    printPreOrder(node.left)
    printPreOrder(node.right)
}
func printPostOrder(node* Node) {
    if node == nil {
        return
    }
    printPostOrder(node.left)
    printPostOrder(node.right)
    fmt.Print(node.value,", ")
}
func printOutOrder(node* Node) {
    if node == nil {
        return
    }
    printOutOrder(node.right)
    fmt.Print(node.value,", ")
    printOutOrder(node.left)
}
func prettyPrintOutOrder(node* Node) {
    printOutOrder(node)
    println("")
}
func prettyPrintPreOrder(node* Node) {
    printPreOrder(node)
    println("")
}
func prettyPrintPostOrder(node* Node) {
    printPostOrder(node)
    println("")
}
func prettyPrintInOrder(node* Node) {
    printInOrder(node)
    println("")
}
func findNode(target float64, theTree BinaryTree, current* Node) bool {
    if target == current.value {
        return true
    } else if target < current.value {
        if current.left != nil {
            current = current.left
        } else {
            return false
        }
    } else if target >= current.value {
        if current.right != nil {
            current = current.right
        } else {
            return false
        }
    }
    return findNode(target,theTree,current)
}
func insertNode(node* Node, root* Node) *Node {
    if root == nil {
        return node
    } else if node.value < root.value {
        root.left = insertNode(node, root.left)
    } else {
        root.right = insertNode(node, root.right)
    }
    return root
}
func assembleTree(theList []float64) *BinaryTree {
    tree := newBinaryTree()
    tree.root = newNode(theList[0])
    var index int = 1
    for index<=len(theList)-1 {
        insertNode(newNode(theList[index]), tree.root)
        index++;
    }
    return tree;
}
func main() {
    var theList = []float64{10,4,8.9,2,11,33,5,1,3,66,999,0.5,-6,-67,-6.8}
    tree := assembleTree(theList)
    prettyPrintInOrder(tree.root)
    fmt.Println(findNode(5, *tree, tree.root))
    prettyPrintOutOrder(tree.root)
}
