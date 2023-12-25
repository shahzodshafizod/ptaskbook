#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

int main() {
	char name[50];
	cin >> name;
	int K;
	cin >> K;
	ofstream file(name);
	fill_n(ostream_iterator<char>(file), K, '*');
	file.close();
	return 0;
}