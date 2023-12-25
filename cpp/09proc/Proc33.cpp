#include <iostream>
using namespace std;

double RadToDeg(double R);

int main()
{
	//Task("Proc33");
	double R;
	for (int i = 1; i < 6; i++)
	{
		cin >> R;
		cout << RadToDeg(R);
	}
	
	return 0;
}

double RadToDeg(double R)
{
	return 180 * R / 3.14;
}