#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File43");
	char fileName[100], copyFileName[100];
	cin.getline(fileName, 100);
	cin.getline(copyFileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(copyFileName, ios::binary);
	char C;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		ofs.write(reinterpret_cast<const char*>(&C),sizeof(C));
	}
	ifs.close();
	ofs.close();
	
	return 0;
}
