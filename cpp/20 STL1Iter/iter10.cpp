#include <iostream>
#include <iterator>
#include <algorithm>
using namespace std;

class f {
public:
	int index;
	f(int _index) {
		index = _index;
	}
	bool operator() (double e) {
		return ++index%2 != 0;
	}
};

int main() {
	remove_copy_if(istream_iterator<double>(cin), istream_iterator<double>(), ostream_iterator<double>(cout, "\t"), f(0));
	return 0;
}