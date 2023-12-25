#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

bool comp (int a, int b) {
	return a > b;
}

int main() {
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream file1(name1);
	ifstream file2(name2);
	merge(istream_iterator<int>(file1), istream_iterator<int>(), istream_iterator<int>(file2), istream_iterator<int>(),
		ostream_iterator<int>(cout, "\t"), comp);
	file1.close();
	file2.close();
	return 0;
}