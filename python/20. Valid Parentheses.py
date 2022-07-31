def isValid(s):
    if len(s) % 2 != 0:
        return False
    stack = []

    closingParenthesis = {")": "(", "}": "{", "]": "["}

    for i in range(len(s)):
        if s[i] in closingParenthesis:
            if stack and closingParenthesis[s[i]] == stack[-1]:
                stack.pop()
            else:
                return False

        else:
            stack.append(s[i])

    return False if stack else True
    print(x)

print(isValid("(){(})"))