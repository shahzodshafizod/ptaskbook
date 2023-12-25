#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File8");
	char fileName1[100], fileName2[100];
	cin.getline(fileName1, 100)
	cin.getline(fileName2, 100);
	ifstream ifs(fileName1, ios::binary);
	ofstream ofs(fileName2, ios::binary);
	ifs.seekg(0, ios::end);
	int elements = ifs.tellg() / sizeof(double);
	double number;
	ifs.seekg(0);
	ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
	ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));

	ifs.seekg((elements-1)*sizeof(double));
	ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
	ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	ifs.close();
	ofs.close();
	
	return 0;
}
