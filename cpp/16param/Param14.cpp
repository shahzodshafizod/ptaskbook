#include <iostream>
using namespace std;

void Split1(const double A[], int NA, double** B, int& NB, double** C, int& NC);

int main()
{
	//Task("Param14");
	int NA, NB, NC;
	double *A = NULL, *B = NULL, *C = NULL;

	cin >> NA;
	A = new double [NA];
	for (int i = 0; i < NA; i++)
		cin >> A[i];
	
	Split1(A, NA, &B, NB, &C, NC);

	cout << NB;
	for (int i = 0; i < NB; i++)
		cout << B[i];

	cout << NC;
	for (int i = 0; i < NC; i++)
		cout << C[i];

	delete [] A;
	delete [] B;
	delete [] C;
	A = B = C = NULL;

	return 0;
}

void Split1(const double A[], int NA, double** B, int& NB, double** C, int& NC)
{
	NB = NA / 2 + NA % 2;
	(*B) = new double [NB];

	NC = NA / 2;
	(*C) = new double [NC];
	
	for (int i = 0, j = 0, k = 0; i < NA; i++)
		if (i % 2 == 0)
			(*B)[j++] = A[i];
		else
			(*C)[k++] = A[i];
}