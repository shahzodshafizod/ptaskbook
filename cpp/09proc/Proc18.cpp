#include <iostream>
using namespace std;

double CircleS(double R);

int main()
{
	//Task("Proc18");
	double R;
	for (int i = 1; i <= 3; i++)
	{
		cin >> R;
		cout << CircleS(R);
	}
	
	return 0;
}

double CircleS(double R)
{
	return 3.14 * R * R;
}