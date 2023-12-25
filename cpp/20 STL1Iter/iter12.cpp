#include <iostream>
#include <fstream>
#include <iterator>
using namespace std;

int main() {
	int K;
	cin >> K;
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream fin(name1);
	ofstream fout(name2);
	ostream_iterator<string> out(fout, "\n");
	for (istream_iterator<string> in(fin); in != istream_iterator<string>(); ) {
		if ((*in).length() <= K) {
			out = *in;
		}
		++in;
	}
	fin.close();
	fout.close();
	return 0;
}