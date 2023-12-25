#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File22");
	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream ifs(fileInName, ios::binary);
	ofstream ofs(fileOutName, ios::binary);
	double prev, curr, next;
	ifs.seekg(0, ios::end);
	int elements = ifs.tellg() / sizeof(double);
	
	int element = elements;
	ofs.write(reinterpret_cast<const char*>(&element), sizeof(element));

	ifs.seekg((element-1)*sizeof(double));
	ifs.read(reinterpret_cast<char*>(&next), sizeof(next));
	Show(next);

	element--;
	ifs.seekg((element-1)*sizeof(double));
	ifs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	Show(curr);

	while (element > 0)
	{
		element--;
		ifs.seekg((element-1)*sizeof(double));
		ifs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
		Show(prev);
		
		if ( (curr > prev) && (curr > next) || (curr < prev) && (curr < next) )
		{
			element++;
			ofs.write(reinterpret_cast<const char*>(&element), sizeof(element));
			element--;
		}
		next = curr;
		curr = prev;
	}
	element++;
	ofs.write(reinterpret_cast<const char*>(&element), sizeof(element));

	ifs.close();
	ofs.close();
	
	return 0;
}
