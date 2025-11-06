Kood Notes Tool

This is a command-line tool for managing simple, single-line notes. It allows you to group notes into different collections, such as shopping_list, work_ideas, or movies_to_watch.
How to Use It

1. Run the tool:

To start, run the program from your terminal with the name of the note collection. If the collection doesn't exist, it will be created.

./notestool my_first_collection


2. Use the menu:

Once it's running, you will see a menu that lets you show your current notes, add a new one, delete an old one, or exit the program.
How Your Notes Are Stored

Notes are saved in a plain text file located in the project folder. The file will have the same name as the collection you create. This makes it simple and portable, and the file can be opened in a regular text editor.

A Quick Example

Here’s what a typical session looks like:

$ ./notestool my_todos
Welcome to the Kood notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
Finish the project report

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - Finish the project report

