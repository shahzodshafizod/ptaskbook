#include <iostream>
#include <fstream>
using namespace std;

int main()
{
    //Task("File78");
	char fileName[100], transFileName[100];
	cin.getline(fileName, 100);
	cin.getline(transFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream trans(transFileName, ios::binary);
	double element;
	file.read(reinterpret_cast<char*>(&element), sizeof(element));
	int cols = (int)element;
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	elements--;
	int rows = elements / cols;
	
	double** matr = new double* [rows];
	file.seekg(sizeof(double));
	for (int i = 0; i < rows; i++)
	{
		matr[i] = new double [cols];
		for (int j = 0; j < cols; j++)
			file.read(reinterpret_cast<char*>(&matr[i][j]), sizeof(matr[i][j]));
	}
	file.close();

	double** transponed = new double* [cols];
	for (int i = 0; i < cols; i++)
	{
		transponed[i] = new double [rows];
		for (int j = 0; j < rows; j++)
			transponed[i][j] = matr[j][i];
	}

	for (int i = 0; i < rows; i++)
		delete [] matr[i];
	delete [] matr;
	matr = transponed;
	transponed = NULL;

	rows += cols;
	cols = rows - cols;
	rows -= cols;

	element = cols;
	trans.write(reinterpret_cast<const char*>(&element), sizeof(element));
	for (int i = 0; i < rows; i++)
		for (int j = 0; j < cols; j++)
			trans.write(reinterpret_cast<const char*>(&matr[i][j]), sizeof(matr[i][j]));
	
	for (int i = 0; i < rows; i++)
		delete [] matr[i];
	delete [] matr;
	matr = NULL;

	trans.close();
	
	return 0;
}
