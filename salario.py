from itertools import groupby
from getpass import getpass
from os import system
import os

def all_equal(iterable):
    g = groupby(iterable)
    return next(g, True) and not next(g, False)

def clearConsole():
    command = 'clear'
    if os.name in ('nt', 'dos'):
        command = 'cls'
    os.system(command)

clearConsole()
print("\n-------- Salario Matcher --------")
salario = []
for i in range(3):
    salario.append(getpass("Salario {}: ".format(i+1)))

print(("\n----------------\n-----Match!-----\n----------------\n"
       if all_equal(salario) else
       "\n----------------\n----No Match----\n----------------\n"))
print("---------------------------------\n")
