#include <stdio.h>
#include "TNode.h"
#include <stdlib.h>

int main() {
	struct TNode* P1 = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(P1);
	printf("P1->Data = ");
	scanf("%d", &(P1->Data));
	
	P1->Left = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(P1->Left);
	printf("P1->Left->Data = ");
	scanf("%d", &(P1->Left->Data));
	
	P1->Right = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(P1->Right);
	printf("P1->Right->Data = ");
	scanf("%d", &(P1->Right->Data));
	
	printf("\nP1->Data = %d", P1->Data);
	printf("\nP1->Left->Data = %d", P1->Left->Data);
	printf("\nP1->Right->Data = %d", P1->Right->Data);
	printf("\nP1->Left = %p", P1->Left);
	printf("\nP1->Right = %p", P1->Right);
	return 0;
}