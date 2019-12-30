use std::fmt::Display;
use std::fmt::Debug;

#[derive(Clone, Debug, std::cmp::PartialEq)]
struct TreeNode<T> {
    pub left: Option<Box<TreeNode<T>>>,
    pub right: Option<Box<TreeNode<T>>>,
    pub value: T,
}

#[derive(Clone, Debug)]
struct BinaryTree<T> {
    root: TreeNode<T>,
}

impl<T: Display> BinaryTree<T> {
    pub fn print_in_order(&self, root: &TreeNode<T>) {
        if let Some(left) = &root.left {
            self.print_in_order(&*left);
        }
        println!("{}", root.value);
        if let Some(right) = &root.right {
            self.print_in_order(&*right);
        }
    }
}
impl<T> TreeNode<T> {
    fn new(val: T) -> Self {
        return TreeNode{value: val, left: None, right: None};
    }
}
fn main() {
    let mut tree = BinaryTree{root: TreeNode::new(9)};
    tree.root.left = Some(Box::new(TreeNode::new(3)));
    tree.root.right = Some(Box::new(TreeNode::new(17)));
    tree.print_in_order(&tree.clone().root);
}
