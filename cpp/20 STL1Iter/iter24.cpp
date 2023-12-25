#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream fin(name1);
	ofstream fout(name2);
	//transform();
	fin.close();
	fout.close();
	return 0;
}