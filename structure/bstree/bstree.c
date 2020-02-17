/*
 * 二叉查找树的定义和实现
 * */

#include <stdio.h>
#include <stdlib.h>
#include "bstree.h"

// 创建节点
static Node* create_bstree_node(Type key, Node *parent, Node *left, Node *right)
{
    Node *p;

    // malloc 失败的情况
    if ((p = (Node *)malloc(sizeof(Node))) == NULL)
        return NULL;
    p->key = key;
    p->left = left;
    p->right = right;
    p->parent = parent;
    
    return p;
}

// 前序遍历
// 访问根节点，前序遍历左子树，前序遍历右子树
void preorder_bstree(BSTree tree)
{
    if (tree != NULL)
    {
        printf("%d ", tree->key);
        preorder_bstree(tree->left);
        preorder_bstree(tree->right);
    }
}

// 中序遍历
// 左，根，右
void inorder_bstree(BSTree tree)
{
    if (tree != NULL)
    {
        inorder_bstree(tree->left);
        printf("%d ", tree->key);
        inorder_bstree(tree->right);
    }
}


// 后序遍历
// 左，右，根
void postorder_bstree(BSTree tree)
{
    if (tree != NULL)
    {
        postorder_bstree(tree->left);
        postorder_bstree(tree->right);
        printf("%d ", tree->key);
    }
}


// 递归版本的查找
Node* bstree_search(BSTree x, Type key)
{
    if (x == NULL || x->key == key)
        return x;

    if (key < x->key)
        return bstree_search(x->left, key);
    else
        return bstree_search(x->right, key);
}

// 非递归版本的查找
Node* iterative_bstree_search(BSTree x, Type key)
{
    while ((x != NULL) && (x->key != key))
    {
        if (key < x->key)
            x = x->left;
        else
            x = x->right;
    }

    return x;
}

// 查找最大值
// 最右的节点
Node* bstree_maximun(BSTree tree)
{
    if (tree == NULL)
        return NULL;

    while(tree->right != NULL)
        tree = tree->right;

    return tree;
}

// 查找最小值
Node* bstree_minimum(BSTree tree)
{
    if (tree == NULL)
        return NULL;

    while (tree->left != NULL) {
        tree = tree->left;
    }
    
    return tree;
}


// 查找节点的前驱
// 即该节点的左子树中的最大节点
Node* bstree_predecessor(Node *x)
{
    if (x->left != NULL)
        return bstree_maximun(x->left);

    Node *y = x->parent;
    while ((y != NULL) && (x == y->left))
    {
        x = y;
        y = y->parent;
    }

    return y;
}


// 查找节点的后继
// 即该节点右子树中的最小节点
Node* bstree_successor(Node *x)
{
    if (x->right != NULL)
        return bstree_minimum(x->right);
    
    Node *y = x->parent;
    while ((y != NULL) && (x == y->right))
    {
        x = y;
        y = y->parent;
    }

    return y;
}


// 节点插入
// 内部函数
static Node* bstree_insert(BSTree tree, Node *z)
{
    Node *y = NULL;
    Node *x = tree;

    // 查找 z 的插入位置
    while (x != NULL)
    {
        y = x;
        if (z->key < x->key)
            x = x->left;
        else
            x = x->right;
    }

    z->parent = y;
    if (y == NULL)
        tree = z;
    else if (z->key < y->key)
        y->left = z;
    else
        y->right = z;
    
    return tree;
}

// 对外的函数
Node *insert_bstree(BSTree tree, Type key)
{
    Node *z;

    if ((z == create_bstree_node(key, NULL, NULL, NULL)) == NULL)
        return tree;

    return bstree_insert(tree, z);
}


// 删除节点
// 内部函数
static Node* bstree_delete(BSTree tree, Node *z)
{
    Node *x = NULL;
    Node *y = NULL;

    if ((z->left == NULL) || (z->right == NULL))
        y = z;
    else
        y = bstree_successor(z);

    if (y->left != NULL)
        x = y->left;
    else
        x = y->right;

    if (x != NULL)
        x->parent = y->parent;

    if (y->parent == NULL)
        tree = x;
    else if (y == y->parent->left)
        y->parent->left = x;
    else
        y->parent->right = x;

    if (y != z)
        z->key = y->key;

    if (y != NULL)
        free(y);

    return tree;
}

// 对外暴露的接口函数
Node* delete_bstree(BSTree tree, Type key)
{
    Node *z, *node;

    if ((z = bstree_search(tree, key)) != NULL)
        tree = bstree_delete(tree, z);

    return tree;
}


// 打印二叉树
void print_bstree(BSTree tree, Type key, int direction)
{
    if (tree != NULL)
    {
        if (direction == 0)
            printf("%2d is root\n", tree->key);
        else
            printf("%2d is %2d's %6s child\n", tree->key, key, direction == 1 ? "right" : "left");

        print_bstree(tree->left, tree->key, -1);
        print_bstree(tree->right, tree->key, 1);
    }
}

// 销毁二叉树
void destory_bstree(BSTree tree)
{
    if (tree == NULL)
        return;
    
    if (tree->left != NULL)
        destory_bstree(tree->left);
    if (tree->left != NULL)
        destory_bstree(tree->right);

    free(tree);
}

















































