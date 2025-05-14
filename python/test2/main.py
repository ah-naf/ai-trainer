import re
from collections import Counter

def count_word_frequencies(text):
    text = text.lower()
    text = re.sub(r'[^\w\s]', ' ', text)
    
    words = text.split()
    
    word_counts = Counter(words)
    
    sorted_word_counts = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))
    
    print("Word Frequencies:")
    for word, count in sorted_word_counts:
        print(f"{word}: {count}")
    
    return word_counts

if __name__ == "__main__":
    text = "Hello world! This is a test. Hello, this test is only a test."
    count_word_frequencies(text)