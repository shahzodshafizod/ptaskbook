#include <iostream>
#include <cstdlib>
using namespace std;

int intCmpReverse(const void* pa, const void* pb);

int main()
{
	//Task("Array64");
	
	int NA;
	cin >> NA;
	int* A = new int [NA];
	for (int i = 0; i < NA; i++)
		cin >> A[i];

	int NB;
	cin >> NB;
	int* B = new int [NB];
	for (int i = 0; i < NB; i++)
		cin >> B[i];

	int NC;
	cin >> NC;
	int* C = new int [NC];
	for (int i = 0; i < NC; i++)
		cin >> C[i];

	int ND = NA + NB + NC;
	int* D = new int [ND];
	int iD = 0;

	for (int i = 0; i < NA; i++)
		D[iD++] = A[i];

	for (int i = 0; i < NB; i++)
		D[iD++] = B[i];

	for (int i = 0; i < NC; i++)
		D[iD++] = C[i];

	qsort(D, ND, sizeof(int), intCmpReverse);

	for (int i = 0; i < ND; i++)
		cout << D[i];

	delete [] A;
	delete [] B;
	delete [] C;
	delete [] D;
	A = B = C = D = 0;
	
	return 0;
}

int intCmpReverse(const void* pa, const void* pb)
{
	const int* a = reinterpret_cast<const int*>(pa);
	const int* b = reinterpret_cast<const int*>(pb);

	return *b - *a;
}