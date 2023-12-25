#include <iostream>
#include <iterator>
#include <algorithm>
#include <fstream>
using namespace std;

class F {
public:
	int index;
	char C;
	F(char _C, int _index) {
		index = _index;
		C = _C;
	}
	char operator() () {
		return C + index++;
	}
};

int main() {
	char name[50];
	cin >> name;
	int N;
	cin >> N;
	ofstream file(name);
	generate_n(ostream_iterator<char>(file), N, F('A', 0));
	file.close();
	return 0;
}