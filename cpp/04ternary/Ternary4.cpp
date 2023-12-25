#include <iostream>
using namespace std;

int main()
{
	//Task("If4");
	int a, b, c;
	cin >> a >> b >> c;
	int positives = 0;
	positives += a > 0 ? 1 : 0;
	positives += b > 0 ? 1 : 0;
	positives += c > 0 ? 1 : 0;
	cout << positives;
	
	return 0;
}
