#include <iostream>
#include <algorithm>
#include <iterator>
using namespace std;

int main() {
	cout << count(istream_iterator<int>(cin), istream_iterator<int>(), 0);
	return 0;
}