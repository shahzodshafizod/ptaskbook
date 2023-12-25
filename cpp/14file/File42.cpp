#include <iostream>
#include <fstream>
using namespace std;

void ReplaceDatas(const char* FNfrom, const char* FNto);

int main()
{
	//Task("File42");
	char fileName1[100], fileName2[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	
	ReplaceDatas(fileName1, tempFileName);
	ReplaceDatas(fileName2, fileName1);
	ReplaceDatas(tempFileName, fileName2);
	remove(tempFileName);
	
	return 0;
}

void ReplaceDatas(const char* FNfrom, const char* FNto)
{
	ifstream ifs(FNfrom, ios::binary);
	ofstream ofs(FNto, ios::binary);
	char C;
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
}