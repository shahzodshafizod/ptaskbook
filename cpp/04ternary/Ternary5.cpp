#include <iostream>
using namespace std;

int main()
{
	//Task("If5");
	int a, b, c;
	cin >> a >> b >> c;
	int positives = 0, negatives = 0;
	
	positives += a > 0 ? 1 : 0;
	positives += b > 0 ? 1 : 0;
	positives += c > 0 ? 1 : 0;
	
	negatives += a < 0 ? 1 : 0;
	negatives += b < 0 ? 1 : 0;
	negatives += c < 0 ? 1 : 0;
	
	cout << positives << negatives;
	
	return 0;
}
