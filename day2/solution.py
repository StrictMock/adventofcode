
#!/usr/bin/env python

import re

def read_passwords():
    with open('input.txt', 'r') as fin:
        return fin.readlines()


def parse(entry):
    regex = "([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)"
    matched = re.search(regex, entry)
    first_num = int(matched.group(1))
    second_num = int(matched.group(2))
    letter = matched.group(3)
    password = matched.group(4)

    return first_num, second_num, letter, password


def part_one(passwords):
    counter = 0
    for p in passwords:
        min_count, max_count, letter, password = parse(p)
        
        if password.count(letter) in range(min_count, max_count+1):
            counter = counter + 1

    print("part one: {}".format(counter))


def part_two(passwords):
    counter = 0
    for p in passwords:
        first_pos, second_pos, letter, password = parse(p)
        
        if (letter in password[first_pos-1]) != (letter in password[second_pos-1]):
            counter = counter + 1

    print("part two: {}".format(counter))


def main():
    passwords = read_passwords()
    part_one(passwords)
    part_two(passwords)


if __name__ == "__main__":
    main()