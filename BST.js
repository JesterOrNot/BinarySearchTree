class Node {
    constructor(value) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}
class BinaryTree {
    constructor() {
        this.root = null;
    }
}
function printInOrder(node) {
    if(node!==null) {
        printInOrder(node.left);
        console.log(node.value + ", ");
        printInOrder(node.right);
    }
}
function printPreOrder(node) {
    if(node!==null) {
        console.log(node.value + ", ");
        printPreOrder(node.left);
        printPreOrder(node.right);
    }
}
function printPostOrder(node) {
    if(node!==null) {
        printPostOrder(node.left);
        printPostOrder(node.right);
        console.log(node.value + ", ");
    }
}
function findNode(target,tree,current) {
    if (target == current.value) {
        reutrn(true);
    } else if (target < current.value) {
        if(current.left!==null) {
            current = current.left;
        } else {
            return(false);
        }
    } else if(current>=current.value) {
        if(current.right!==null) {
            current = current.right;
        } else {
            return(false);
        }
    }
}
function insertNode(node,root) {
    if (root === null) {
        return node;
    } else if (node.value < root.value) {
        root.left = insertNode(node, root.left);
    } else {
        root.right = insertNode(node, root.right);
    }
    return root;
}
function assembleTree(theList) {
    var tree = new BinaryTree();
    tree.root = new Node(theList[0]);
    index = 1;
    while(index<=theList.length-1) {
        insertNode(new Node(theList[index]), tree.root);
        index++;
    }
    return(tree);
}
var theTree = assembleTree([1,2,3,0.5]);
printInOrder(theTree.root);