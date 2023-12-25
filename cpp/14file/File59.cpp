#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File59");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	ifs.seekg(0, ios::end);
	int pos = ifs.tellg();
	int elements = pos / sizeof(char);
	char C;
	for (pos -= sizeof(char); pos >= 0; pos -= sizeof(char))
	{
		ifs.seekg(pos);
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (C == ' ')
			break;
	}
	ifs.seekg(0);
	for (int i = 0; i < pos; i += sizeof(char))
	{
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		ofs.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
