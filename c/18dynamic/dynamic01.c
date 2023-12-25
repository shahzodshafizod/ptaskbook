#include <stdio.h>
#include "TNode.h"
#include <stdlib.h>

int main() {
	struct TNode* P1 = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(P1);
	printf("P1->Data = ");
	scanf("%d", &(P1->Data));
	
	P1->Next = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(P1->Next);
	printf("P1->Next->Data = ");
	scanf("%d", &(P1->Next->Data));
	
	printf("\nP1->Data = %d", P1->Data);
	printf("\nP1->Next->Data = %d", P1->Next->Data);
	printf("\nP1->Next = %p", P1->Next);
	
	free(P1->Next);
	free(P1);
	return 0;
}