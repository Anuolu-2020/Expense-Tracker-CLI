# Expense Tracker CLI

## Overview

This is a command-line interface (CLI) tool for managing expenses. You can use the following commands to 
interact with the expense tracker:

- `tracker -ls`: List all expenses in the database.
- `tracker -add`  Add a new expense.
- `tracker -clear`: Delete all expenses from the database.

## Usage
### List Expenses

To list all expenses in the database, use the following command:
```bash
tracker -ls

#Expected result
ID   AMOUNT CATEGORY  DESCRIPTION        DATE

```
To add expenses in the database, use the following command:
```bash
tracker -add
## Provide the values step by step

What's the amount
2000

What category is the expense
Monthly

Expense Description
Bought Pizza

Expense added succesfully
```
To delete expenses in the database, use the following command:
```
tracker -clear

## Expected result
Expenses data deleted
```
