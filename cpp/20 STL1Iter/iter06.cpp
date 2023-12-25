#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	ofstream file(name);
	copy(istream_iterator<char>(cin), istream_iterator<char>(), ostream_iterator<char>(file, " "));
	file.close();
	return 0;
}