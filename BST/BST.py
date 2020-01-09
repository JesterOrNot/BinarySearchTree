class Node:
    def __init__(self, value):
        self.left = None
        self.right = None
        self.value = value
class BinaryTree(object):
    def __init__(self):
        super(BinaryTree,self).__init__()
        self.root = None
def printInOrder(root):
    if not root:
        return
    printInOrder(root.left)
    print(root.value,end=", ")
    printInOrder(root.right)
def printPreOrder(root):
    if not root:
        return
    print(root.value,end=", ")
    printPreOrder(root.left)
    printPreOrder(root.right)
def printPostOrder(root):
    if not root:
        return
    printPostOrder(root.left)
    printPostOrder(root.right)
    print(root.value,end=", ")
def findNode(target,tree,current):
    if target == current.value:
        return True
    elif target < current.value:
        if current.left is not None:
            current = current.left
        else:
            return False
    elif target >= current.value:
        if current.right is not None:
            current = current.right
        else:
            return False
    return findNode(target,tree,current)
def insertNode(node,root):
    if root == None:
        return node
    elif node.value < root.value:
        root.left = insertNode(node,root.left)
    else:
        root.right = insertNode(node,root.right)
    return root
def assembleTree(theList):
    tree = BinaryTree()
    tree.root = Node(theList[0])
    index = 1
    while index <= len(theList)-1:
        insertNode(Node(theList[index]),tree.root)
        index+=1
    return(tree)
tree = assembleTree([5,2,1,0,3.10,3.5,-3])
printInOrder(tree.root)
print()