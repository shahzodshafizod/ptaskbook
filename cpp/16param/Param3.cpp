#include <iostream>
using namespace std;

void MinmaxNum(double A[], int N, int& NMin, int& NMax);

int main()
{
	//Task("Param3");
	int N, NMin, NMax;
	double* mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		MinmaxNum(mass, N, NMin, NMax);
		cout << NMin << NMax;
	}
	
	return 0;
}

void MinmaxNum(double A[], int N, int& NMin, int& NMax)
{
	double minimal, maximal;
	minimal = maximal = A[0];
	NMin = NMax = 1;
	for (int i = 1; i < N; i++)
	{
		if (A[i] > maximal)
		{
			maximal = A[i];
			NMax = i+1;
		}
		if (A[i] < minimal)
		{
			minimal = A[i];
			NMin = i+1;
		}
	}
}