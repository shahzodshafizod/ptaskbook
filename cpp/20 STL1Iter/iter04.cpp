#include <iostream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	copy(istream_iterator<double>(cin), istream_iterator<double>(), ostream_iterator<double>(cout, "\t"));
	return 0;
}