#include <iostream>
#include <deque>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	istream_iterator<int> first(cin), last;
	ostream_iterator<int> out(cout, "\t");
	deque<int> d(first, last);
	copy(first, last, d.begin()); //copy from console to deque
	copy(d.rbegin()+d.size()/2, d.rend(), out);
	cout << endl;
	copy(d.rbegin(), d.rbegin()+d.size()/2, out);
	return 0;
}