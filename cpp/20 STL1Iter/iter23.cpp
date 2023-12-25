#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

double getDiff (double a, double b) {
	return b - a;
}

int main() {
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream file1(name1);
	ifstream file2(name2);
	transform(istream_iterator<double>(file1), istream_iterator<double>(), istream_iterator<double>(file2),
		ostream_iterator<double>(cout, " "), getDiff);
	file1.close();
	file2.close();
	return 0;
}