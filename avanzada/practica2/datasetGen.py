import random
import itertools

def generate_all_latin_squares(n):

    numbers = list(range(1, n + 1))
    permutations = itertools.permutations(numbers)

    latin_squares = []
    for perm in permutations:
        square = [[perm[(i + j) % n] for j in range(n)] for i in range(n)]
        latin_squares.append(square)

    return latin_squares


def print_latin_squares(squares):

    for idx, square in enumerate(squares, start=1):
        print(f"Latin Square {idx}:")
        for row in square:
            print(' '.join(map(str, row)))
        print()

def modify_density(square, density):

    n = len(square)
    modified_square = []

    for row in square:
        modified_row = []
        for val in row:
            if density == 1:
                modified_row.append(val)
            elif density == 0:
                modified_row.append('*')
            else:
                if random.random() <= density:
                    modified_row.append(val)
                else:
                    modified_row.append('*')
        modified_square.append(modified_row)

    return modified_square


def modify_density_in_squares(squares, density):

    modified_squares = []
    for square in squares:
        modified_square = modify_density(square, density)
        modified_squares.append(modified_square)

    return modified_squares


def write_squares_to_file(squares, file_name):

    with open(file_name, 'w') as file:
        for idx, square in enumerate(squares):
            for row in square:
                file.write(' '.join(map(str, row)) + '\n')
            if idx < len(squares) - 1:
                file.write("#\n")

def generate_dataset(size,density,nOfSquares):
    all_squares = generate_all_latin_squares(size)
    modified = modify_density_in_squares(all_squares, density)
    random.shuffle(modified)
    write_squares_to_file(modified[:nOfSquares],"latin_squares.txt")

if __name__ == "__main__":
    size = 3 
    all_squares = generate_all_latin_squares(size)
    modified = modify_density_in_squares(all_squares, 0.5)
    print_latin_squares(modified)
    write_squares_to_file(modified,"latin_squares.txt")
    