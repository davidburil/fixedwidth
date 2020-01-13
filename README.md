# Infer Column Widths of a Fixed-Width Text File
A fixed-width file use a fixed column width for each column (though not all columns necessarily have the same width) and pads the remaining space on the left or on the right, usually with spaces:

```
   287540 Smith  Jones  Accountant         $55,000
   204878 Ross   Betsy  Senior Accountant  $66,000
   208417 Arthur Wilbur CEO               $123,000
```

Parsing a fixed-width file can be difficult. Either, the user has to know the column widths in advance and pass that to a parsing method, or the method has to infer the widths of the columns. 
 
This package Go propose automated infer the columns.

----
Project inspired by blog post at [dev.to](https://dev.to/awwsmm/java-infer-column-widths-of-a-fixed-width-text-file-2hh0)