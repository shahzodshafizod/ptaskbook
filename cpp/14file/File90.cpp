#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File90");
	char SA[100], SB[100], SC[100];
	cin.getline(SA, 100);
	cin.getline(SB, 100);
	cin.getline(SC, 100);
	ifstream fileA(SA, ios::binary);
	ifstream fileB(SB, ios::binary);
	ofstream fileC(SC, ios::binary);

	int elements;

	fileA.seekg(0, ios::end);
	elements = fileA.tellg() / sizeof(double);
	int ArowCols = 0;
	for (int i = 1; elements > 0; i++)
	{
		ArowCols++;
		elements -= i;
	}
	fileA.seekg(0);

	fileB.seekg(0, ios::end);
	elements = fileB.tellg() / sizeof(double);
	int BrowCols = 0;
	for (int i = 1; elements > 0; i++)
	{
		BrowCols++;
		elements -= i;
	}
	fileB.seekg(0);

	if (ArowCols == BrowCols)
	{
		double** A = new double* [ArowCols];
		for (int i = 0; i < ArowCols; i++)
		{
			A[i] = new double [ArowCols];
			for (int j = 0; j < ArowCols; j++)
				A[i][j] = 0;
		}
		for (int i = 0; i < ArowCols; i++)
			for (int j = 0; j <= i; j++)
				fileA.read(reinterpret_cast<char*>(&A[i][j]), sizeof(A[i][j]));
		
		double** B = new double* [BrowCols];
		for (int i = 0; i < BrowCols; i++)
		{
			B[i] = new double [BrowCols];
			for (int j = 0; j < BrowCols; j++)
				B[i][j] = 0;
		}
		for (int i = 0; i < BrowCols; i++)
			for (int j = 0; j <= i; j++)
				fileB.read(reinterpret_cast<char*>(&B[i][j]), sizeof(B[i][j]));

		int CrowCols = ArowCols;
		double** C = new double* [CrowCols];
		for (int i = 0; i < CrowCols; i++)
		{
			C[i] = new double [CrowCols];
			for (int j = 0; j < CrowCols; j++)
			{
				C[i][j] = 0;
				for (int k = 0; k < CrowCols; k++)
					C[i][j] += A[i][k] * B[k][j];
			}
		}
		for (int i = 0; i < CrowCols; i++)
			for (int j = 0; j <= i; j++)
				fileC.write(reinterpret_cast<const char*>(&C[i][j]), sizeof(C[i][j]));

		for (int i = 0; i < CrowCols; i++)
		{
			delete [] A[i];
			delete [] B[i];
			delete [] C[i];
		}
		delete [] A;
		delete [] B;
		delete [] C;
		A = B = C = NULL;
	}

	fileA.close();
	fileB.close();
	fileC.close();
	
	return 0;
}
