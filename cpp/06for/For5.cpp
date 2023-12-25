#include <iostream>
using namespace std;

int main()
{
	//Task("For5");
	double price;
	cin >> price;
	for (double i = 0.1; i <= 1.0; i += 0.1)
		cout << i*price;
	
	return 0;
}
