## Getting Started with Postgres Data Types
Relational database systems store structured data for use by modern web applications, and **structured querying language** (SQL) is the language programmers use to store and retrieve data. There are several popular SQL databases used, including MySQL, MariaDB, PostgreSQL, and SQLite. The most popular **PostgreSQL** (aka Postgres) is an advanced open-source database supported by a large developer community. This article will review the benefit of using different Postgres data types.

## What are Postgres Data Types?
Collecting and processing data is a common application used by almost every web application. Most relational databases include common structured data types such as numbers, binary, and text. One of the reasons developers like Postgres is due to the extensive number of additional data types available for programmers to define in their queries. Postgres also includes the ability to define custom data types by combining multiple types.

## Why use Postgres Data Types?
Postgres data types improve data quality by validating the data in a database record matches the correct type. In a simple example, If a form requires an email address and the user enters a date, the form returns an error message stating incorrect data. Data types also define field size and page layout. Complex databases often require specific and well-defined data types to ensure the operation delivers accurate results. Postgres also defines an extensive list of data beyond text and numbers. You can store network addresses, geometric shapes, monetary types, JSON, and multidimensional arrays.

## Top Data Types
All relational database programs include a standard set of data types, including numeric, date, text, and boolean types. All data has specific characteristics that need to be consistent throughout calculations, operations, and reporting. Numbers are always numeric. Text always consists of text. A date has a different format than currency.

Data types also define the size of the fields, which in turn, affect page layout. There are fixed-length types and variable-length types.

Postgres includes an extensive set of predefined data types for structured, semi-structured, and non-structured (text) data. View a complete data type feature matrix.

Below we’ve highlighted some of these more common data types.

## Text and Character Data Types
Text and character types are either fixed or variable and determine space allocation in the database structure. These types include text, char, and varchar. There is no difference between type or char.

## Numeric Data Types
Postgres provides several numeric data types, including integers, floating points, arbitrary precision, and a special integer type called serial. Integers store numbers without fractions or decimals. Basic integers include integer, smallint, and bigint. Floating points describe numbers with decimals but without exact precision. Data with arbitrary precision are defined using numeric types. Serial types create unique ids or primary keys. Binary strings are stored using the bytea data type.

## Boolean Data Types
*True* and *false* statements are boolean types and are standard types featured in most SQL databases. These statements are stored as *true*, *false*, or *null*.

## Geometric Data Types
Website design often depends on shapes to contain information within the website page. Postgres geometric data types store two-dimensional spatial objects. Types include points, lines, lsig, box, path, polygon, and circle.

## Custom Data Types
Postgres is popular with developers for its ability to create custom data types by combining multiple standard types using the create type. The name of the new type must be unique. The create type includes five variations: Composite, Enumerated, Range, and Base. Composite types are a collection of other addressable types. Using the Composite type avoids the need to create a table to list the different types. Enumerated types list a static, ordered set of values, while Range types represent a range of values. You can create a brand new type using the Base type variation. Please note, the author of a new Base type must be a Postgres superuser since an error in the definition can confuse or even crash the server.

## Conclusion
Postgres is a robust, full-functional SQL database supported by a dynamic community of developers. The support for multiple data types is one of Postgres’ distinguishing features. In addition to storing standard types, the ability to define custom types and geometric shapes help Postgres stand out as an exceptional database tool.
