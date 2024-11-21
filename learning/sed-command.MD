### The SED 

(Stream Editor) command in Unix/Linux is a powerful utility used to process and manipulate text in files. It can perform a variety of operations such as searching, find-and-replace, insertion, deletion, and more, without the need to open the file in an editor. This makes it a highly efficient tool for managing and editing text, especially when working with large files or automating tasks in scripts.

SED is a powerful text stream editor. Can do insertion, deletion, search, and replace(substitution).
SED command in UNIX supports regular expression which allows it to perform complex pattern matching.


#### string replace
 ```bash sed 's/word-infile/replacement-ofword/' input-filename 
    - s: string
    - word-infile: word that present in input-filename
    - replacement-ofword: replacement of word in input-filename ``` 
    
Every occurrences of word-infile has changed to replacement-ofword
    