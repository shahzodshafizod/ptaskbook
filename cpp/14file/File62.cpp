#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File62");
	char *mass, C, fileName[100];
	cin.getline(fileName, 100);
	fstream file(fileName, ios::binary | ios::in | ios::out);
	file.seekg(0, ios::end);
	int elements = file.tellg() / sizeof(char);
	mass = new char [elements];
	file.seekg(0);
	for (int i = 0; i < elements; i++)
		file.read(reinterpret_cast<char*>(&mass[i]), sizeof(mass[i]));
	for (int i = 0; i < elements-1; i++)
		for (int j = 1; j < elements-i; j++)
			if (int(mass[j]) < int(mass[j-1]))
			{
				C = mass[j];
				mass[j] = mass[j-1];
				mass[j-1] = C;
			}
			
	file.seekg(0);
	for (int i = 0; i < elements; i++)
		file.write(reinterpret_cast<const char*>(&mass[i]), sizeof(mass[i]));
	file.close();
	
	return 0;
}
