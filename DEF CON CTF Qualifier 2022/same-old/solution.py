from binascii import crc32
from itertools import product
from string import ascii_letters, digits
from sys import exit

team_name = 'HiCr33k'
allowed_chars = list(ascii_letters + digits)
the_crc = crc32('the'.encode())

i = 0
while True:
    print(f'trying length: {i}')
    for prod in product(allowed_chars, repeat=i):
        if crc32((team_name + ''.join(prod)).encode()) == the_crc:
            print(prod, ''.join(prod))
            exit()
    i += 1
