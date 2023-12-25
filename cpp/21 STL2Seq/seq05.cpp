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
	int seyak = l.size()/3;
	
	//think about advancedIter;
	
	list<int>::iterator advancedIter = l.begin();
	advance(advancedIter, seyak);
	copy(l.begin(), advancedIter, out);
	cout << endl;
	
	advancedIter = l.rbegin();
	advance(advancedIter, -(2*seyak));
	copy(l.rbegin()+seyak, advancedIter, out);
	cout << endl;
	
	advancedIter = l.rbegin();
	advance(advancedIter, -seyak);
	copy(l.rbegin(), advancedIter, out);
	return 0;
}