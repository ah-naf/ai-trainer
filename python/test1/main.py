def is_valid_string(s):
    if len(s) < 6:
        return False
    
    digit_count = sum(c.isdigit() for c in s)
    if digit_count < 2 or digit_count > 3:
        return False
    
    prev_digit_index = -2
    for i, char in enumerate(s):
        if char.isdigit():
            if i - prev_digit_index <= 1:
                return False
            prev_digit_index = i
            
    return True