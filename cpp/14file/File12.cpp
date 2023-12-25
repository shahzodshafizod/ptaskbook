#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File12");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ifs.seekg(0, ios::end);
	int elements = ifs.tellg() / sizeof(int);
	int* mass = new int [elements];
	ifs.seekg(0);
	for (int i = 0; i < elements; i++)
		ifs.read(reinterpret_cast<char*>(&mass[i]), sizeof(int));
	ifs.close();
	char fileName1[100], fileName2[100];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	ofstream file1(fileName1, ios::binary);
	ofstream file2(fileName2, ios::binary);
	for (int i = 0; i < elements; i++)
		if (mass[i] % 2 == 0)
			file1.write(reinterpret_cast<const char*>(&mass[i]), sizeof(mass[i]));
		else
			file2.write(reinterpret_cast<const char*>(&mass[i]), sizeof(mass[i]));
	file1.close();
	file2.close();
	delete [] mass;
	mass = NULL;
	
	return 0;
}
