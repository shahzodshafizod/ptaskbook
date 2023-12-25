#include <iostream>
using namespace std;

int main()
{
	//Task("For6");
	double price;
	cin >> price;
	for (double u = 1.2; u <= 2.0; u += 0.2)
		cout << u*price;
	
	return 0;
}
