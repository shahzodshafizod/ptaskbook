#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File85");
	int I, J;
	cin >> I >> J;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	int RowCols = (elements-4) / 3 + 2;
	
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

		file.seekg(0);
		for (int i = 0; i < RowCols; i++)
			for (int j = i-1; j < i+2; j++)
			{
				if ( (j < 0) || (j >= RowCols) )
					continue;
				file.read(reinterpret_cast<char*>(&matr[i][j]), sizeof(matr[i][j]));
			}
		
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
