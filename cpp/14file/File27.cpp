#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File27");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	char tempFileName[] = "tempFileNameForFile27.txt";
	ofstream ofs(tempFileName, ios::binary);
	ifs.seekg(0, ios::end);
	int intSize = sizeof(int);
	int elements = ifs.tellg() / intSize;
	int number;
	int i = 0;
	for (; i < elements/2; i++)
	{
		ifs.seekg(i*intSize);
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));

		ifs.seekg((elements-1-i)*intSize);
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	if (elements % 2 != 0)
	{
		ifs.seekg(i*intSize);
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
