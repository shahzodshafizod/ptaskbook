#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File2");
	char fileName[100];
	cin.getline(fileName, 100);
	int N;
	cin >> N;
	ofstream file(fileName, ios::binary);
	for (int i = 2, j = 1; j <= N; i += 2, j++)
		file.write( reinterpret_cast<const char*>(&i), sizeof(i) );
	file.close();
	
	return 0;
}
