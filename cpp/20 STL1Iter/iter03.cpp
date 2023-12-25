#include <iostream>
#include <fstream>
#include <iterator>
#include <cstring>
#include <algorithm>
using namespace std;

bool isLeng6(string s) {
	return s.length() == 6;
}

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	cout << count_if(istream_iterator<string>(file), istream_iterator<string>(), isLeng6);
	file.close();
	return 0;
}