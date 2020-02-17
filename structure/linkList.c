#include <stdio.h>
#include <stdlib.h>

typedef struct LinkListNode {
    int value;
    struct LinkListNode *next;
} LinkListNode;

// 初始化链表
LinkListNode *init() {
    LinkListNode *head = (LinkListNode *)malloc(sizeof(LinkListNode));
    head->value = 0;
    head->next = NULL;
    return head;
}

// 链表插入
LinkListNode *insert(LinkListNode *l, int value, int postion) {
    LinkListNode *node = l;
    // 特殊处理 postion 为 0 的情况
    if (postion == 0) {
        LinkListNode *start = (LinkListNode *)malloc(sizeof(LinkListNode));
        start->value = value;
        start->next = node;
        return start;
    }

    // 其他情况
    int i = 0;
    while (i < postion-1) {
        if ((node = node->next) == NULL) {
            printf("invalid postion: %d\n", postion);
            return l;
        }    
    }
    LinkListNode *new = (LinkListNode *)malloc(sizeof(LinkListNode));
    new->value = value;
    LinkListNode *temp = node->next;
    node->next = new;
    if (temp != NULL)
        new->next = temp;
    return l;
}

// 打印链表
void showLinkList(LinkListNode *l) {
    LinkListNode *node = l;
    printf("\nprint current LinkListNode:\n");
    printf("{");
    int start = 0;
    while (node != NULL) {
        if (start == 0) {
            start = 1;
        } else {
            printf(", ");
        }
        printf("%d", node->value);
        node = node->next;
    }
    printf("}\n");
}

int main() {
    LinkListNode *linkList = init();
    printf("init a LinkListNode, value: %d\n", linkList->value); 
    linkList = insert(linkList, 1, 0);
    linkList = insert(linkList, 2, 0);
    linkList = insert(linkList, 0, 0);
    linkList = insert(linkList, -1, 1);
    showLinkList(linkList);
}




