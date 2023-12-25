#include "TStack.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#ifndef NULL
#define NULL 0
#endif

void
TStack_constr(struct TStack* this) {
	this->Top = NULL;
	srand(time(0));
}

void
TStack_destr(struct TStack* this) {
	TStack_free(this);
}

void
TStack_handmake(struct TStack* this) {
	struct TNode* newNode = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(newNode);
	printf("Data = ");
	scanf("%d", &(newNode->Data));
	newNode->Next = this->Top;
	this->Top = newNode;
	printf("Continue? (y/n):\t");
	char C;
	scanf("%c%c", &C, &C);
	if (C == 'Y' || C == 'y') {
		TStack_handmake(this);
	}
}

void
TStack_automake(struct TStack* this, int count) {
	if (count <= 0) {
		return;
	}
	struct TNode* newNode = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(newNode);
	--count;
	newNode->Data = TStack_getRandomData(10, 99);
	newNode->Next = this->Top;
	this->Top = newNode;
	newNode = NULL;
	TStack_automake(this, count);
}

void
TStack_free(struct TStack* this) {
	if (this->Top == NULL) {
		return;
	}
	struct TNode* tempNode = this->Top;
	this->Top = this->Top->Next;
	free(tempNode);
}

void
TStack_Push(struct TStack* this, int D) {
	struct TNode* newNode = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(newNode);
	newNode->Data = D;
	newNode->Next = this->Top;
	this->Top = newNode;
}

int
TStack_Pop(struct TStack* this) {
	if (this->Top == NULL) {
		return 0;
	}
	struct TNode* tempNode = this->Top;
	this->Top = this->Top->Next;
	int data = tempNode->Data;
	free(tempNode);
	return data;
}

bool
TStack_IsEmpty(const struct TStack* this) {
	return this->Top == NULL;
}

int
TStack_Peek(const struct TStack* this) {
	if (this->Top == NULL) {
		return 0;
	}
	return this->Top->Data;
}

struct TNode*
TStack_getTop(const struct TStack* this) {
	return this->Top;
}

void
TStack_print(const struct TStack* this) {
	for (struct TNode* i = this->Top; i != NULL; i = i->Next) {
		printf("%d", i->Data);
		if (i->Next != NULL) {
			printf(" - ");
		}
	}
	printf("\n");
}

int
TStack_getRandomData(int from, int to) {
	int res = rand() % (to - from + 1) + from;
	return res;
}