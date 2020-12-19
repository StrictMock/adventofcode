
#!/usr/bin/env python

def read_input():
    with open('input.txt', 'r') as fin:
        return [x.rstrip() for x in fin.readlines()]

def part_one(data):
    unique_questions = set()
    sum = 0
    for answer in data:
        if not answer:
            sum = sum + len(unique_questions)
            unique_questions = set()
            continue
        for question in answer:
            unique_questions.add(question)

    return sum


def part_two(data):
    unique_questions = {}
    sum = 0
    group_count = 0
    for answer in data:
        if not answer:
            sum = sum + len([x for x in unique_questions.values() if x == group_count])
            unique_questions = {}
            group_count = 0
            continue
        group_count = group_count + 1
        for question in answer:
            unique_questions[question] = unique_questions.get(question, 0) + 1

    return sum


def main():
    data = read_input()
    print("part one: {}".format(part_one(data)))
    print("part two: {}".format(part_two(data)))


if __name__ == "__main__":
    main()
