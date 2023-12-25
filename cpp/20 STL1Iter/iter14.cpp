#include <iostream>
#include <iterator>
using namespace std;

int main() {
	ostream_iterator<double> out(cout, " ");
	int i = 0;
	for (istream_iterator<double> in(cin); in != istream_iterator<double>(); ) {
		++i;
		if (i % 2 == 0) {
			out = *in;
		}
		++in;
	}
	return 0;
}