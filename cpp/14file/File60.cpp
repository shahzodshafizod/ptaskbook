#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File60");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	char C;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if ( ifs.eof() || (C == ' ') )
		{
			ifs.clear();
			break;
		}
	}
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
