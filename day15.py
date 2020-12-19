from collections import defaultdict

elveswords = [0, 3, 6]

"""
    if i+1 exists, do nothing
    else
        has i been seen before?
        no -> 0
        yes? -> walk back indices, last - previous

"""

n = 2020
i = 0
while i < n:
    # print(input[i], input)
    if i < len(elveswords) - 1:
        i += 1
        continue
    s = elveswords[i]
    indices = [i for i, e in enumerate(elveswords) if e == s]  # has it been seen before?
    if len(indices) == 0:
        # print("never seen", s, "before")
        elveswords.append(0)  # never seen before
    else:
        # we have seen it
        lim = indices[len(indices) - 2:]
        diff = lim[1] - lim[0] if len(lim) > 1 else i - lim[0]
        elveswords.append(diff)
        # print(s, "has seen it before in pos", indices[len(indices) - 2:], "diff:", diff)
    if i == n - 1:
        print(s)
    i += 1

# For example, suppose the starting numbers are 0,3,6:
#
# Turn 1: The 1st number spoken is a starting number, 0.
# Turn 2: The 2nd number spoken is a starting number, 3.
# Turn 3: The 3rd number spoken is a starting number, 6.
# Turn 4: Now, consider the last number spoken, 6. Since that was the first time the number had been spoken, the 4th number spoken is 0.
# Turn 5: Next, again consider the last number spoken, 0. Since it had been spoken before, the next number to speak is the difference between the turn number when it was last spoken (the previous turn, 4) and the turn number of the time it was most recently spoken before then (turn 1). Thus, the 5th number spoken is 4 - 1, 3.
# Turn 6: The last number spoken, 3 had also been spoken before, most recently on turns 5 and 2. So, the 6th number spoken is 5 - 2, 3.
# Turn 7: Since 3 was just spoken twice in a row, and the last two turns are 1 turn apart, the 7th number spoken is 1.
# Turn 8: Since 1 is new, the 8th number spoken is 0.
# Turn 9: 0 was last spoken on turns 8 and 4, so the 9th number spoken is the difference between them, 4.
# Turn 10: 4 is new, so the 10th number spoken is 0.
