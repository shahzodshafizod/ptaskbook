#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File84");
	int I, J;
	cin >> I >> J;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	int RowCols = 0;
	for (int i = 1; elements > 0; i++)
	{
		elements -= i;
		RowCols++;
	}
	file.seekg(0);

	double element = -1;
	if ( (I <= RowCols) && (J <= RowCols) )
	{
		double** matr = new double* [RowCols];
		for (int i = 0; i < RowCols; i++)
		{
			matr[i] = new double [RowCols];
			for (int j = 0; j < RowCols; j++)
				matr[i][j] = 0;
		}
		
		for (int i = 0; i < RowCols; i++)
			for (int j = 0; j <= i; j++)
				file.read(reinterpret_cast<char*>(&matr[i][j]), sizeof(matr[i][j]));

		element = matr[I-1][J-1];
		for (int i = 0; i < RowCols; i++)
			delete [] matr[i];
		delete [] matr;
		matr = NULL;
	}
	file.close();

	cout << RowCols << element;
	
	return 0;
}
