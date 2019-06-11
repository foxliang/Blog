"""
百分制成绩转等级制
>90    --> A
80-89  --> B
70-79  --> C
60-69  --> D
<60    --> E
"""


score = float(input('请输入成绩：'))
if score >= 90:
    grade =' A'
elif score >= 80:
    geade = 'B'
elif score >= 70:
    grade = 'C'
elif score >= 60:
    grade = 'D'
else:
    grade = 'E'
print('对应的等级是:',grade)
