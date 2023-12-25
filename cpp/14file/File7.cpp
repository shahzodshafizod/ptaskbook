#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File7");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ifs.seekg(0, ios::end);
	int number, elements = ifs.tellg() / sizeof(int);
	
	ifs.seekg(0);
	for (int i = 0; i < 2; i++)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		cout << number;
	}
	
	ifs.seekg((elements-2)*sizeof(int));
	for (int i = 0; i < 2; i++)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		cout << number;
	}
	ifs.close();
	
	return 0;
}
