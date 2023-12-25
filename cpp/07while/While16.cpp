#include <iostream>
using namespace std;

int main()
{
	//Task("While16");
	double P, length = 10;
	cin >> P;
	int day = 1;
	double sum = length;
	while (sum <= 200)
	{
		length += length * P / 100;
		sum += length;
		day++;
	}
	cout << day << sum;
	
	return 0;
}
