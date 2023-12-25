#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File61");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	ifs.seekg(0, ios::end);
	int pos = ifs.tellg();
	char C;
	for (pos -= sizeof(C); true; pos -= sizeof(C))
	{
		ifs.seekg(pos);
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (ifs.eof() || (C == ' '))
		{
			ifs.clear();
			break;
		}
	}
	ifs.seekg(pos + sizeof(char));
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		ofs.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
