#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int doubleCoder(char C) {
	return 2*((int)C);
}

int main() {
	char name[50];
	cin >> name;
	ofstream file(name);
	transform(istream_iterator<char>(cin), istream_iterator<char>(), ostream_iterator<int>(file, " "), doubleCoder);
	file.close();
	return 0;
}