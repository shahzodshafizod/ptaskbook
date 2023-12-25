#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree1;
	tree1.make();
	tree1.display();
	
	Tree tree2(tree1);
	tree2.display();
	return 0;
}