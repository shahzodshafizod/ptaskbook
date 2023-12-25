#include <iostream>
#include <vector>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	istream_iterator<int> first(cin), last;
	ostream_iterator<int> out(cout, "\t");
	vector<int> v(first, last);
	copy(first, last, v.begin());
	copy(v.begin() + v.size()/2, v.end(), out);
	cout << endl;
	copy(v.begin(), v.begin() + v.size()/2, out);
	return 0;
}