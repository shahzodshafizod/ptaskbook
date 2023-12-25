#include <iostream>
#include <fstream>
#include <cmath>
using namespace std;

int main()
{
	//Task("File74");
	int I, J;
	cin >> I >> J;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	int rows, cols;
	rows = cols = (int)sqrt((double)elements);
	double element = 0;
	if ( (I <= rows) && (J <= cols) )
	{
		bool exit = false;
		double** mass = new double*[rows];
		file.seekg(0);
		for (int i = 0; i < rows; i++)
		{
			mass[i] = new double [cols];
			for (int j = 0; j < cols; j++)
			{
				file.read(reinterpret_cast<char*>(&mass[i][j]), sizeof(mass[i][j]));
				if ( (i == I-1) && (j == J-1) )
				{
					element = mass[i][j];
					exit = true;
					break;
				}
			}
			if (exit)
				break;
		}
	}
	cout << element;
	
	return 0;
}
