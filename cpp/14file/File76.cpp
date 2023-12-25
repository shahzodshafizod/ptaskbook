#include <iostream>
#include <fstream>
#include <cmath>
using namespace std;

int main()
{
	//Task("File76");
	char SA[100], SB[100], SC[100];
	cin.getline(SA, 100);
	cin.getline(SB, 100);
	cin.getline(SC, 100);
	
	ifstream fileA(SA, ios::binary);
	ifstream fileB(SB, ios::binary);
	ofstream fileC(SC, ios::binary);
	
	int ARowCols, BRowCols, elements;

	fileA.seekg(0, ios::end);
	elements = fileA.tellg() / sizeof(double);
	ARowCols = (int)sqrt((double)elements);

	fileB.seekg(0, ios::end);
	elements = fileB.tellg() / sizeof(double);
	BRowCols = (int)sqrt((double)elements);

	if (ARowCols == BRowCols)
	{
		fileA.seekg(0);
		fileB.seekg(0);
		double** massA = new double* [ARowCols];
		double** massB = new double* [ARowCols];
		for (int i = 0; i < ARowCols; i++)
		{
			massA[i] = new double [ARowCols];
			massB[i] = new double [ARowCols];
			for (int j = 0; j < ARowCols; j++)
			{
				fileA.read(reinterpret_cast<char*>(&massA[i][j]), sizeof(massA[i][j]));
				fileB.read(reinterpret_cast<char*>(&massB[i][j]), sizeof(massB[i][j]));
			}
		}
		double** massC = new double* [ARowCols];
		for (int i = 0; i < ARowCols; i++)
		{
			massC[i] = new double [ARowCols];
			for (int j = 0; j < ARowCols; j++)
			{
				massC[i][j] = 0;
				for (int k = 0; k < ARowCols; k++)
					massC[i][j] += massA[i][k] * massB[k][j];
				fileC.write(reinterpret_cast<const char*>(&massC[i][j]), sizeof(massC[i][j]));
			}
		}
		for (int i = 0; i < ARowCols; i++)
		{
			delete [] massA[i];
			delete [] massB[i];
			delete [] massC[i];
		}
		delete [] massA;
		delete [] massB;
		delete [] massC;
		massA = massB = massC = NULL;
	}
	fileA.close();
	fileB.close();
	fileC.close();
	
	return 0;
}
