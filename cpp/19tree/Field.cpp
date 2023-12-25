#include "Field.h"
#include <iostream>
using namespace std;

Field::Field(TNode* node, int _height, int _width)
	: cellWidth(4) {
	this->height = _height;
	this->width = _width;
	this->init(node);
}

void Field::init(TNode* node) {
	this->matrix = new TNode** [this->height];
	this->chars = new TChar* [this->height];
	for (int i = 0; i < this->height; ++i) {
		this->matrix[i] = new TNode* [this->width];
		this->chars[i] = new TChar [this->width];
		for (int j = 0; j < this->width; ++j) {
			this->matrix[i][j] = NULL;
			this->chars[i][j] = SPACE;
		}
	}
	int index = 0;
	this->fillMatrix(node, 0, index);
	this->fillChars();
}

Field::~Field() {
	for (int i = 0; i < this->height; ++i) {
		delete [] this->matrix[i];
		delete [] this->chars[i];
	}
	delete [] this->matrix;
	delete [] this->chars;
}

void Field::fillMatrix(TNode* node, int row, int& col) {
	if (node == NULL) {
		return;
	}
	fillMatrix(node->Left, row+1, col);
	this->matrix[row][col++] = node;
	fillMatrix(node->Right, row+1, col);
}

void Field::fillChars() {
	for (int i = 0; i < this->height; ++i) {
		for (int j = 0; j < this->width; ++j) {
			if (this->matrix[i][j] == NULL) {
				continue;
			}
			this->chars[i][j] = NUMBER;
			if (this->matrix[i][j]->Left != NULL) {
				for (int k = j-1; k >= 0; --k) {
					if (this->matrix[i+1][k] == NULL) {
						this->chars[i][k] = CENTER;
						continue;
					}
					this->chars[i][k] = LEFT;
					break;
				}
			}
			if (this->matrix[i][j]->Right != NULL) {
				for (++j; j < this->width; ++j) {
					if (this->matrix[i+1][j] == NULL) {
						this->chars[i][j] = CENTER;
						continue;
					}
					this->chars[i][j] = RIGHT;
					break;
				}
			}
		}
	}
}

void Field::printChars(char C, int count) {
	for (int i = 0; i < count; ++i) {
		cout << C;
	}
}

TNode* Field::getNextNode(int i, int j) {
	for (; j < this->width; ++j) {
		if (this->matrix[i][j] != NULL) {
			return this->matrix[i][j];
		}
	}
	return NULL;
}

int Field::getSpaces(int number) {
	int spaces = this->cellWidth;
	if (number < 0) {
		number *= -1;
		--spaces;
	}
	while (number > 0) {
		number /= 10;
		--spaces;
	}
	return spaces;
}

void Field::display(TNode* node) {
	int data, spaces;
	TNode* nextNode;
	for (int i = 0; i < this->height; ++i) {
		for (int j = 0; j < this->width; ++j) {
			switch (this->chars[i][j]) {
				case SPACE:
					this->printChars(' ', this->cellWidth);
					break;
				case LEFT:
					nextNode = getNextNode(i, j);
					data = nextNode->Data;
					spaces = nextNode == node ? 2 : getSpaces(data);
					this->printChars(' ', spaces/2);
					this->printChars((char)218, 1);
					this->printChars((char)196, this->cellWidth-1);
					break;
				case CENTER:
					this->printChars((char)196, this->cellWidth);
					break;
				case RIGHT:
					this->printChars((char)196, this->cellWidth-1);
					this->printChars((char)191, 1);
					this->printChars(' ', spaces/2 + spaces%2);
					break;
				case NUMBER:
					if (this->matrix[i][j]->Left == NULL) {
						data = this->matrix[i][j]->Data;
						spaces = this->matrix[i][j] == NULL ? this->cellWidth-2 : getSpaces(data);
						this->printChars(' ', spaces/2);
					}
					
					if (this->matrix[i][j] == node) { cout << "##"; } else { cout << data; }
					
					if (this->matrix[i][j]->Right == NULL) {
						this->printChars(' ', spaces/2 + spaces%2);
					}
					break;
			}
		}
		cout << endl;
	}
}