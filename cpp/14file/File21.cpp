#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File21");
	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream ifs(fileInName, ios::binary);
	ofstream ofs(fileOutName, ios::binary);
	double prev, curr, next;
	int index = 1;
	ifs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
	ifs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	if (prev > curr)
		ofs.write(reinterpret_cast<const char*>(&index), sizeof(index));
	index++;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&next), sizeof(next));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if ( (curr > prev) && (curr > next) )
			ofs.write(reinterpret_cast<const char*>(&index), sizeof(index));
		index++;
		prev = curr;
		curr = next;
	}
	if (curr > prev)
		ofs.write(reinterpret_cast<const char*>(&index), sizeof(index));
	ifs.close();
	ofs.close();
	
	return 0;
}
