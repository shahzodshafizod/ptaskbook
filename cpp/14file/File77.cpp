#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File77");
	int I, J;
	cin >> I >> J;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	double element;
	file.read(reinterpret_cast<char*>(&element), sizeof(element));
	int cols = (int)element;
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(double);
	elements--; // because a count of cols is also keeped in file and it's not an element of matrix
	int rows = elements / cols;
	element = 0;
	if ( (I <= rows) && (J <= cols) )
	{
		file.seekg( sizeof(double) );
		file.seekg( (cols*(I-1))*sizeof(double), ios::cur );
		file.seekg( (J-1)*sizeof(double), ios::cur );
		file.read(reinterpret_cast<char*>(&element), sizeof(element));
	}
	file.close();
	cout << element;
	
	return 0;
}
