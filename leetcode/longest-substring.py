def lengthOfLongestSubstring(s):
    str_length = len(s)
        
    if str_length == 0:
        return 0
    
    max_length = 0
    start = 0
    end = 0
    current_item = {}
    
    while start <= str_length:
        
        for i in range(end, str_length):
            if current_item.get(s[i]):
                if start < str_length:
                    current_item.pop(s[start])
                end = i
                length = end - start
                if max_length < length:
                    max_length = length
                break
            else:
                current_item[s[i]] = True
        else:
            length = str_length - start
            if max_length < length:
                max_length = length
        print(start, end, current_item, max_length) 
        start += 1
    
    return max_length

a = "bbbbbb"
a1 = "pwwkew"
a2 = "qrsvbspk"
a3 = "au"
a4 = ""
print(lengthOfLongestSubstring(a))
print(lengthOfLongestSubstring(a1))
print(a2); print(lengthOfLongestSubstring(a2))
print(lengthOfLongestSubstring(a3))
print(lengthOfLongestSubstring(a4))







