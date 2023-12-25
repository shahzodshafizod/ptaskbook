#include <iostream>
#include <fstream>
#include <iterator>
using namespace std;

int main() {
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream fin(name1);
	ofstream fout(name2);
	ostream_iterator<int> out(fout, "\n");
	for(istream_iterator<int> in(fin); in != istream_iterator<int>(); ) {
		if (*in != 0) {
			out = *in;
		}
		++in;
	}
	fin.close();
	fout.close();
	return 0;
}