#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File87");
	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream fileIn(fileInName, ios::binary);
	ofstream fileOut(fileOutName, ios::binary);
	fileIn.seekg(0, ios::end);
	int elements = fileIn.tellg() / sizeof(double);
	int RowCols = 0;
	for (int i = 1; elements > 0; i++)
	{
		RowCols++;
		elements -= i;
	}

	double** matr = new double* [RowCols];
	for (int i = 0; i < RowCols; i++)
	{
		matr[i] = new double [RowCols];
		for (int j = 0; j < RowCols; j++)
			matr[i][j] = 0;
	}

	fileIn.seekg(0);
	for (int i = 0; i < RowCols; i++)
		for (int j = 0; j <= i; j++)
			fileIn.read(reinterpret_cast<char*>(&matr[i][j]), sizeof(matr[i][j]));

	for (int i = 0; i < RowCols; i++)
		for (int j = 0; j < RowCols; j++)
			fileOut.write(reinterpret_cast<const char*>(&matr[i][j]), sizeof(matr[i][j]));

	for (int i = 0; i < RowCols; i++)
		delete [] matr[i];
	delete [] matr;
	matr = NULL;

	fileIn.close();
	fileOut.close();
	
	return 0;
}
