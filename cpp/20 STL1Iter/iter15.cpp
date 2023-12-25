#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	ofstream file(name);
	replace_copy(istream_iterator<int>(cin), istream_iterator<int>(), ostream_iterator<int>(file, "\t"), 0, 10);
	file.close();
	return 0;
}