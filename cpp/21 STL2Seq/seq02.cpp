#include <iostream>
#include <list>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	istream_iterator<int> first(cin), last;
	ostream_iterator<int> out(cout, "\t");
	list<int> l(first, last);
	copy(first, last, l.begin());
	copy(l.begin(), l.end(), out);
	cout << endl;
	copy(l.rbegin(), l.rend(), out);
	return 0;
}