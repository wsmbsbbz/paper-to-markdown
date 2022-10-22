import random
# NOTE: randint(a, b) return N such that a <= N <= b

class R:
    def randomList(self):
        min = eval(input("Min -> int:"))
        max = eval(input("Max -> int:"))
        amount = eval(input("Amount -> int:"))
        repetition = True if input("Repetition(y/n):") == 'y' else False
        sorting = True if input("Sorting(y/n):") == 'y' else False

        if amount > max - min + 1 and repetition is False:
            print("couldn't no-repetitions when amount > max - min + 1")
            exit()
        array = []
        opt = [i for i in range(min, max + 1)]

        if repetition is True:
            array = random.choices(population=opt, k=amount)
        else:
            for _ in range(amount):
                i = random.randint(0, len(opt) - 1)
                array.append(opt.pop(i))

        if sorting is True:
            array.sort()

        return array

    def randomInt(self):
        min = eval(input("Min(included) -> int:"))
        max = eval(input("Max(included) -> int:"))
        try:
            return random.randint(min, max)
        except ValueError:
            print('ValueError!', "'Max' should not be less than 'Min'")
            exit()

if __name__ == '__main__':
    r = R()
    # TODO: if randomList
    # print(r.randomList())
    # TODO: if randint
    print(r.randomInt())
