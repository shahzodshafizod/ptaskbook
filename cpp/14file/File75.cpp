#include <iostream>
#include <fstream>
#include <cmath>
using namespace std;

int main()
{
	//Task("File75");
	char fileName[100], transponeFileName[100];
	cin.getline(fileName, 100);
	cin.getline(transponeFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream transponeFile(transponeFileName, ios::binary);
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	int rows, cols;
	rows = cols = (int)sqrt((double)elements);
	double** mass = new double* [rows];
	file.seekg(0);
	for (int i = 0; i < rows; i++)
	{
		mass[i] = new double [cols];
		for (int j = 0; j < cols; j++)
			file.read(reinterpret_cast<char*>(&mass[i][j]), sizeof(mass[i][j]));
	}
	double temp;
	for (int i = 0; i < rows; i++)
	{
		for (int j = i+1; j < cols; j++)
		{
			temp = mass[i][j];
			mass[i][j] = mass[j][i];
			mass[j][i] = temp;
		}
	}
	for (int i = 0; i < rows; i++)
	{
		for (int j = 0; j < cols; j++)
			transponeFile.write(reinterpret_cast<const char*>(&mass[i][j]), sizeof(mass[i][j]));
		delete [] mass[i];
	}
	delete [] mass;
	mass = NULL;
	file.close();
	transponeFile.close();
	
	return 0;
}
