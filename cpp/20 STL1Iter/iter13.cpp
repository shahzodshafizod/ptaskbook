#include <iostream>
#include <fstream>
#include <iterator>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	ostream_iterator<int> out(cout, "\t");
	int i = 0;
	for (istream_iterator<int> in(file); in != istream_iterator<int>(); ) {
		++i;
		if (i % 2 != 0) {
			out = *in;
		}
		++in;
	}
	file.close();
	return 0;
}