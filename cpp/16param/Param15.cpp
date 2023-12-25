#include <iostream>
using namespace std;

void Split2(const int A[], int NA, int** B, int& NB, int** C, int& NC);

int main()
{
	//Task("Param15");
	int NA, NB, NC;
	int *A = NULL, *B = NULL, *C = NULL;

	cin >> NA;
	A = new int [NA];
	for (int i = 0; i < NA; i++)
		cin >> A[i];

	Split2(A, NA, &B, NB, &C, NC);

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

void Split2(const int A[], int NA, int** B, int& NB, int** C, int& NC)
{
	NB = 0, NC = 0;
	for (int i = 0; i < NA; i++)
		if (A[i] % 2 == 0)
			NB++;
		else
			NC++;
	
	(*B) = new int [NB];
	(*C) = new int [NC];

	for (int i = 0, j = 0, k = 0; i < NA; i++)
		if (A[i] % 2 == 0)
			(*B)[j++] = A[i];
		else
			(*C)[k++] = A[i];

}