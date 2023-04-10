import random
import string
import sys
num_words = int(sys.argv[1])
basic_words = set(''.join(random.choice(string.ascii_lowercase) for _ in range(random.randint(30,60))) for _ in range(num_words*925//1000))

compound_words = set()
while len(compound_words) < num_words*75//1000:
    bw1 = random.choice(list(basic_words))
    bw2 = random.choice(list(basic_words))
    compound_words.add(bw1 + bw2)

all_words = basic_words.union(compound_words)

first_line = ''.join(random.sample(list(all_words), random.randint(num_words//100 + 1, num_words//10)))

with open('./tmp/f' + sys.argv[1] + '.txt', 'w') as f:
    f.write(first_line + '\n')
    for word in all_words:
        f.write(word + '\n')

def modify_sentence(sentence):
    LF = len(sentence)
    probability = 1 / (LF * 10)
    new_sentence = ""
    for index, letter in enumerate(sentence):
        if random.random() <= probability:
            new_sentence += random.choice(string.ascii_lowercase)
        else:
            new_sentence += letter
    return new_sentence

with open('./tmp/fMod' + sys.argv[1] + '.txt', 'w') as f:
    f.write(modify_sentence(first_line) + '\n')
    for word in all_words:
        f.write(word + '\n')