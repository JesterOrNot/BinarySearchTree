#[derive(Clone)]
struct TreeNode<T> {
    pub left: Box<TreeNode<T>>,
    pub right: Box<TreeNode<T>>,
    pub value: T,
}

#[derive(Clone)]
struct BinaryTree<T> {
    root: TreeNode<T>,
}
impl BinaryTree<i32> {
    pub fn print_in_order(&self, root: TreeNode<i32>) {
        if root.value == None {
            return
        }
        self.print_in_order(*root.left);
        println!("{}", &root.value);
        self.print_in_order(*root.right);
    }
}
impl TreeNode<i32> {
    fn new(val: i32) -> TreeNode<i32> {
        return TreeNode::<i32>{value: val, left: Box::from(TreeNode::new(0)), right: Box::from(TreeNode::new(0))};
    }
}
fn main() {
    let mut tree = BinaryTree{root: TreeNode::new(12)};
    tree.root.value = 8;
    tree.root.left.value = 3;
    tree.clone().print_in_order(tree.root);
}
