#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File20");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	double prev, curr, next;
	int counts = 0;
	ifs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
	ifs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	if (prev != curr)
		counts++;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&next), sizeof(next));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if ( (curr < prev) && (curr < next) || (curr > prev) && (curr > next) )
			counts++;
		prev = curr;
		curr = next;
	}
	if (curr != prev)
		counts++;
	
	cout << counts;
	
	return 0;
}
