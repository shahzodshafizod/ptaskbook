#include <iostream>
using namespace std;

int main()
{
	//Task("If5");
	int a, b, c;
	cin >> a >> b >> c;
	int positives = 0, negatives = 0;
	if (a > 0)
		positives++;
	else if (a < 0)
		negatives++;
	if (b > 0)
		positives++;
	else if (b < 0)
		negatives++;
	if (c > 0)
		positives++;
	else if (c < 0)
		negatives++;
	
	cout << positives << negatives;
	
	return 0;
}
