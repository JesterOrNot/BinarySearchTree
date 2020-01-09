#include <iostream>
using namespace std;
struct Node {
    struct Node* left;
    struct Node* right;
    double value;
};
struct BinaryTree {
    struct Node* root;
};
struct BinaryTree* newBinaryTree() {
    struct BinaryTree* theTree = new(struct BinaryTree);
    theTree->root = NULL;
    return(theTree);
}
struct Node* newNode(double value) {
    struct Node* node = new(struct Node);
    node->left = NULL;
    node->right = NULL;
    node->value = value;
    return(node);
};
void printInOrder(struct Node* node) {
    if(node != NULL) {
        printInOrder(node->left);
        cout << node->value << ", ";
        printInOrder(node->right);
    }
}
void printPreOrder(struct Node* node) {
    if (node != NULL) {
        cout << node->value << ", ";
        printPreOrder(node->left);
        printPreOrder(node->right);
    }
}
void printPostOrder(struct Node* node) {
    if (node != NULL) {
        printPostOrder(node->left);
        printPostOrder(node->right);
        cout << node->value << ", ";
    }
}
void printOutOrder(struct Node* node) {
    if (node != NULL) {
        printOutOrder(node->right);
        cout << node->value << ", ";
        printOutOrder(node->left);
    }
}
bool findNode(double target, struct BinaryTree* tree, struct Node* current) {
    if (target==current->value) {
        return(true);
    } else if (target < current->value) {
        if(current->left!=NULL) {
            current = current->left;
        } else {
            return(false);
        }
    } else if (target >= current->value) {
        if(current->right!=NULL) {
            current = current->right;
        } else {
            return(false);
        }
    }
    return findNode(target, tree, current);
}
struct Node* insertNode(struct Node* node, struct Node* root) {
    if (root == NULL) {
        return node;
    } else if (node->value < root->value) {
        root->left = insertNode(node, root->left);
    } else {
        root->right = insertNode(node, root->right);
    }
    return(root);
}
struct BinaryTree* assembleTree(double theArr[]) {
    struct BinaryTree* tree;
    tree->root = newNode(theArr[0]);
    int size = *(&theArr+1) - theArr;
    int index = 0;
    while(index <= size - 2) {
        insertNode(newNode(theArr[index+1]),tree->root);
        index++;
    }
    return(tree);
}
int main() {
    double theList[] = {10,4,8.9,2,11,33,5,1,3,9,66,999,0.5,-6,-67,-6.8};
    cout << endl;
    struct BinaryTree* tree = assembleTree(theList);
    printInOrder(tree->root);
    cout << findNode(5, tree, tree->root);
}
