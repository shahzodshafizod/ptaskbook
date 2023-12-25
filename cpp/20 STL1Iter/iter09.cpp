#include <iostream>
#include <iterator>
#include <fstream>
#include <algorithm>
using namespace std;

class f {
public:
	int index;
	f(int _index) {
		index = _index;
	}
	bool operator() (int e) {
		return ++index % 2 == 0;
	}
};

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	remove_copy_if(istream_iterator<int>(file), istream_iterator<int>(), ostream_iterator<int>(cout, "\t"), f(0));
	file.close();
	return 0;
}