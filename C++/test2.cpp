#include <iostream>
#include <vector>
using namespace std;

int main()
{
    int rows, cols;
    cout << "Rows: ";
    cin >> rows;
    cout << "Columns: ";
    cin >> cols;

    vector<vector<int>> matrix(rows, vector<int>(cols));
    cout << "Matrix:\n";
    for (int i = 0; i < rows; ++i)
        for (int j = 0; j < cols; ++j)
            cin >> matrix[i][j];

    cout << "\nOriginal matrix:\n";
    for (int i = 0; i < rows; ++i)
    {
        for (int j = 0; j < cols; ++j)
            cout << matrix[i][j] << " ";
        cout << "\n";
    }

    vector<vector<int>> transposed(cols, vector<int>(rows));
    for (int i = 0; i < rows; ++i)
        for (int j = 0; j < cols; ++j)
            transposed[j][i] = matrix[i][j];

    cout << "\nTransposed matrix:\n";
    for (int i = 0; i < cols; ++i)
    {
        for (int j = 0; j < rows; ++j)
            cout << transposed[i][j] << " ";
        cout << "\n";
    }

    return 0;
}
