/**
 * @author Sean Hellum
 * <p>This is a BST library</p>*/
import java.util.List;
import java.util.Arrays;
public class BST {
    public static class BinaryTree {
        Node root;
        BinaryTree(double value) {
            root = new Node(value);
        }
        BinaryTree() {
            root = null;
        }
    }
    public static class Node {
        double value;
        Node left, right;
        public Node(double item) {
            value = item;
            left = null;
            right = null;
        }
    }
    /**
     * @param node The root node of the tree i.e. printInorder(tree.root);*/
    public static void printInorder(Node node) {
        if (node == null) {
            return;
        }
        printInorder(node.left);
        System.out.print(node.value + ",");
        printInorder(node.right);
    }
    public static void printPreorder(Node node) {
        if (node == null) {
            return;
        }
        System.out.print(node.value + ",");
        printPreorder(node.left);
        printPreorder(node.right);
    }
    public static void deleteTree(Node node) {
        node=null;
    }
    public static void printPostorder(Node node) {
        if (node == null) {
            return;
        }
        printPostorder(node.left);
        printPostorder(node.right);
        System.out.print(node.value + ",");
    }
    public static void printOutOrder(Node node) {
        if (node == null) {
            return;
        }
        printOutOrder(node.right);
        System.out.print(node.value + ",");
        printOutOrder(node.left);
    }
    public static void prettyPrintInOrder(Node node) {
        printInorder(node);
        System.out.println();
    }
    public static void prettyPrintPreOrder(Node node) {
        printPreorder(node);
        System.out.println();
    }
    public static void prettyPrintPostOrder(Node node) {
        printPostorder(node);
        System.out.println();
    }
    public static void prettyPrintOutOrder(Node node) {
        printOutOrder(node);
    }
    public static boolean findNode(double target,Node current) {
        if (target == current.value) {
            return true;
        } else if (target < current.value) {
            if(current.left!=null) {
                current = current.left;
            } else {
                return false;
            }
        } else if (target >= current.value) {
            if (current.right!=null) {
                current = current.right;
            } else {
                return false;
            }
        }
        return findNode(target, current);
    }
    public static Node insertNode(Node node, Node root) {
        if (root == null) {
           return node;
        } else if (node.value < root.value) {
            root.left = insertNode(node, root.left);
        } else {
            root.right = insertNode(node, root.right);
        }
        return root;
    }
    public static BinaryTree assembleTree(List<Double> list) {
        BinaryTree tree = new BinaryTree();
        tree.root = new Node(list.get(0));
        int index = 1;
        while(index<=list.size()-1) {
            insertNode(new Node(list.get(index)), tree.root);
            index++;
        }
        return tree;
    }
}