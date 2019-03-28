# box

**box** is a text-based window manager on terminal. Now, provide Box, ListBox, RotateBox.

[API Document](https://godoc.org/github.com/septemhill/box)

## Box
![image](https://github.com/septemhill/box/blob/master/box.gif)

Box means Textbox, you need give coordinate **X**, **Y** and box **Width**, **Height**.

[Box Example](https://github.com/septemhill/box/blob/master/example/box_example/main.go)

## ListBox
![image](https://github.com/septemhill/box/blob/master/listbox.gif)

ListBox also need coordinate **X**, **Y**, **Width**, **Height**, and additional **Display List** string

[ListBox Example](https://github.com/septemhill/box/blob/master/example/listbox_example/main.go)

## RotateBox
![image](https://github.com/septemhill/box/blob/master/rotatebox.gif)

RotateBox seems like ListBox, but there are a little different. It uses **display list count** instead of **height**

[RotateBox Example](https://github.com/septemhill/box/blob/master/example/rotatebox_example/main.go)
