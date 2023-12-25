#include <iostream>
#include <fstream>
#include <iterator>
#include <algorithm>
using namespace std;

bool comp (string str1, string str2) {
	int len1 = str1.length();
	int len2 = str2.length();
	if (len1 == len2) {
		bool ok = true;
		for (int i = 0; i < len1; ++i) {
			if (str1[i] == str2[i]) {
				continue;
			}
			if (str1[i] > str2[i]) {
				ok = false;
			}
			break;
		}
		return ok;
	} else {
		return len1 < len2;
	}
}

int main() {
	char name1[50], name2[50];
	cin >> name1 >> name2;
	ifstream file1(name1);
	ifstream file2(name2);
	merge(istream_iterator<string>(file1), istream_iterator<string>(), istream_iterator<string>(file2), istream_iterator<string>(),
		ostream_iterator<string>(cout, "\t"), comp);
	file1.close();
	file2.close();
	return 0;
}