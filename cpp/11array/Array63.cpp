#include <iostream>
using namespace std;

int main()
{
	//Task("Array63");
	
	double A[5], B[5];

	for (int i = 0; i < 5; i++)
		cin >> A[i];

	for (int i = 0; i < 5; i++)
		cin >> B[i];

	double C[10];
	for (int indexC = 0, indexA = 0, indexB = 0; indexC < 10; indexC++)
	{
		if (indexA >= 5)
			C[indexC] = B[indexB++];
		
		else if (indexB >= 5)
			C[indexC] = A[indexA++];

		else if (A[indexA] < B[indexB])
			C[indexC] = A[indexA++];
		
		else
			C[indexC] = B[indexB++];
	}

	for (int i = 0; i < 10; i++)
		cout << C[i];
	
	return 0;
}
