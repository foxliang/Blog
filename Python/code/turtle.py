"""
用turtle模块绘图
这是一个非常有趣的模块 它模拟一只乌龟在窗口上爬行的方式来进行绘图
"""

import turtle

turtle.pensize(2)
turtle.penup()
turtle.goto(-180, 150)
turtle.pencolor('red')
turtle.fillcolor('yellow')
turtle.pendown()
turtle.begin_fill()
for _ in range(36):
    turtle.forward(200)
    turtle.right(170)
turtle.end_fill()
turtle.mainloop()
