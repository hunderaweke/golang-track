# Library Management System User Guide

This guide provides instructions for using the command-line interface (CLI) of the Library Management System. Follow the steps below to perform various operations like adding books, borrowing books, and managing library members.

## Starting the Application

1. Open a terminal.
2. Navigate to the directory containing the compiled application.
3. Run the application:

   ```sh
   ./library_management
   ```

## Menu Options

Upon starting the application, you will be presented with the following menu:

```
Choose an option:
1. Add Book
2. Get Books
3. Get Available Books
4. Get Borrowed Books
5. Add Member
6. Get Members
7. Borrow Book
8. Return Book
9. Exit
```

### 1. Add Book

To add a new book to the library:

1. Select option `1` by entering `1` and pressing `Enter`.
2. Enter the book title and press `Enter`.
3. Enter the book author and press `Enter`.
4. You will see a confirmation message with the book ID.

### 2. Get Books

To list all books in the library:

1. Select option `2` by entering `2` and pressing `Enter`.
2. The application will display a list of all books with their IDs, titles, and authors.

### 3. Get Available Books

To list all available (not borrowed) books:

1. Select option `3` by entering `3` and pressing `Enter`.
2. The application will display a list of all available books with their IDs, titles, and authors.

### 4. Get Borrowed Books

To list all books borrowed by a specific member:

1. Select option `4` by entering `4` and pressing `Enter`.
2. Enter the member ID and press `Enter`.
3. The application will display a list of all books borrowed by the specified member.

### 5. Add Member

To add a new member to the library:

1. Select option `5` by entering `5` and pressing `Enter`.
2. Enter the member's name and press `Enter`.
3. You will see a confirmation message with the member ID.

### 6. Get Members

To list all members of the library:

1. Select option `6` by entering `6` and pressing `Enter`.
2. The application will display a list of all members with their IDs and names.

### 7. Borrow Book

To borrow a book for a member:

1. Select option `7` by entering `7` and pressing `Enter`.
2. Enter the book ID and press `Enter`.
3. Enter the member ID and press `Enter`.
4. The application will display a confirmation message if the book was successfully borrowed, or an error message if the book is not available.

### 8. Return Book

To return a borrowed book:

1. Select option `8` by entering `8` and pressing `Enter`.
2. Enter the book ID and press `Enter`.
3. Enter the member ID and press `Enter`.
4. The application will display a confirmation message if the book was successfully returned, or an error message if the book was not borrowed by the member.

### 9. Exit

To exit the application:

1. Select option `9` by entering `9` and pressing `Enter`.
2. The application will close.

