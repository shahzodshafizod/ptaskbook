#include "Tree.h"
#include <time.h>
#include <stdlib.h>
#include <stdio.h>

#ifndef NULL
#define NULL 0;
#endif

void
Tree_constr(struct Tree* this) {
	Tree_init(this);
}

void
Tree_destr(struct Tree* this) {
	Tree_free(this);
}

void
Tree_automake(struct Tree* this, int count) {
	if (count <= 0) {
		this->current = this->root;
		return;
	}
	if (this->root == NULL) {
		this->root = (struct TNode*)malloc(sizeof(struct TNode));
		TNode_constr(this->root);
		this->nodeCount++;
		--count;
		this->root->Data = Tree_getRandomData(1, 100);
		this->current = this->root;
		if (count <= 0) { return; }
	}
	int direction;
	do {
		direction = Tree_getRandomDirection();
	} while (direction == 3 && this->current->Parent == NULL);
	char exit = 0;
	switch (direction) {
		case 1:
			if (this->current->Left != NULL) { this->current = this->current->Left; exit = 1; }
			break;
		case 2:
			if (this->current->Right != NULL) { this->current = this->current->Right; exit = 1; }
			break;
		case 3:
			this->current = this->current->Parent; exit = 1;
			break;
	}
	if (exit) {
		Tree_automake(this, count);
		return;
	}
	struct TNode* newNode = (struct TNode*)malloc(sizeof(struct TNode));
	TNode_constr(newNode);
	--count;
	this->nodeCount++;
	newNode->Data = Tree_getRandomData(1, 100);
	newNode->Parent = this->current;
	switch (direction) {
		case 1: this->current->Left = newNode; break;
		case 2: this->current->Right = newNode; break;
	}
	this->current = newNode;
	Tree_automake(this, count);
}

void
Tree_free(struct Tree* this) {
	if (this->current == NULL) {
		return;
	}
	struct TNode* node = this->current;
	this->current = node->Left;
	Tree_free(this);
	this->current = node->Right;
	Tree_free(this);
	free(node);
}

void
Tree_infix(struct Tree* this) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	struct TNode* node = this->current;
	this->current = node->Left;
	Tree_infix(this);
	printf("%d\t", node->Data);
	this->current = node->Right;
	Tree_infix(this);
}

void
Tree_display(struct Tree* this) {
	if (this->root == NULL) {
		return;
	}
	if (this->field == NULL) {
		this->field = (struct Field*)malloc(sizeof(struct Field));
		Field_constr(this->field, this->root, Tree_getLevel(this)+1, Tree_getNodeCount(this, ALL, 0));
	}
	Field_display(this->field);
}

int
Tree_getNodeCount(struct Tree* this, enum Flag flag, int data) {
	if (flag == ALL) {
		return this->nodeCount;
	}
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	struct TNode* node = this->current;
	int count = 0;
	switch (flag) {
		case DATA: count += node->Data == data; break;
		case LEFT: count += node->Parent != NULL && node->Parent->Left == node; break;
		case LEAF: count += node->Left == NULL && node->Right == NULL; break;
		case RIGHT_LEAF:
			if (node->Parent != NULL && node->Parent->Right == node && node->Left == NULL && node->Right == NULL) {
				++count;
			} break;
	}
	this->current = node->Left;
	count += Tree_getNodeCount(this, flag, data);
	this->current = node->Right;
	count += Tree_getNodeCount(this, flag, data);
	return count;
}

int
Tree_getLevel(struct Tree* this) {
	if (this->level < 0) {
		Tree_setLevel(this, 0);
	}
	return this->level;
}

void
Tree_setLevel(struct Tree* this, int level) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	if (level > this->level) { this->level = level; }
	struct TNode* node = this->current;
	this->current = node->Left;
	Tree_setLevel(this, level+1);
	this->current = node->Right;
	Tree_setLevel(this, level+1);
}

int
Tree_getSum(struct Tree* this, enum Flag flag) {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	struct TNode* node = this->current;
	int sum = 0;
	switch (flag) {
		case ALL: sum += node->Data; break;
		case LEAF: sum += node->Left == NULL && node->Right == NULL ? node->Data : 0; break;
	}
	this->current = node->Left;
	sum += Tree_getSum(this, flag);
	this->current = node->Right;
	sum += Tree_getSum(this, flag);
	return sum;
}

void
Tree_toArray(struct Tree* this, int array[], int level, enum Flag flag) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	struct TNode* node = this->current;
	switch (flag) {
		case COUNT: array[level]++; break;
		case SUM: array[level] += node->Data; break;
	}
	this->current = node->Left;
	Tree_toArray(this, array, level+1, flag);
	this->current = node->Right;
	Tree_toArray(this, array, level+1, flag);
}

void
Tree_print(struct Tree* this, enum Flag flag) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	struct TNode* node = this->current;
	if (flag == PREFIX) { printf("%d\t", node->Data); }
	this->current = node->Left;
	Tree_print(this, flag);
	if (flag == INFIX) { printf("%d\t", node->Data); }
	this->current = node->Right;
	Tree_print(this, flag);
	if (flag == POSTFIX) { printf("%d\t", node->Data); }
}

void
Tree_init(struct Tree* this) {
	this->current = NULL;
	this->root = NULL;
	this->field = NULL;
	this->nodeCount = 0;
	this->level = -1;
	srand(time(0));
}

int
Tree_getRandomData(int from, int to) {
	int res = rand() % (to - from + 1) + from;
	return res;
}

int
Tree_getRandomDirection() {
	int res = rand() % 3 + 1;
	return res;
}