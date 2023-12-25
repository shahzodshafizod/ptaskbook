#include <iostream>
#include <fstream>
using namespace std;

void InvIntFile(const char* S);

int main()
{
	//Task("Param50");
	char fileName[100];
	for (int i = 1; i < 4; i++)
	{
		cin.getline(fileName, 100);
		InvIntFile(fileName);
	}
	
	return 0;
}

void InvIntFile(const char* S)
{
	fstream file(S, ios::binary | ios::in | ios::out);
	if (file.is_open())
	{
		file.seekg(0, ios::end);
		int number1, number2, numbers = file.tellg() / sizeof(int);
		for (int i = 0, j = numbers-1; i < j; i++, j--)
		{
			file.seekg(i*sizeof(int));
			file.read(reinterpret_cast<char*>(&number1), sizeof(number1));

			file.seekg(j*sizeof(int));
			file.read(reinterpret_cast<char*>(&number2), sizeof(number2));

			file.seekg(i*sizeof(int));
			file.write(reinterpret_cast<const char*>(&number2), sizeof(number2));

			file.seekg(j*sizeof(int));
			file.write(reinterpret_cast<const char*>(&number1), sizeof(number1));
		}
		file.close();
	}
}