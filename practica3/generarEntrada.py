import random
import string
import sys
num_words = int(sys.argv[1])
words = set(''.join(random.choice(string.ascii_lowercase) for _ in range(random.randint(2,5))) for _ in range(num_words))
first_line = ''.join(random.sample(list(words), random.randint(num_words/10, num_words)))
with open('./tmp/f' + sys.argv[1] + '.txt', 'w') as f:
    f.write(first_line + '\n')
    for word in words:
        f.write(word + '\n')