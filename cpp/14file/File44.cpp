#include <iostream>
#include <fstream>
using namespace std;

void ReplaceDatas(const char* from, const char* to);

int main()
{
	//Task("File44");
	
	char fileName1[100], fileName2[100], fileName3[100];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	cin.getline(fileName3, 100);
	
	ifstream file1(fileName1, ios::binary);
	ifstream file2(fileName2, ios::binary);
	ifstream file3(fileName3, ios::binary);
	
	file1.seekg(0, ios::end);
	file2.seekg(0, ios::end);
	file3.seekg(0, ios::end);

	int file1Size = file1.tellg();
	int file2Size = file2.tellg();
	int file3Size = file3.tellg();

	int xurd, kalon;
	if ( (file1Size < file2Size) && (file1Size < file3Size) )
		xurd = 1;
	else if (file2Size < file3Size)
		xurd = 2;
	else
		xurd = 3;

	if ( (file1Size > file2Size) && (file1Size > file3Size) )
		kalon = 1;
	else if (file2Size > file3Size)
		kalon = 2;
	else
		kalon = 3;
	
	ReplaceDatas( xurd == 1 ? fileName1 : xurd == 2 ? fileName2 : fileName3,
				kalon == 1 ? fileName1 : kalon == 2 ? fileName2 : fileName3 );
	
	return 0;
}

void ReplaceDatas(const char* from, const char* to)
{
	ifstream ifs(from, ios::binary);
	ofstream ofs(to, ios::binary);
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