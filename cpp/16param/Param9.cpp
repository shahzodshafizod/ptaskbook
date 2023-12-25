#include <iostream>
using namespace std;

void RemoveForInc(double** A, int& N);

int main()
{
	//Task("Param9");
	int N;
	double* mass = NULL;
	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		RemoveForInc(&mass, N);
		
		cout << N;
		for (int j = 0; j < N; j++)
			cout << mass[j];

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void RemoveForInc(double** A, int& N)
{
	for (int i = 1; i < N; i++)
	{
		if ((*A)[i] < (*A)[i-1])
		{
			double* temp = new double [N-1];
			for (int k = 0, j = 0; k < N; k++)
			{
				if (k == i)
					continue;
				temp[j++] = (*A)[k];
			}
			delete [] (*A);
			(*A) = temp;
			temp = NULL;
			N--;
			i--;
		}
	}
}