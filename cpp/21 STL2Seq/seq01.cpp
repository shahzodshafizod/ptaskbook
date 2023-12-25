#include <iostream>
#include <fstream>
#include <vector>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	istream_iterator<int> first(file), last;
	vector<int> v(first, last);
	copy(first, last, v.begin());
	for_each(v.begin(), v.end(), [] (int n) {cout << n << '\t';});
	file.close();
	return 0;
}