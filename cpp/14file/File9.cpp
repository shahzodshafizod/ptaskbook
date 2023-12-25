#include <iostream>
#include <fstream>
using namespace std;

void Do(fstream& fs, char* fileName);

int main()
{
	//Task("File9");
	char fileName1[100], fileName2[100];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	fstream one(fileName1, ios::binary | ios::in | ios::out);
	fstream two(fileName2, ios::binary | ios::in | ios::out);
	if (one.is_open())
	{
		Do(one, fileName2);
		one.close();
	}
	else
	{
		Do(two, fileName1);
		two.close();
	}
	
	return 0;
}

void Do(fstream& fs, char* fileName)
{
	fs.seekg(0, ios::end);
	int elements = fs.tellg() / sizeof(double);
	ofstream ofs(fileName, ios::binary);
	double number;
	fs.seekg((elements-1)*sizeof(double));
	fs.read(reinterpret_cast<char*>(&number), sizeof(number));
	ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	
	fs.seekg(0);
	fs.read(reinterpret_cast<char*>(&number), sizeof(number));
	ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));

	ofs.close();
}