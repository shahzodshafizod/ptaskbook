#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	copy(istream_iterator<int>(file), istream_iterator<int>(), ostream_iterator<int>(cout, "\t"));
	file.close();
	return 0;
}