#include <iostream>
#include <iterator>
#include <algorithm>
using namespace std;

bool isDigit(char C) {
	return C >= '0' && C <= '9';
}

int main() {
	replace_copy_if(istream_iterator<char>(cin), istream_iterator<char>(), ostream_iterator<char>(cout, " "), isDigit, '_');
	return 0;
}