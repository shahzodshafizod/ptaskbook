#include <iostream>
#include <fstream>
#include <iterator>
#include <string>
#include <algorithm>
using namespace std;

class f {
public:
	int len;
	f(int _len) {
		len = _len;
	}
	bool operator() (string e) {
		return e.length() > len;
	}
};

int main() {
	int K;
	cin >> K;
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream fin(name1);
	ofstream fout(name2);
	remove_copy_if(istream_iterator<string>(fin), istream_iterator<string>(), ostream_iterator<string>(fout, "\n"), f(K));
	fin.close();
	fout.close();
	return 0;
}