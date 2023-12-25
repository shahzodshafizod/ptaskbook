#include "Field.h"
#include <stdlib.h>
#include <stdio.h>

#ifndef NULL
#define NULL 0
#endif

void
Field_constr(struct Field* this, struct TNode* node, int height, int width) {
	this->cellWidth = 4;
	this->height = height;
	this->width = width;
	this->left = (char)218;
	this->center = (char)196;
	this->right = (char)191;
	Field_init(this, node);
}

void
Field_init(struct Field* this, struct TNode* node) {
	this->nodes = (struct TNode***)malloc(this->height * sizeof(struct TNode**));
	this->chars = (enum Flag**)malloc(this->height * sizeof(enum Flag*));
	for (int i = 0; i < this->height; ++i) {
		this->nodes[i] = (struct TNode**)malloc(this->width * sizeof(struct TNode*));
		this->chars[i] = (enum Flag*)malloc(this->width * sizeof(enum Flag));
		for (int j = 0; j < this->width; ++j) {
			this->nodes[i][j] = NULL;
			this->chars[i][j] = SPACE;
		}
	}
	int col = 0;
	Field_fillNodes(this, node, 0, &col);
	Field_fillFlags(this);
}

void
Field_fillNodes(struct Field* this, struct TNode* node, int row, int* col) {
	if (node == NULL) {
		return;
	}
	Field_fillNodes(this, node->Left, row+1, col);
	this->nodes[row][(*col)++] = node;
	Field_fillNodes(this, node->Right, row+1, col);
}

void
Field_fillFlags(struct Field* this) {
	for (int i = 0; i < this->height; ++i) {
		for (int j = 0; j < this->width; ++j) {
			if (this->nodes[i][j] == NULL) {
				continue;
			}
			this->chars[i][j] = DATA;
			if (this->nodes[i][j]->Left != NULL) {
				for (int k = j-1; k >= 0; --k) {
					if (this->nodes[i+1][k] != NULL) {
						this->chars[i][k] = LEFT;
						break;
					}
					this->chars[i][k] = CENTER;
				}
			}
			if (this->nodes[i][j]->Right != NULL) {
				for (++j; j < this->width; ++j) {
					if (this->nodes[i+1][j] != NULL) {
						this->chars[i][j] = RIGHT;
						break;
					}
					this->chars[i][j] = CENTER;
				}
			}
		}
	}
}

void
Field_display(struct Field* this) {
	struct TNode* nextNode = NULL;
	int data, spaces;
	for (int i = 0; i < this->height; ++i) {
		for (int j = 0; j < this->width; ++j) {
			switch (this->chars[i][j]) {
				case SPACE:
					Field_printChars(' ', this->cellWidth);
					break;
				case LEFT:
					nextNode = Field_getNextNode(this, i, j);
					data = nextNode->Data;
					spaces = this->cellWidth - Field_getSpaces(data);
					Field_printChars(' ', spaces/2);
					Field_printChars(this->left, 1);
					Field_printChars(this->center, this->cellWidth-1);
					break;
				case CENTER:
					Field_printChars(this->center, this->cellWidth);
					break;
				case RIGHT:
					Field_printChars(this->center, this->cellWidth-1);
					Field_printChars(this->right, 1);
					Field_printChars(' ', spaces - spaces/2);
					break;
				case DATA:
					if (this->nodes[i][j]->Left == NULL) {
						nextNode = this->nodes[i][j];
						data = nextNode->Data;
						spaces = this->cellWidth - Field_getSpaces(data);
						Field_printChars(' ', spaces/2);
					}
					printf("%d", data);
					if (this->nodes[i][j]->Right == NULL) {
						Field_printChars(' ', spaces - spaces/2);
					}
					break;
			}
		}
		nextNode = NULL;
		printf("\n");
	}
}

struct TNode*
Field_getNextNode(struct Field* this, int row, int col) {
	for (++col; col < this->width; ++col) {
		if (this->nodes[row][col] != NULL) {
			return this->nodes[row][col];
		}
	}
	return NULL;
}

void
Field_printChars(char C, int count) {
	for (int i = 0; i < count; ++i) {
		printf("%c", C);
	}
}

int
Field_getSpaces(int data) {
	int spaces = 0;
	if (data < 0) {
		data *= -1;
		++spaces;
	}
	while (data > 0) {
		data /= 10;
		++spaces;
	}
	return spaces;
}