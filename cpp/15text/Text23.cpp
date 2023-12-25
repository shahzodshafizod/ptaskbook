#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text23");
	int K;
	cin >> K;
	cin.ignore();
	const int stringSize = 80;
	char fileName1[100], fileName2[100];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	ifstream file1(fileName1);
	ofstream file2(fileName2);
	char** str = new char* [K];
	for (int i = 0; i < K; i++)
	{
		str[i] = new char[stringSize];
		file1.getline(str[i], stringSize);
	}

	int index = -1;
	while (file1.peek() != -1)
	{
		index = (index+1) % K;
		file1.getline(str[index], stringSize);
	}
	file1.close();
	
	for (int i = index+1; i < K; i++)
		file2 << str[i] << endl;
	
	for (int i = 0; i <= index; i++)
		file2 << str[i] << endl;

	file2.close();
	
	for (int i = 0; i < K; i++)
		delete [] str[i];
	delete [] str;
	str = NULL;
	
	return 0;
}
