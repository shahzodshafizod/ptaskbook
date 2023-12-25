#include <iostream>
using namespace std;

int main()
{
	//Task("Array32");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	int index = 0;
	if (mass[0] < mass[1])
	{
		index = 1;
		cout << index;
	}
	else
	{
		for (int i = 1; i < N-1; i++)
			if ( (mass[i] < mass[i-1]) && (mass[i] < mass[i+1]) )
			{
				index = i+1;
				break;
			}
		
		if (index > 0)
			cout << index;
		
		else if (mass[N-1] < mass[N-2])
		{
			index = N;
			cout << index;
		}
		
		else
			cout << index;
	}
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
