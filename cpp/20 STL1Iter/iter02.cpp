#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

bool isPositive(double d) {
	return d > 0;
}

int main() {
	char name[50];
	cin >> name;
	ifstream file(name);
	//cout << count_if(istream_iterator<double>(file), istream_iterator<double>(), [](double d) {return d > 0;});
	cout << count_if(istream_iterator<double>(file), istream_iterator<double>(), isPositive);
	file.close();
	return 0;
}