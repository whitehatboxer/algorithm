#include <stdio.h>
#include <stdlib.h>

typedef struct LinkListNode {
    int val;
    struct LinkListNode *next;
} LinkListNode;

typedef struct LinkList {
   LinkListNode *head;
   LinkListNode *tail;
} LinkList;

int main() {
    LinkList *ll1, *ll2, *ll_res;
    // 初始化结构体指针
    ll1 = malloc(sizeof(LinkList));
    ll2 = malloc(sizeof(LinkList));
    ll_res = malloc(sizeof(LinkList));
    LinkListNode *ll1_start = ll1->head;
}

