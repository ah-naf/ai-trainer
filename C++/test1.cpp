#include <iostream>
#include <vector>
#include <string>

std::vector<std::string> formatText(const std::vector<std::string> &words, int maxWidth)
{
    std::vector<std::string> lines;
    std::string current;
    for (const auto &w : words)
    {
        if (current.empty())
        {
            current = w;
        }
        else if (current.size() + 1 + w.size() <= maxWidth)
        {
            current += ' ' + w;
        }
        else
        {
            lines.push_back(current);
            current = w;
        }
        if (w.size() > maxWidth && current == w)
        {
            lines.push_back(current);
            current.clear();
        }
    }
    if (!current.empty())
        lines.push_back(current);
    return lines;
}

int main()
{
    std::vector<std::string> words = {"This", "is", "an", "example", "of", "text", "justification."};
    int maxWidth = 10;
    auto res = formatText(words, maxWidth);
    for (const auto &line : res)
        std::cout << line << '\n';
    return 0;
}
