#include <iostream>
using namespace std;

int main()
{
	//Task("If4");
	int a, b, c;
	cin >> a >> b >> c;
	int positives = 0;
	if (a > 0)
		positives++;
	if (b > 0)
		positives++;
	if (c > 0)
		positives++;
	cout << positives;
	
	return 0;
}
