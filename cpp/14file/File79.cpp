#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File79");
	char SA[100], SB[100], SC[100];
	cin.getline(SA, 100);
	cin.getline(SB, 100);
	cin.getline(SC, 100);
	ifstream fileA(SA, ios::binary);
	ifstream fileB(SB, ios::binary);
	ofstream fileC(SC, ios::binary);

	double element;
	int elements;

	fileA.read(reinterpret_cast<char*>(&element), sizeof(element));
	int Acols = (int)element;
	fileA.seekg(0, ios::end);
	elements = fileA.tellg() / sizeof(double);
	elements--; // because a count of cols is also keeped in the file and it's not an element of matrix
	int Arows = elements / Acols;
	
	fileB.read(reinterpret_cast<char*>(&element), sizeof(element));
	int Bcols = (int)element;
	fileB.seekg(0, ios::end);
	elements = fileB.tellg() / sizeof(double);
	elements--;
	int Brows = elements / Bcols;

	if (Acols == Brows)
	{
		fileA.seekg(sizeof(double));
		double** A = new double* [Arows];
		for (int i = 0; i < Arows; i++)
		{
			A[i] = new double [Acols];
			for (int j = 0; j < Acols; j++)
				fileA.read(reinterpret_cast<char*>(&A[i][j]), sizeof(A[i][j]));
		}

		fileB.seekg(sizeof(double));
		double** B = new double* [Brows];
		for (int i = 0; i < Brows; i++)
		{
			B[i] = new double [Bcols];
			for (int j = 0; j < Bcols; j++)
				fileB.read(reinterpret_cast<char*>(&B[i][j]), sizeof(B[i][j]));
		}

		int Crows = Arows;
		int Ccols = Bcols;
		double** C = new double* [Crows];
		for (int i = 0; i < Crows; i++)
		{
			C[i] = new double [Ccols];
			for (int j = 0; j < Ccols; j++)
			{
				C[i][j] = 0;
				for (int k = 0; k < Acols; k++)
					C[i][j] += A[i][k]*B[k][j];
			}
		}
		
		element = Ccols;
		fileC.write(reinterpret_cast<const char*>(&element), sizeof(element));
		for (int i = 0; i < Crows; i++)
			for (int j = 0; j < Ccols; j++)
				fileC.write(reinterpret_cast<const char*>(&C[i][j]), sizeof(C[i][j]));

		for (int i = 0; i < Arows; i++)
			delete [] A[i];
		delete [] A;
		A = NULL;

		for (int i = 0; i < Brows; i++)
			delete [] B[i];
		delete [] B;
		B = NULL;

		for (int i = 0; i < Crows; i++)
			delete [] C[i];
		delete [] C;
		C = NULL;
	}
	fileA.close();
	fileB.close();
	fileC.close();
	
	return 0;
}
