#include <iostream>
using namespace std;

void SortAfz(double& a, double& b);

int main()
{
	//Task("Minmax22");
	int N;
	cin >> N;

	double number, minimal1, minimal2;
	
	cin >> number;
	minimal1 = number;
	
	cin >> number;
	minimal2 = number;
	SortAfz(minimal1, minimal2);

	for (int i = 3; i <= N; i++)
	{
		cin >> number;

		if (number < minimal2)
			minimal2 = number;
		
		SortAfz(minimal1, minimal2);
	}

	cout << minimal1 << minimal2;
	
	return 0;
}

void SortAfz(double& a, double& b)
{
	if (a > b)
	{
		double temp = a;
		a = b;
		b = temp;
	}
}