array = [
	[1,3,4,9,12],
	[2,4,9,12,14],
	[4,7,11,13,15],
	[6,9,13,19,22],
	[8,14,19,21,24]
]

def search(key, m=4, n=0):
	
	while array[n][m] != key and ( m !=0 and n != 4):
		if array[n][m] > key:
			m -= 1
		if array[n][m] < key:
			n += 1
	if array[n][m] == key:
		print(m, n)
		return True
	else:
		print(m, n)
		return False
	
	
print(search(24))