#include <iostream>
using namespace std;

int MaxNum(double A[], int N);

int main()
{
	//Task("Param2");
	int N;
	double* mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new double [N];
		
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		cout << MaxNum(mass, N);
		
		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

int MaxNum(double A[], int N)
{
	double maximal = A[0];
	int index = 0;
	for (int i = 1; i < N; i++)
		if (A[i] > maximal)
		{
			maximal = A[i];
			index = i;
		}
	
	return (index+1);
}