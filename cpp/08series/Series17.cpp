#include <iostream>
using namespace std;

int main()
{
	//Task("Series17");
	double B;
	cin >> B;
	int N;
	cin >> N;
	double number;
	bool outedB = false;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		if ( (outedB == false) && (B < number) )
		{
			cout << B;
			outedB = true;
		}
		cout << number;
	}
	if (outedB == false)
		cout << B;
	
	return 0;
}
